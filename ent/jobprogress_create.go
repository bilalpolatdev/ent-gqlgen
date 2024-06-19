// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"gqlgen-ent/ent/jobdetail"
	"gqlgen-ent/ent/jobprogress"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// JobProgressCreate is the builder for creating a JobProgress entity.
type JobProgressCreate struct {
	config
	mutation *JobProgressMutation
	hooks    []Hook
}

// SetOne sets the "One" field.
func (jpc *JobProgressCreate) SetOne(i int) *JobProgressCreate {
	jpc.mutation.SetOne(i)
	return jpc
}

// SetNillableOne sets the "One" field if the given value is not nil.
func (jpc *JobProgressCreate) SetNillableOne(i *int) *JobProgressCreate {
	if i != nil {
		jpc.SetOne(*i)
	}
	return jpc
}

// SetTwo sets the "Two" field.
func (jpc *JobProgressCreate) SetTwo(i int) *JobProgressCreate {
	jpc.mutation.SetTwo(i)
	return jpc
}

// SetNillableTwo sets the "Two" field if the given value is not nil.
func (jpc *JobProgressCreate) SetNillableTwo(i *int) *JobProgressCreate {
	if i != nil {
		jpc.SetTwo(*i)
	}
	return jpc
}

// SetThree sets the "Three" field.
func (jpc *JobProgressCreate) SetThree(i int) *JobProgressCreate {
	jpc.mutation.SetThree(i)
	return jpc
}

// SetNillableThree sets the "Three" field if the given value is not nil.
func (jpc *JobProgressCreate) SetNillableThree(i *int) *JobProgressCreate {
	if i != nil {
		jpc.SetThree(*i)
	}
	return jpc
}

// SetFour sets the "Four" field.
func (jpc *JobProgressCreate) SetFour(i int) *JobProgressCreate {
	jpc.mutation.SetFour(i)
	return jpc
}

// SetNillableFour sets the "Four" field if the given value is not nil.
func (jpc *JobProgressCreate) SetNillableFour(i *int) *JobProgressCreate {
	if i != nil {
		jpc.SetFour(*i)
	}
	return jpc
}

// SetFive sets the "Five" field.
func (jpc *JobProgressCreate) SetFive(i int) *JobProgressCreate {
	jpc.mutation.SetFive(i)
	return jpc
}

// SetNillableFive sets the "Five" field if the given value is not nil.
func (jpc *JobProgressCreate) SetNillableFive(i *int) *JobProgressCreate {
	if i != nil {
		jpc.SetFive(*i)
	}
	return jpc
}

// SetSix sets the "Six" field.
func (jpc *JobProgressCreate) SetSix(i int) *JobProgressCreate {
	jpc.mutation.SetSix(i)
	return jpc
}

// SetNillableSix sets the "Six" field if the given value is not nil.
func (jpc *JobProgressCreate) SetNillableSix(i *int) *JobProgressCreate {
	if i != nil {
		jpc.SetSix(*i)
	}
	return jpc
}

