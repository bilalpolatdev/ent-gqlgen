// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"gqlgen-ent/ent/companycareer"
	"gqlgen-ent/ent/companyengineer"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CompanyCareerCreate is the builder for creating a CompanyCareer entity.
type CompanyCareerCreate struct {
	config
	mutation *CompanyCareerMutation
	hooks    []Hook
}

// SetCareer sets the "Career" field.
func (ccc *CompanyCareerCreate) SetCareer(s string) *CompanyCareerCreate {
	ccc.mutation.SetCareer(s)
	return ccc
}

// SetNillableCareer sets the "Career" field if the given value is not nil.
func (ccc *CompanyCareerCreate) SetNillableCareer(s *string) *CompanyCareerCreate {
	if s != nil {
		ccc.SetCareer(*s)
	}
	return ccc
}

// AddEngineerCareerIDs adds the "engineerCareers" edge to the CompanyEngineer entity by IDs.
func (ccc *CompanyCareerCreate) AddEngineerCareerIDs(ids ...int) *CompanyCareerCreate {
	ccc.mutation.AddEngineerCareerIDs(ids...)
	return ccc
}

// AddEngineerCareers adds the "engineerCareers" edges to the CompanyEngineer entity.
func (ccc *CompanyCareerCreate) AddEngineerCareers(c ...*CompanyEngineer) *CompanyCareerCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ccc.AddEngineerCareerIDs(ids...)
}

// Mutation returns the CompanyCareerMutation object of the builder.
func (ccc *CompanyCareerCreate) Mutation() *CompanyCareerMutation {
	return ccc.mutation
}

// Save creates the CompanyCareer in the database.
func (ccc *CompanyCareerCreate) Save(ctx context.Context) (*CompanyCareer, error) {
	return withHooks(ctx, ccc.sqlSave, ccc.mutation, ccc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ccc *CompanyCareerCreate) SaveX(ctx context.Context) *CompanyCareer {
	v, err := ccc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccc *CompanyCareerCreate) Exec(ctx context.Context) error {
	_, err := ccc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccc *CompanyCareerCreate) ExecX(ctx context.Context) {
	if err := ccc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ccc *CompanyCareerCreate) check() error {
	return nil
}

func (ccc *CompanyCareerCreate) sqlSave(ctx context.Context) (*CompanyCareer, error) {
	if err := ccc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ccc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ccc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ccc.mutation.id = &_node.ID
	ccc.mutation.done = true
	return _node, nil
}

func (ccc *CompanyCareerCreate) createSpec() (*CompanyCareer, *sqlgraph.CreateSpec) {
	var (
		_node = &CompanyCareer{config: ccc.config}
		_spec = sqlgraph.NewCreateSpec(companycareer.Table, sqlgraph.NewFieldSpec(companycareer.FieldID, field.TypeInt))
	)
	if value, ok := ccc.mutation.Career(); ok {
		_spec.SetField(companycareer.FieldCareer, field.TypeString, value)
		_node.Career = value
	}
	if nodes := ccc.mutation.EngineerCareersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   companycareer.EngineerCareersTable,
			Columns: []string{companycareer.EngineerCareersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(companyengineer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CompanyCareerCreateBulk is the builder for creating many CompanyCareer entities in bulk.
type CompanyCareerCreateBulk struct {
	config
	err      error
	builders []*CompanyCareerCreate
}

// Save creates the CompanyCareer entities in the database.
func (cccb *CompanyCareerCreateBulk) Save(ctx context.Context) ([]*CompanyCareer, error) {
	if cccb.err != nil {
		return nil, cccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(cccb.builders))
	nodes := make([]*CompanyCareer, len(cccb.builders))
	mutators := make([]Mutator, len(cccb.builders))
	for i := range cccb.builders {
		func(i int, root context.Context) {
			builder := cccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CompanyCareerMutation)
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
					_, err = mutators[i+1].Mutate(root, cccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, cccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cccb *CompanyCareerCreateBulk) SaveX(ctx context.Context) []*CompanyCareer {
	v, err := cccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cccb *CompanyCareerCreateBulk) Exec(ctx context.Context) error {
	_, err := cccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cccb *CompanyCareerCreateBulk) ExecX(ctx context.Context) {
	if err := cccb.Exec(ctx); err != nil {
		panic(err)
	}
}
