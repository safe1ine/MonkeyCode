package repo

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"

	"github.com/chaitin/MonkeyCode/backend/consts"
	"github.com/chaitin/MonkeyCode/backend/db"
	"github.com/chaitin/MonkeyCode/backend/db/apikey"
	"github.com/chaitin/MonkeyCode/backend/db/model"
	"github.com/chaitin/MonkeyCode/backend/db/task"
	"github.com/chaitin/MonkeyCode/backend/db/taskrecord"
	"github.com/chaitin/MonkeyCode/backend/domain"
	"github.com/chaitin/MonkeyCode/backend/pkg/entx"
)

type ProxyRepo struct {
	db    *db.Client
	redis *redis.Client
}

func NewProxyRepo(db *db.Client, redis *redis.Client) domain.ProxyRepo {
	return &ProxyRepo{db: db, redis: redis}
}

func (r *ProxyRepo) SelectModelWithLoadBalancing(modelName string, modelType consts.ModelType) (*db.Model, error) {
	ctx := context.Background()
	m, err := r.db.Model.Query().
		Where(model.ModelName(modelName), model.ModelType(modelType)).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *ProxyRepo) ValidateApiKey(ctx context.Context, key string) (*db.ApiKey, error) {
	rkey := "sk-" + key
	data, err := r.redis.Get(ctx, rkey).Result()
	if err == nil {
		key := db.ApiKey{}
		if err := json.Unmarshal([]byte(data), &key); err != nil {
			return nil, err
		}
		return &key, nil
	}

	if !errors.Is(err, redis.Nil) {
		return nil, err
	}

	a, err := r.db.ApiKey.Query().
		Where(apikey.Key(key), apikey.Status(consts.ApiKeyStatusActive)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}

	if err := r.redis.Set(ctx, rkey, string(b), 24*time.Hour).Err(); err != nil {
		return nil, err
	}

	return a, nil
}

func (r *ProxyRepo) Record(ctx context.Context, record *domain.RecordParam) error {
	userID, err := uuid.Parse(record.UserID)
	if err != nil {
		return err
	}
	modelID, err := uuid.Parse(record.ModelID)
	if err != nil {
		return err
	}

	return entx.WithTx(ctx, r.db, func(tx *db.Tx) error {
		isNew := false
		t, err := tx.Task.Query().ForUpdate().Where(task.TaskID(record.TaskID)).Only(ctx)
		if db.IsNotFound(err) {
			t, err = tx.Task.Create().
				SetTaskID(record.TaskID).
				SetRequestID(record.RequestID).
				SetUserID(userID).
				SetModelID(modelID).
				SetProgramLanguage(record.ProgramLanguage).
				SetInputTokens(record.InputTokens).
				SetOutputTokens(record.OutputTokens).
				SetIsAccept(record.IsAccept).
				SetModelType(record.ModelType).
				SetWorkMode(record.WorkMode).
				SetCodeLines(record.CodeLines).
				Save(ctx)
			isNew = true
		}
		if err != nil {
			return err
		}
		if !isNew {
			up := tx.Task.UpdateOneID(t.ID)
			if record.OutputTokens > 0 {
				up.AddOutputTokens(record.OutputTokens)
			}
			if t.InputTokens == 0 && record.InputTokens > 0 {
				up.SetInputTokens(record.InputTokens)
			}
			if t.RequestID != record.RequestID {
				up.SetRequestID(record.RequestID)
				up.AddInputTokens(record.InputTokens)
			}
			if err := up.Exec(ctx); err != nil {
				return err
			}
		}

		if record.Role == consts.ChatRoleUser {
			count, err := tx.TaskRecord.Query().
				Where(
					taskrecord.TaskID(t.ID),
					taskrecord.Role(consts.ChatRoleUser),
					taskrecord.Prompt(record.Prompt),
				).
				Count(ctx)
			if err != nil {
				return err
			}
			if count > 0 {
				return nil
			}
		}

		_, err = tx.TaskRecord.Create().
			SetTaskID(t.ID).
			SetRole(record.Role).
			SetPrompt(record.Prompt).
			SetCompletion(record.Completion).
			SetOutputTokens(record.OutputTokens).
			Save(ctx)

		return err
	})
}

func (r *ProxyRepo) UpdateByTaskID(ctx context.Context, taskID string, fn func(*db.TaskUpdateOne)) error {
	rc, err := r.db.Task.Query().Where(task.TaskID(taskID)).Only(ctx)
	if err != nil {
		return err
	}

	return entx.WithTx(ctx, r.db, func(tx *db.Tx) error {
		up := tx.Task.UpdateOneID(rc.ID)
		fn(up)
		return up.Exec(ctx)
	})
}

func (r *ProxyRepo) AcceptCompletion(ctx context.Context, req *domain.AcceptCompletionReq) error {
	return entx.WithTx(ctx, r.db, func(tx *db.Tx) error {
		rc, err := tx.Task.Query().Where(task.TaskID(req.ID)).Only(ctx)
		if err != nil {
			return err
		}

		if err := tx.Task.UpdateOneID(rc.ID).
			SetIsAccept(true).
			SetCompletion(req.Completion).
			Exec(ctx); err != nil {
			return err
		}

		return tx.TaskRecord.Update().
			Where(taskrecord.TaskID(rc.ID)).
			SetCompletion(req.Completion).Exec(ctx)
	})
}
