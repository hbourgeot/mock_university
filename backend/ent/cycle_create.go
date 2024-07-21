// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mocku/backend/ent/cycle"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CycleCreate is the builder for creating a Cycle entity.
type CycleCreate struct {
	config
	mutation *CycleMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cc *CycleCreate) SetName(s string) *CycleCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetStartDate sets the "start_date" field.
func (cc *CycleCreate) SetStartDate(t time.Time) *CycleCreate {
	cc.mutation.SetStartDate(t)
	return cc
}

// SetNillableStartDate sets the "start_date" field if the given value is not nil.
func (cc *CycleCreate) SetNillableStartDate(t *time.Time) *CycleCreate {
	if t != nil {
		cc.SetStartDate(*t)
	}
	return cc
}

// SetEndDate sets the "end_date" field.
func (cc *CycleCreate) SetEndDate(t time.Time) *CycleCreate {
	cc.mutation.SetEndDate(t)
	return cc
}

// SetNillableEndDate sets the "end_date" field if the given value is not nil.
func (cc *CycleCreate) SetNillableEndDate(t *time.Time) *CycleCreate {
	if t != nil {
		cc.SetEndDate(*t)
	}
	return cc
}

// Mutation returns the CycleMutation object of the builder.
func (cc *CycleCreate) Mutation() *CycleMutation {
	return cc.mutation
}

// Save creates the Cycle in the database.
func (cc *CycleCreate) Save(ctx context.Context) (*Cycle, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CycleCreate) SaveX(ctx context.Context) *Cycle {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CycleCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CycleCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CycleCreate) defaults() {
	if _, ok := cc.mutation.StartDate(); !ok {
		v := cycle.DefaultStartDate()
		cc.mutation.SetStartDate(v)
	}
	if _, ok := cc.mutation.EndDate(); !ok {
		v := cycle.DefaultEndDate()
		cc.mutation.SetEndDate(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CycleCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Cycle.name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := cycle.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Cycle.name": %w`, err)}
		}
	}
	if _, ok := cc.mutation.StartDate(); !ok {
		return &ValidationError{Name: "start_date", err: errors.New(`ent: missing required field "Cycle.start_date"`)}
	}
	if _, ok := cc.mutation.EndDate(); !ok {
		return &ValidationError{Name: "end_date", err: errors.New(`ent: missing required field "Cycle.end_date"`)}
	}
	return nil
}

func (cc *CycleCreate) sqlSave(ctx context.Context) (*Cycle, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CycleCreate) createSpec() (*Cycle, *sqlgraph.CreateSpec) {
	var (
		_node = &Cycle{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(cycle.Table, sqlgraph.NewFieldSpec(cycle.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(cycle.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.StartDate(); ok {
		_spec.SetField(cycle.FieldStartDate, field.TypeTime, value)
		_node.StartDate = value
	}
	if value, ok := cc.mutation.EndDate(); ok {
		_spec.SetField(cycle.FieldEndDate, field.TypeTime, value)
		_node.EndDate = value
	}
	return _node, _spec
}

// CycleCreateBulk is the builder for creating many Cycle entities in bulk.
type CycleCreateBulk struct {
	config
	err      error
	builders []*CycleCreate
}

// Save creates the Cycle entities in the database.
func (ccb *CycleCreateBulk) Save(ctx context.Context) ([]*Cycle, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Cycle, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CycleMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CycleCreateBulk) SaveX(ctx context.Context) []*Cycle {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CycleCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CycleCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
