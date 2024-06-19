// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"gqlgen-ent/ent/companyowner"
	"gqlgen-ent/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CompanyOwnerDelete is the builder for deleting a CompanyOwner entity.
type CompanyOwnerDelete struct {
	config
	hooks    []Hook
	mutation *CompanyOwnerMutation
}

// Where appends a list predicates to the CompanyOwnerDelete builder.
func (cod *CompanyOwnerDelete) Where(ps ...predicate.CompanyOwner) *CompanyOwnerDelete {
	cod.mutation.Where(ps...)
	return cod
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cod *CompanyOwnerDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cod.sqlExec, cod.mutation, cod.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cod *CompanyOwnerDelete) ExecX(ctx context.Context) int {
	n, err := cod.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cod *CompanyOwnerDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(companyowner.Table, sqlgraph.NewFieldSpec(companyowner.FieldID, field.TypeInt))
	if ps := cod.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cod.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cod.mutation.done = true
	return affected, err
}

// CompanyOwnerDeleteOne is the builder for deleting a single CompanyOwner entity.
type CompanyOwnerDeleteOne struct {
	cod *CompanyOwnerDelete
}

// Where appends a list predicates to the CompanyOwnerDelete builder.
func (codo *CompanyOwnerDeleteOne) Where(ps ...predicate.CompanyOwner) *CompanyOwnerDeleteOne {
	codo.cod.mutation.Where(ps...)
	return codo
}

// Exec executes the deletion query.
func (codo *CompanyOwnerDeleteOne) Exec(ctx context.Context) error {
	n, err := codo.cod.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{companyowner.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (codo *CompanyOwnerDeleteOne) ExecX(ctx context.Context) {
	if err := codo.Exec(ctx); err != nil {
		panic(err)
	}
}