// SetCreatedAt sets the "created_at" field.
func (jpc *JobProgressCreate) SetCreatedAt(t time.Time) *JobProgressCreate {
	jpc.mutation.SetCreatedAt(t)
	return jpc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (jpc *JobProgressCreate) SetNillableCreatedAt(t *time.Time) *JobProgressCreate {
	if t != nil {
		jpc.SetCreatedAt(*t)
	}
	return jpc
}

// SetUpdatedAt sets the "updated_at" field.
func (jpc *JobProgressCreate) SetUpdatedAt(t time.Time) *JobProgressCreate {
	jpc.mutation.SetUpdatedAt(t)
	return jpc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (jpc *JobProgressCreate) SetNillableUpdatedAt(t *time.Time) *JobProgressCreate {
	if t != nil {
		jpc.SetUpdatedAt(*t)
	}
	return jpc
}

// AddProgresIDs adds the "progress" edge to the JobDetail entity by IDs.
func (jpc *JobProgressCreate) AddProgresIDs(ids ...int) *JobProgressCreate {
	jpc.mutation.AddProgresIDs(ids...)
	return jpc
}

// AddProgress adds the "progress" edges to the JobDetail entity.
func (jpc *JobProgressCreate) AddProgress(j ...*JobDetail) *JobProgressCreate {
	ids := make([]int, len(j))
	for i := range j {
		ids[i] = j[i].ID
	}
	return jpc.AddProgresIDs(ids...)
}

// Mutation returns the JobProgressMutation object of the builder.
func (jpc *JobProgressCreate) Mutation() *JobProgressMutation {
	return jpc.mutation
}

// Save creates the JobProgress in the database.
func (jpc *JobProgressCreate) Save(ctx context.Context) (*JobProgress, error) {
	jpc.defaults()
	return withHooks(ctx, jpc.sqlSave, jpc.mutation, jpc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (jpc *JobProgressCreate) SaveX(ctx context.Context) *JobProgress {
	v, err := jpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jpc *JobProgressCreate) Exec(ctx context.Context) error {
	_, err := jpc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jpc *JobProgressCreate) ExecX(ctx context.Context) {
	if err := jpc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (jpc *JobProgressCreate) defaults() {
	if _, ok := jpc.mutation.One(); !ok {
		v := jobprogress.DefaultOne
		jpc.mutation.SetOne(v)
	}
	if _, ok := jpc.mutation.Two(); !ok {
		v := jobprogress.DefaultTwo
		jpc.mutation.SetTwo(v)
	}
	if _, ok := jpc.mutation.Three(); !ok {
		v := jobprogress.DefaultThree
		jpc.mutation.SetThree(v)
	}
	if _, ok := jpc.mutation.Four(); !ok {
		v := jobprogress.DefaultFour
		jpc.mutation.SetFour(v)
	}
	if _, ok := jpc.mutation.Five(); !ok {
		v := jobprogress.DefaultFive
		jpc.mutation.SetFive(v)
	}
	if _, ok := jpc.mutation.Six(); !ok {
		v := jobprogress.DefaultSix
		jpc.mutation.SetSix(v)
	}
	if _, ok := jpc.mutation.CreatedAt(); !ok {
		v := jobprogress.DefaultCreatedAt()
		jpc.mutation.SetCreatedAt(v)
	}
	if _, ok := jpc.mutation.UpdatedAt(); !ok {
		v := jobprogress.DefaultUpdatedAt()
		jpc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (jpc *JobProgressCreate) check() error {
	if _, ok := jpc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "JobProgress.created_at"`)}
	}
	if _, ok := jpc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "JobProgress.updated_at"`)}
	}
	return nil
}

func (jpc *JobProgressCreate) sqlSave(ctx context.Context) (*JobProgress, error) {
	if err := jpc.check(); err != nil {
		return nil, err
	}
	_node, _spec := jpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, jpc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	jpc.mutation.id = &_node.ID
	jpc.mutation.done = true
	return _node, nil
}

func (jpc *JobProgressCreate) createSpec() (*JobProgress, *sqlgraph.CreateSpec) {
	var (
		_node = &JobProgress{config: jpc.config}
		_spec = sqlgraph.NewCreateSpec(jobprogress.Table, sqlgraph.NewFieldSpec(jobprogress.FieldID, field.TypeInt))
	)
	if value, ok := jpc.mutation.One(); ok {
		_spec.SetField(jobprogress.FieldOne, field.TypeInt, value)
		_node.One = value
	}
	if value, ok := jpc.mutation.Two(); ok {
		_spec.SetField(jobprogress.FieldTwo, field.TypeInt, value)
		_node.Two = value
	}
	if value, ok := jpc.mutation.Three(); ok {
		_spec.SetField(jobprogress.FieldThree, field.TypeInt, value)
		_node.Three = value
	}
	if value, ok := jpc.mutation.Four(); ok {
		_spec.SetField(jobprogress.FieldFour, field.TypeInt, value)
		_node.Four = value
	}
	if value, ok := jpc.mutation.Five(); ok {
		_spec.SetField(jobprogress.FieldFive, field.TypeInt, value)
		_node.Five = value
	}
	if value, ok := jpc.mutation.Six(); ok {
		_spec.SetField(jobprogress.FieldSix, field.TypeInt, value)
		_node.Six = value
	}
	if value, ok := jpc.mutation.CreatedAt(); ok {
		_spec.SetField(jobprogress.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := jpc.mutation.UpdatedAt(); ok {
		_spec.SetField(jobprogress.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := jpc.mutation.ProgressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   jobprogress.ProgressTable,
			Columns: []string{jobprogress.ProgressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(jobdetail.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// JobProgressCreateBulk is the builder for creating many JobProgress entities in bulk.
type JobProgressCreateBulk struct {
	config
	err      error
	builders []*JobProgressCreate
}

// Save creates the JobProgress entities in the database.
func (jpcb *JobProgressCreateBulk) Save(ctx context.Context) ([]*JobProgress, error) {
	if jpcb.err != nil {
		return nil, jpcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(jpcb.builders))
	nodes := make([]*JobProgress, len(jpcb.builders))
	mutators := make([]Mutator, len(jpcb.builders))
	for i := range jpcb.builders {
		func(i int, root context.Context) {
			builder := jpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*JobProgressMutation)
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
					_, err = mutators[i+1].Mutate(root, jpcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, jpcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, jpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (jpcb *JobProgressCreateBulk) SaveX(ctx context.Context) []*JobProgress {
	v, err := jpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jpcb *JobProgressCreateBulk) Exec(ctx context.Context) error {
	_, err := jpcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jpcb *JobProgressCreateBulk) ExecX(ctx context.Context) {
	if err := jpcb.Exec(ctx); err != nil {
		panic(err)
	}
}