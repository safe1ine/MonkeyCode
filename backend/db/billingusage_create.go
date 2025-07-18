// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chaitin/MonkeyCode/backend/db/billingusage"
)

// BillingUsageCreate is the builder for creating a BillingUsage entity.
type BillingUsageCreate struct {
	config
	mutation *BillingUsageMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetDeletedAt sets the "deleted_at" field.
func (buc *BillingUsageCreate) SetDeletedAt(t time.Time) *BillingUsageCreate {
	buc.mutation.SetDeletedAt(t)
	return buc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (buc *BillingUsageCreate) SetNillableDeletedAt(t *time.Time) *BillingUsageCreate {
	if t != nil {
		buc.SetDeletedAt(*t)
	}
	return buc
}

// SetUserID sets the "user_id" field.
func (buc *BillingUsageCreate) SetUserID(s string) *BillingUsageCreate {
	buc.mutation.SetUserID(s)
	return buc
}

// SetModelName sets the "model_name" field.
func (buc *BillingUsageCreate) SetModelName(s string) *BillingUsageCreate {
	buc.mutation.SetModelName(s)
	return buc
}

// SetTokens sets the "tokens" field.
func (buc *BillingUsageCreate) SetTokens(i int64) *BillingUsageCreate {
	buc.mutation.SetTokens(i)
	return buc
}

// SetOperation sets the "operation" field.
func (buc *BillingUsageCreate) SetOperation(s string) *BillingUsageCreate {
	buc.mutation.SetOperation(s)
	return buc
}

// SetCreatedAt sets the "created_at" field.
func (buc *BillingUsageCreate) SetCreatedAt(t time.Time) *BillingUsageCreate {
	buc.mutation.SetCreatedAt(t)
	return buc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (buc *BillingUsageCreate) SetNillableCreatedAt(t *time.Time) *BillingUsageCreate {
	if t != nil {
		buc.SetCreatedAt(*t)
	}
	return buc
}

// SetUpdatedAt sets the "updated_at" field.
func (buc *BillingUsageCreate) SetUpdatedAt(t time.Time) *BillingUsageCreate {
	buc.mutation.SetUpdatedAt(t)
	return buc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (buc *BillingUsageCreate) SetNillableUpdatedAt(t *time.Time) *BillingUsageCreate {
	if t != nil {
		buc.SetUpdatedAt(*t)
	}
	return buc
}

// SetID sets the "id" field.
func (buc *BillingUsageCreate) SetID(s string) *BillingUsageCreate {
	buc.mutation.SetID(s)
	return buc
}

// Mutation returns the BillingUsageMutation object of the builder.
func (buc *BillingUsageCreate) Mutation() *BillingUsageMutation {
	return buc.mutation
}

// Save creates the BillingUsage in the database.
func (buc *BillingUsageCreate) Save(ctx context.Context) (*BillingUsage, error) {
	if err := buc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, buc.sqlSave, buc.mutation, buc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (buc *BillingUsageCreate) SaveX(ctx context.Context) *BillingUsage {
	v, err := buc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (buc *BillingUsageCreate) Exec(ctx context.Context) error {
	_, err := buc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buc *BillingUsageCreate) ExecX(ctx context.Context) {
	if err := buc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (buc *BillingUsageCreate) defaults() error {
	if _, ok := buc.mutation.CreatedAt(); !ok {
		if billingusage.DefaultCreatedAt == nil {
			return fmt.Errorf("db: uninitialized billingusage.DefaultCreatedAt (forgotten import db/runtime?)")
		}
		v := billingusage.DefaultCreatedAt()
		buc.mutation.SetCreatedAt(v)
	}
	if _, ok := buc.mutation.UpdatedAt(); !ok {
		if billingusage.DefaultUpdatedAt == nil {
			return fmt.Errorf("db: uninitialized billingusage.DefaultUpdatedAt (forgotten import db/runtime?)")
		}
		v := billingusage.DefaultUpdatedAt()
		buc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (buc *BillingUsageCreate) check() error {
	if _, ok := buc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`db: missing required field "BillingUsage.user_id"`)}
	}
	if _, ok := buc.mutation.ModelName(); !ok {
		return &ValidationError{Name: "model_name", err: errors.New(`db: missing required field "BillingUsage.model_name"`)}
	}
	if _, ok := buc.mutation.Tokens(); !ok {
		return &ValidationError{Name: "tokens", err: errors.New(`db: missing required field "BillingUsage.tokens"`)}
	}
	if _, ok := buc.mutation.Operation(); !ok {
		return &ValidationError{Name: "operation", err: errors.New(`db: missing required field "BillingUsage.operation"`)}
	}
	if _, ok := buc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`db: missing required field "BillingUsage.created_at"`)}
	}
	if _, ok := buc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`db: missing required field "BillingUsage.updated_at"`)}
	}
	return nil
}

