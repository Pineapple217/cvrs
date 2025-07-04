// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Pineapple217/cvrs/pkg/ent/predicate"
	"github.com/Pineapple217/cvrs/pkg/ent/releaseappearance"
)

// ReleaseAppearanceDelete is the builder for deleting a ReleaseAppearance entity.
type ReleaseAppearanceDelete struct {
	config
	hooks    []Hook
	mutation *ReleaseAppearanceMutation
}

// Where appends a list predicates to the ReleaseAppearanceDelete builder.
func (rad *ReleaseAppearanceDelete) Where(ps ...predicate.ReleaseAppearance) *ReleaseAppearanceDelete {
	rad.mutation.Where(ps...)
	return rad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rad *ReleaseAppearanceDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rad.sqlExec, rad.mutation, rad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rad *ReleaseAppearanceDelete) ExecX(ctx context.Context) int {
	n, err := rad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rad *ReleaseAppearanceDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(releaseappearance.Table, nil)
	if ps := rad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rad.mutation.done = true
	return affected, err
}

// ReleaseAppearanceDeleteOne is the builder for deleting a single ReleaseAppearance entity.
type ReleaseAppearanceDeleteOne struct {
	rad *ReleaseAppearanceDelete
}

// Where appends a list predicates to the ReleaseAppearanceDelete builder.
func (rado *ReleaseAppearanceDeleteOne) Where(ps ...predicate.ReleaseAppearance) *ReleaseAppearanceDeleteOne {
	rado.rad.mutation.Where(ps...)
	return rado
}

// Exec executes the deletion query.
func (rado *ReleaseAppearanceDeleteOne) Exec(ctx context.Context) error {
	n, err := rado.rad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{releaseappearance.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rado *ReleaseAppearanceDeleteOne) ExecX(ctx context.Context) {
	if err := rado.Exec(ctx); err != nil {
		panic(err)
	}
}
