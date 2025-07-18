package domain

import (
	"context"
	"net/http"

	"github.com/rokku-c/go-openai"

	"github.com/chaitin/MonkeyCode/backend/consts"
	"github.com/chaitin/MonkeyCode/backend/db"
)

type Proxy interface {
	AcceptCompletion(ctx context.Context, req *AcceptCompletionReq) error
	HandleCompletion(ctx context.Context, w http.ResponseWriter, req CompletionRequest)
	HandleChatCompletion(ctx context.Context, w http.ResponseWriter, req *openai.ChatCompletionRequest)
	HandleEmbeddings(ctx context.Context, w http.ResponseWriter, req *openai.EmbeddingRequest)
}

type ProxyUsecase interface {
	SelectModelWithLoadBalancing(modelName string, modelType consts.ModelType) (*Model, error)
	Record(ctx context.Context, record *RecordParam) error
	ValidateApiKey(ctx context.Context, key string) (*ApiKey, error)
	AcceptCompletion(ctx context.Context, req *AcceptCompletionReq) error
}

type ProxyRepo interface {
	Record(ctx context.Context, record *RecordParam) error
	UpdateByTaskID(ctx context.Context, taskID string, fn func(*db.TaskUpdateOne)) error
	AcceptCompletion(ctx context.Context, req *AcceptCompletionReq) error
	SelectModelWithLoadBalancing(modelName string, modelType consts.ModelType) (*db.Model, error)
	ValidateApiKey(ctx context.Context, key string) (*db.ApiKey, error)
}

type VersionInfo struct {
	Version string `json:"version"`
	URL     string `json:"url"`
}

type AcceptCompletionReq struct {
	ID         string `json:"id"`         // 记录ID
	Completion string `json:"completion"` // 补全内容
}

type RecordParam struct {
	RequestID       string
	TaskID          string
	UserID          string
	ModelID         string
	ModelType       consts.ModelType
	Role            consts.ChatRole
	Prompt          string
	ProgramLanguage string
	InputTokens     int64
	OutputTokens    int64
	IsAccept        bool
	Completion      string
	WorkMode        string
	CodeLines       int64
}

func (r *RecordParam) Clone() *RecordParam {
	return &RecordParam{
		RequestID:       r.RequestID,
		TaskID:          r.TaskID,
		UserID:          r.UserID,
		ModelID:         r.ModelID,
		ModelType:       r.ModelType,
		Role:            r.Role,
		Prompt:          r.Prompt,
		ProgramLanguage: r.ProgramLanguage,
		InputTokens:     r.InputTokens,
		OutputTokens:    r.OutputTokens,
		IsAccept:        r.IsAccept,
		Completion:      r.Completion,
		WorkMode:        r.WorkMode,
		CodeLines:       r.CodeLines,
	}
}
