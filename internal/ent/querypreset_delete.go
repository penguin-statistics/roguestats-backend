// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"exusiai.dev/roguestats-backend/internal/ent/predicate"
	"exusiai.dev/roguestats-backend/internal/ent/querypreset"
)

// QueryPresetDelete is the builder for deleting a QueryPreset entity.
type QueryPresetDelete struct {
	config
	hooks    []Hook
	mutation *QueryPresetMutation
}

// Where appends a list predicates to the QueryPresetDelete builder.
func (qpd *QueryPresetDelete) Where(ps ...predicate.QueryPreset) *QueryPresetDelete {
	qpd.mutation.Where(ps...)
	return qpd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (qpd *QueryPresetDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, qpd.sqlExec, qpd.mutation, qpd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (qpd *QueryPresetDelete) ExecX(ctx context.Context) int {
	n, err := qpd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (qpd *QueryPresetDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(querypreset.Table, sqlgraph.NewFieldSpec(querypreset.FieldID, field.TypeString))
	if ps := qpd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, qpd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	qpd.mutation.done = true
	return affected, err
}

// QueryPresetDeleteOne is the builder for deleting a single QueryPreset entity.
type QueryPresetDeleteOne struct {
	qpd *QueryPresetDelete
}

// Where appends a list predicates to the QueryPresetDelete builder.
func (qpdo *QueryPresetDeleteOne) Where(ps ...predicate.QueryPreset) *QueryPresetDeleteOne {
	qpdo.qpd.mutation.Where(ps...)
	return qpdo
}

// Exec executes the deletion query.
func (qpdo *QueryPresetDeleteOne) Exec(ctx context.Context) error {
	n, err := qpdo.qpd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{querypreset.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (qpdo *QueryPresetDeleteOne) ExecX(ctx context.Context) {
	if err := qpdo.Exec(ctx); err != nil {
		panic(err)
	}
}
