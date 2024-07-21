// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"mocku/backend/ent/careers"
	"mocku/backend/ent/professor"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CareersCreate is the builder for creating a Careers entity.
type CareersCreate struct {
	config
	mutation *CareersMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cc *CareersCreate) SetName(s string) *CareersCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetDescription sets the "description" field.
func (cc *CareersCreate) SetDescription(s string) *CareersCreate {
	cc.mutation.SetDescription(s)
	return cc
}

// AddLeaderIDs adds the "leader" edge to the Professor entity by IDs.
func (cc *CareersCreate) AddLeaderIDs(ids ...int) *CareersCreate {
	cc.mutation.AddLeaderIDs(ids...)
	return cc
}

// AddLeader adds the "leader" edges to the Professor entity.
func (cc *CareersCreate) AddLeader(p ...*Professor) *CareersCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cc.AddLeaderIDs(ids...)
}

// Mutation returns the CareersMutation object of the builder.
func (cc *CareersCreate) Mutation() *CareersMutation {
	return cc.mutation
}

// Save creates the Careers in the database.
func (cc *CareersCreate) Save(ctx context.Context) (*Careers, error) {
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CareersCreate) SaveX(ctx context.Context) *Careers {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CareersCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CareersCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CareersCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Careers.name"`)}
	}
	if v, ok := cc.mutation.Name(); ok {
		if err := careers.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Careers.name": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Careers.description"`)}
	}
	if v, ok := cc.mutation.Description(); ok {
		if err := careers.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Careers.description": %w`, err)}
		}
	}
	return nil
}

func (cc *CareersCreate) sqlSave(ctx context.Context) (*Careers, error) {
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

func (cc *CareersCreate) createSpec() (*Careers, *sqlgraph.CreateSpec) {
	var (
		_node = &Careers{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(careers.Table, sqlgraph.NewFieldSpec(careers.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(careers.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.Description(); ok {
		_spec.SetField(careers.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if nodes := cc.mutation.LeaderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   careers.LeaderTable,
			Columns: []string{careers.LeaderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(professor.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CareersCreateBulk is the builder for creating many Careers entities in bulk.
type CareersCreateBulk struct {
	config
	err      error
	builders []*CareersCreate
}

// Save creates the Careers entities in the database.
func (ccb *CareersCreateBulk) Save(ctx context.Context) ([]*Careers, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Careers, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CareersMutation)
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
func (ccb *CareersCreateBulk) SaveX(ctx context.Context) []*Careers {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CareersCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CareersCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
