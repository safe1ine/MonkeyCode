// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/chaitin/MonkeyCode/backend/db/extension"
	"github.com/chaitin/MonkeyCode/backend/db/predicate"
)

// ExtensionDelete is the builder for deleting a Extension entity.
type ExtensionDelete struct {
	config
	hooks    []Hook
	mutation *ExtensionMutation
}

// Where appends a list predicates to the ExtensionDelete builder.
func (ed *ExtensionDelete) Where(ps ...predicate.Extension) *ExtensionDelete {
	ed.mutation.Where(ps...)
	return ed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ed *ExtensionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ed.sqlExec, ed.mutation, ed.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ed *ExtensionDelete) ExecX(ctx context.Context) int {
	n, err := ed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ed *ExtensionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(extension.Table, sqlgraph.NewFieldSpec(extension.FieldID, field.TypeUUID))
	if ps := ed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ed.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ed.mutation.done = true
	return affected, err
}

// ExtensionDeleteOne is the builder for deleting a single Extension entity.
type ExtensionDeleteOne struct {
	ed *ExtensionDelete
}

// Where appends a list predicates to the ExtensionDelete builder.
func (edo *ExtensionDeleteOne) Where(ps ...predicate.Extension) *ExtensionDeleteOne {
	edo.ed.mutation.Where(ps...)
	return edo
}

// Exec executes the deletion query.
func (edo *ExtensionDeleteOne) Exec(ctx context.Context) error {
	n, err := edo.ed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{extension.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (edo *ExtensionDeleteOne) ExecX(ctx context.Context) {
	if err := edo.Exec(ctx); err != nil {
		panic(err)
	}
}
