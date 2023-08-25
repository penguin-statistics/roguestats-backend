// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"exusiai.dev/roguestats-backend/internal/ent/event"
	"exusiai.dev/roguestats-backend/internal/ent/research"
)

// ResearchCreate is the builder for creating a Research entity.
type ResearchCreate struct {
	config
	mutation *ResearchMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (rc *ResearchCreate) SetName(s string) *ResearchCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetSchema sets the "schema" field.
func (rc *ResearchCreate) SetSchema(m map[string]interface{}) *ResearchCreate {
	rc.mutation.SetSchema(m)
	return rc
}

// SetID sets the "id" field.
func (rc *ResearchCreate) SetID(s string) *ResearchCreate {
	rc.mutation.SetID(s)
	return rc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (rc *ResearchCreate) SetNillableID(s *string) *ResearchCreate {
	if s != nil {
		rc.SetID(*s)
	}
	return rc
}

// AddEventIDs adds the "events" edge to the Event entity by IDs.
func (rc *ResearchCreate) AddEventIDs(ids ...string) *ResearchCreate {
	rc.mutation.AddEventIDs(ids...)
	return rc
}

// AddEvents adds the "events" edges to the Event entity.
func (rc *ResearchCreate) AddEvents(e ...*Event) *ResearchCreate {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return rc.AddEventIDs(ids...)
}

// Mutation returns the ResearchMutation object of the builder.
func (rc *ResearchCreate) Mutation() *ResearchMutation {
	return rc.mutation
}

// Save creates the Research in the database.
func (rc *ResearchCreate) Save(ctx context.Context) (*Research, error) {
	rc.defaults()
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *ResearchCreate) SaveX(ctx context.Context) *Research {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *ResearchCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *ResearchCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *ResearchCreate) defaults() {
	if _, ok := rc.mutation.ID(); !ok {
		v := research.DefaultID()
		rc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *ResearchCreate) check() error {
	if _, ok := rc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Research.name"`)}
	}
	if v, ok := rc.mutation.Name(); ok {
		if err := research.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Research.name": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Schema(); !ok {
		return &ValidationError{Name: "schema", err: errors.New(`ent: missing required field "Research.schema"`)}
	}
	return nil
}

func (rc *ResearchCreate) sqlSave(ctx context.Context) (*Research, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Research.ID type: %T", _spec.ID.Value)
		}
	}
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *ResearchCreate) createSpec() (*Research, *sqlgraph.CreateSpec) {
	var (
		_node = &Research{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(research.Table, sqlgraph.NewFieldSpec(research.FieldID, field.TypeString))
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.Name(); ok {
		_spec.SetField(research.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := rc.mutation.Schema(); ok {
		_spec.SetField(research.FieldSchema, field.TypeJSON, value)
		_node.Schema = value
	}
	if nodes := rc.mutation.EventsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   research.EventsTable,
			Columns: []string{research.EventsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(event.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ResearchCreateBulk is the builder for creating many Research entities in bulk.
type ResearchCreateBulk struct {
	config
	builders []*ResearchCreate
}

// Save creates the Research entities in the database.
func (rcb *ResearchCreateBulk) Save(ctx context.Context) ([]*Research, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Research, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ResearchMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *ResearchCreateBulk) SaveX(ctx context.Context) []*Research {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *ResearchCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *ResearchCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