func (buc *BillingUsageCreate) sqlSave(ctx context.Context) (*BillingUsage, error) {
	if err := buc.check(); err != nil {
		return nil, err
	}
	_node, _spec := buc.createSpec()
	if err := sqlgraph.CreateNode(ctx, buc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected BillingUsage.ID type: %T", _spec.ID.Value)
		}
	}
	buc.mutation.id = &_node.ID
	buc.mutation.done = true
	return _node, nil
}

func (buc *BillingUsageCreate) createSpec() (*BillingUsage, *sqlgraph.CreateSpec) {
	var (
		_node = &BillingUsage{config: buc.config}
		_spec = sqlgraph.NewCreateSpec(billingusage.Table, sqlgraph.NewFieldSpec(billingusage.FieldID, field.TypeString))
	)
	_spec.OnConflict = buc.conflict
	if id, ok := buc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := buc.mutation.DeletedAt(); ok {
		_spec.SetField(billingusage.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := buc.mutation.UserID(); ok {
		_spec.SetField(billingusage.FieldUserID, field.TypeString, value)
		_node.UserID = value
	}
	if value, ok := buc.mutation.ModelName(); ok {
		_spec.SetField(billingusage.FieldModelName, field.TypeString, value)
		_node.ModelName = value
	}
	if value, ok := buc.mutation.Tokens(); ok {
		_spec.SetField(billingusage.FieldTokens, field.TypeInt64, value)
		_node.Tokens = value
	}
	if value, ok := buc.mutation.Operation(); ok {
		_spec.SetField(billingusage.FieldOperation, field.TypeString, value)
		_node.Operation = value
	}
	if value, ok := buc.mutation.CreatedAt(); ok {
		_spec.SetField(billingusage.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := buc.mutation.UpdatedAt(); ok {
		_spec.SetField(billingusage.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.BillingUsage.Create().
//		SetDeletedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BillingUsageUpsert) {
//			SetDeletedAt(v+v).
//		}).
//		Exec(ctx)
func (buc *BillingUsageCreate) OnConflict(opts ...sql.ConflictOption) *BillingUsageUpsertOne {
	buc.conflict = opts
	return &BillingUsageUpsertOne{
		create: buc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.BillingUsage.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (buc *BillingUsageCreate) OnConflictColumns(columns ...string) *BillingUsageUpsertOne {
	buc.conflict = append(buc.conflict, sql.ConflictColumns(columns...))
	return &BillingUsageUpsertOne{
		create: buc,
	}
}

type (
	// BillingUsageUpsertOne is the builder for "upsert"-ing
	//  one BillingUsage node.
	BillingUsageUpsertOne struct {
		create *BillingUsageCreate
	}

	// BillingUsageUpsert is the "OnConflict" setter.
	BillingUsageUpsert struct {
		*sql.UpdateSet
	}
)

// SetDeletedAt sets the "deleted_at" field.
func (u *BillingUsageUpsert) SetDeletedAt(v time.Time) *BillingUsageUpsert {
	u.Set(billingusage.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *BillingUsageUpsert) UpdateDeletedAt() *BillingUsageUpsert {
	u.SetExcluded(billingusage.FieldDeletedAt)
	return u
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *BillingUsageUpsert) ClearDeletedAt() *BillingUsageUpsert {
	u.SetNull(billingusage.FieldDeletedAt)
	return u
}

// SetUserID sets the "user_id" field.
func (u *BillingUsageUpsert) SetUserID(v string) *BillingUsageUpsert {
	u.Set(billingusage.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *BillingUsageUpsert) UpdateUserID() *BillingUsageUpsert {
	u.SetExcluded(billingusage.FieldUserID)
	return u
}

// SetModelName sets the "model_name" field.
func (u *BillingUsageUpsert) SetModelName(v string) *BillingUsageUpsert {
	u.Set(billingusage.FieldModelName, v)
	return u
}

// UpdateModelName sets the "model_name" field to the value that was provided on create.
func (u *BillingUsageUpsert) UpdateModelName() *BillingUsageUpsert {
	u.SetExcluded(billingusage.FieldModelName)
	return u
}

// SetTokens sets the "tokens" field.
func (u *BillingUsageUpsert) SetTokens(v int64) *BillingUsageUpsert {
	u.Set(billingusage.FieldTokens, v)
	return u
}

// UpdateTokens sets the "tokens" field to the value that was provided on create.
func (u *BillingUsageUpsert) UpdateTokens() *BillingUsageUpsert {
	u.SetExcluded(billingusage.FieldTokens)
	return u
}

// AddTokens adds v to the "tokens" field.
func (u *BillingUsageUpsert) AddTokens(v int64) *BillingUsageUpsert {
	u.Add(billingusage.FieldTokens, v)
	return u
}

// SetOperation sets the "operation" field.
func (u *BillingUsageUpsert) SetOperation(v string) *BillingUsageUpsert {
	u.Set(billingusage.FieldOperation, v)
	return u
}

// UpdateOperation sets the "operation" field to the value that was provided on create.
func (u *BillingUsageUpsert) UpdateOperation() *BillingUsageUpsert {
	u.SetExcluded(billingusage.FieldOperation)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *BillingUsageUpsert) SetCreatedAt(v time.Time) *BillingUsageUpsert {
	u.Set(billingusage.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *BillingUsageUpsert) UpdateCreatedAt() *BillingUsageUpsert {
	u.SetExcluded(billingusage.FieldCreatedAt)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *BillingUsageUpsert) SetUpdatedAt(v time.Time) *BillingUsageUpsert {
	u.Set(billingusage.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *BillingUsageUpsert) UpdateUpdatedAt() *BillingUsageUpsert {
	u.SetExcluded(billingusage.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.BillingUsage.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(billingusage.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *BillingUsageUpsertOne) UpdateNewValues() *BillingUsageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(billingusage.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.BillingUsage.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *BillingUsageUpsertOne) Ignore() *BillingUsageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BillingUsageUpsertOne) DoNothing() *BillingUsageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BillingUsageCreate.OnConflict
// documentation for more info.
func (u *BillingUsageUpsertOne) Update(set func(*BillingUsageUpsert)) *BillingUsageUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BillingUsageUpsert{UpdateSet: update})
	}))
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *BillingUsageUpsertOne) SetDeletedAt(v time.Time) *BillingUsageUpsertOne {
	return u.Update(func(s *BillingUsageUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *BillingUsageUpsertOne) UpdateDeletedAt() *BillingUsageUpsertOne {
	return u.Update(func(s *BillingUsageUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *BillingUsageUpsertOne) ClearDeletedAt() *BillingUsageUpsertOne {
	return u.Update(func(s *BillingUsageUpsert) {
		s.ClearDeletedAt()
	})
}

// SetUserID sets the "user_id" field.
func (u *BillingUsageUpsertOne) SetUserID(v string) *BillingUsageUpsertOne {
	return u.Update(func(s *BillingUsageUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *BillingUsageUpsertOne) UpdateUserID() *BillingUsageUpsertOne {
	return u.Update(func(s *BillingUsageUpsert) {
		s.UpdateUserID()
	})
}

// SetModelName sets the "model_name" field.
func (u *BillingUsageUpsertOne) SetModelName(v string) *BillingUsageUpsertOne {
	return u.Update(func(s *BillingUsageUpsert) {
		s.SetModelName(v)
	})
}

// UpdateModelName sets the "model_name" field to the value that was provided on create.
func (u *BillingUsageUpsertOne) UpdateModelName() *BillingUsageUpsertOne {
	return u.Update(func(s *BillingUsageUpsert) {
		s.UpdateModelName()
	})
}

// SetTokens sets the "tokens" field.
func (u *BillingUsageUpsertOne) SetTokens(v int64) *BillingUsageUpsertOne {
	return u.Update(func(s *BillingUsageUpsert) {
		s.SetTokens(v)
	})
}

// AddTokens adds v to the "tokens" field.
func (u *BillingUsageUpsertOne) AddTokens(v int64) *BillingUsageUpsertOne {
	return u.Update(func(s *BillingUsageUpsert) {
		s.AddTokens(v)
	})
}

// UpdateTokens sets the "tokens" field to the value that was provided on create.
func (u *BillingUsageUpsertOne) UpdateTokens() *BillingUsageUpsertOne {
	return u.Update(func(s *BillingUsageUpsert) {
		s.UpdateTokens()
	})
}

// SetOperation sets the "operation" field.
func (u *BillingUsageUpsertOne) SetOperation(v string) *BillingUsageUpsertOne {
	return u.Update(func(s *BillingUsageUpsert) {
		s.SetOperation(v)
	})
}

// UpdateOperation sets the "operation" field to the value that was provided on create.
func (u *BillingUsageUpsertOne) UpdateOperation() *BillingUsageUpsertOne {
	return u.Update(func(s *BillingUsageUpsert) {
		s.UpdateOperation()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *BillingUsageUpsertOne) SetCreatedAt(v time.Time) *BillingUsageUpsertOne {
	return u.Update(func(s *BillingUsageUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *BillingUsageUpsertOne) UpdateCreatedAt() *BillingUsageUpsertOne {
	return u.Update(func(s *BillingUsageUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *BillingUsageUpsertOne) SetUpdatedAt(v time.Time) *BillingUsageUpsertOne {
	return u.Update(func(s *BillingUsageUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *BillingUsageUpsertOne) UpdateUpdatedAt() *BillingUsageUpsertOne {
	return u.Update(func(s *BillingUsageUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *BillingUsageUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("db: missing options for BillingUsageCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BillingUsageUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *BillingUsageUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("db: BillingUsageUpsertOne.ID is not supported by MySQL driver. Use BillingUsageUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *BillingUsageUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// BillingUsageCreateBulk is the builder for creating many BillingUsage entities in bulk.
type BillingUsageCreateBulk struct {
	config
	err      error
	builders []*BillingUsageCreate
	conflict []sql.ConflictOption
}

// Save creates the BillingUsage entities in the database.
func (bucb *BillingUsageCreateBulk) Save(ctx context.Context) ([]*BillingUsage, error) {
	if bucb.err != nil {
		return nil, bucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(bucb.builders))
	nodes := make([]*BillingUsage, len(bucb.builders))
	mutators := make([]Mutator, len(bucb.builders))
	for i := range bucb.builders {
		func(i int, root context.Context) {
			builder := bucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BillingUsageMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = bucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bucb *BillingUsageCreateBulk) SaveX(ctx context.Context) []*BillingUsage {
	v, err := bucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bucb *BillingUsageCreateBulk) Exec(ctx context.Context) error {
	_, err := bucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bucb *BillingUsageCreateBulk) ExecX(ctx context.Context) {
	if err := bucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.BillingUsage.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.BillingUsageUpsert) {
//			SetDeletedAt(v+v).
//		}).
//		Exec(ctx)
func (bucb *BillingUsageCreateBulk) OnConflict(opts ...sql.ConflictOption) *BillingUsageUpsertBulk {
	bucb.conflict = opts
	return &BillingUsageUpsertBulk{
		create: bucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.BillingUsage.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (bucb *BillingUsageCreateBulk) OnConflictColumns(columns ...string) *BillingUsageUpsertBulk {
	bucb.conflict = append(bucb.conflict, sql.ConflictColumns(columns...))
	return &BillingUsageUpsertBulk{
		create: bucb,
	}
}

// BillingUsageUpsertBulk is the builder for "upsert"-ing
// a bulk of BillingUsage nodes.
type BillingUsageUpsertBulk struct {
	create *BillingUsageCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.BillingUsage.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(billingusage.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *BillingUsageUpsertBulk) UpdateNewValues() *BillingUsageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(billingusage.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.BillingUsage.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *BillingUsageUpsertBulk) Ignore() *BillingUsageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *BillingUsageUpsertBulk) DoNothing() *BillingUsageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the BillingUsageCreateBulk.OnConflict
// documentation for more info.
func (u *BillingUsageUpsertBulk) Update(set func(*BillingUsageUpsert)) *BillingUsageUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&BillingUsageUpsert{UpdateSet: update})
	}))
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *BillingUsageUpsertBulk) SetDeletedAt(v time.Time) *BillingUsageUpsertBulk {
	return u.Update(func(s *BillingUsageUpsert) {
		s.SetDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *BillingUsageUpsertBulk) UpdateDeletedAt() *BillingUsageUpsertBulk {
	return u.Update(func(s *BillingUsageUpsert) {
		s.UpdateDeletedAt()
	})
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (u *BillingUsageUpsertBulk) ClearDeletedAt() *BillingUsageUpsertBulk {
	return u.Update(func(s *BillingUsageUpsert) {
		s.ClearDeletedAt()
	})
}

// SetUserID sets the "user_id" field.
func (u *BillingUsageUpsertBulk) SetUserID(v string) *BillingUsageUpsertBulk {
	return u.Update(func(s *BillingUsageUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *BillingUsageUpsertBulk) UpdateUserID() *BillingUsageUpsertBulk {
	return u.Update(func(s *BillingUsageUpsert) {
		s.UpdateUserID()
	})
}

// SetModelName sets the "model_name" field.
func (u *BillingUsageUpsertBulk) SetModelName(v string) *BillingUsageUpsertBulk {
	return u.Update(func(s *BillingUsageUpsert) {
		s.SetModelName(v)
	})
}

// UpdateModelName sets the "model_name" field to the value that was provided on create.
func (u *BillingUsageUpsertBulk) UpdateModelName() *BillingUsageUpsertBulk {
	return u.Update(func(s *BillingUsageUpsert) {
		s.UpdateModelName()
	})
}

// SetTokens sets the "tokens" field.
func (u *BillingUsageUpsertBulk) SetTokens(v int64) *BillingUsageUpsertBulk {
	return u.Update(func(s *BillingUsageUpsert) {
		s.SetTokens(v)
	})
}

// AddTokens adds v to the "tokens" field.
func (u *BillingUsageUpsertBulk) AddTokens(v int64) *BillingUsageUpsertBulk {
	return u.Update(func(s *BillingUsageUpsert) {
		s.AddTokens(v)
	})
}

// UpdateTokens sets the "tokens" field to the value that was provided on create.
func (u *BillingUsageUpsertBulk) UpdateTokens() *BillingUsageUpsertBulk {
	return u.Update(func(s *BillingUsageUpsert) {
		s.UpdateTokens()
	})
}

// SetOperation sets the "operation" field.
func (u *BillingUsageUpsertBulk) SetOperation(v string) *BillingUsageUpsertBulk {
	return u.Update(func(s *BillingUsageUpsert) {
		s.SetOperation(v)
	})
}

// UpdateOperation sets the "operation" field to the value that was provided on create.
func (u *BillingUsageUpsertBulk) UpdateOperation() *BillingUsageUpsertBulk {
	return u.Update(func(s *BillingUsageUpsert) {
		s.UpdateOperation()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *BillingUsageUpsertBulk) SetCreatedAt(v time.Time) *BillingUsageUpsertBulk {
	return u.Update(func(s *BillingUsageUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *BillingUsageUpsertBulk) UpdateCreatedAt() *BillingUsageUpsertBulk {
	return u.Update(func(s *BillingUsageUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *BillingUsageUpsertBulk) SetUpdatedAt(v time.Time) *BillingUsageUpsertBulk {
	return u.Update(func(s *BillingUsageUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *BillingUsageUpsertBulk) UpdateUpdatedAt() *BillingUsageUpsertBulk {
	return u.Update(func(s *BillingUsageUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *BillingUsageUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("db: OnConflict was set for builder %d. Set it on the BillingUsageCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("db: missing options for BillingUsageCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *BillingUsageUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
