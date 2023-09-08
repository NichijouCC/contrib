// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/contrib/entproto/internal/entprototest/ent/messagewithints"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessageWithIntsCreate is the builder for creating a MessageWithInts entity.
type MessageWithIntsCreate struct {
	config
	mutation *MessageWithIntsMutation
	hooks    []Hook
}

// SetInt32s sets the "int32s" field.
func (mwic *MessageWithIntsCreate) SetInt32s(i []int32) *MessageWithIntsCreate {
	mwic.mutation.SetInt32s(i)
	return mwic
}

// SetInt64s sets the "int64s" field.
func (mwic *MessageWithIntsCreate) SetInt64s(i []int64) *MessageWithIntsCreate {
	mwic.mutation.SetInt64s(i)
	return mwic
}

// SetUint32s sets the "uint32s" field.
func (mwic *MessageWithIntsCreate) SetUint32s(u []uint32) *MessageWithIntsCreate {
	mwic.mutation.SetUint32s(u)
	return mwic
}

// SetUint64s sets the "uint64s" field.
func (mwic *MessageWithIntsCreate) SetUint64s(u []uint64) *MessageWithIntsCreate {
	mwic.mutation.SetUint64s(u)
	return mwic
}

// Mutation returns the MessageWithIntsMutation object of the builder.
func (mwic *MessageWithIntsCreate) Mutation() *MessageWithIntsMutation {
	return mwic.mutation
}

// Save creates the MessageWithInts in the database.
func (mwic *MessageWithIntsCreate) Save(ctx context.Context) (*MessageWithInts, error) {
	return withHooks[*MessageWithInts, MessageWithIntsMutation](ctx, mwic.sqlSave, mwic.mutation, mwic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mwic *MessageWithIntsCreate) SaveX(ctx context.Context) *MessageWithInts {
	v, err := mwic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mwic *MessageWithIntsCreate) Exec(ctx context.Context) error {
	_, err := mwic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mwic *MessageWithIntsCreate) ExecX(ctx context.Context) {
	if err := mwic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mwic *MessageWithIntsCreate) check() error {
	if _, ok := mwic.mutation.Int32s(); !ok {
		return &ValidationError{Name: "int32s", err: errors.New(`ent: missing required field "MessageWithInts.int32s"`)}
	}
	if _, ok := mwic.mutation.Int64s(); !ok {
		return &ValidationError{Name: "int64s", err: errors.New(`ent: missing required field "MessageWithInts.int64s"`)}
	}
	if _, ok := mwic.mutation.Uint32s(); !ok {
		return &ValidationError{Name: "uint32s", err: errors.New(`ent: missing required field "MessageWithInts.uint32s"`)}
	}
	if _, ok := mwic.mutation.Uint64s(); !ok {
		return &ValidationError{Name: "uint64s", err: errors.New(`ent: missing required field "MessageWithInts.uint64s"`)}
	}
	return nil
}

func (mwic *MessageWithIntsCreate) sqlSave(ctx context.Context) (*MessageWithInts, error) {
	if err := mwic.check(); err != nil {
		return nil, err
	}
	_node, _spec := mwic.createSpec()
	if err := sqlgraph.CreateNode(ctx, mwic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	mwic.mutation.id = &_node.ID
	mwic.mutation.done = true
	return _node, nil
}

func (mwic *MessageWithIntsCreate) createSpec() (*MessageWithInts, *sqlgraph.CreateSpec) {
	var (
		_node = &MessageWithInts{config: mwic.config}
		_spec = sqlgraph.NewCreateSpec(messagewithints.Table, sqlgraph.NewFieldSpec(messagewithints.FieldID, field.TypeInt))
	)
	if value, ok := mwic.mutation.Int32s(); ok {
		_spec.SetField(messagewithints.FieldInt32s, field.TypeJSON, value)
		_node.Int32s = value
	}
	if value, ok := mwic.mutation.Int64s(); ok {
		_spec.SetField(messagewithints.FieldInt64s, field.TypeJSON, value)
		_node.Int64s = value
	}
	if value, ok := mwic.mutation.Uint32s(); ok {
		_spec.SetField(messagewithints.FieldUint32s, field.TypeJSON, value)
		_node.Uint32s = value
	}
	if value, ok := mwic.mutation.Uint64s(); ok {
		_spec.SetField(messagewithints.FieldUint64s, field.TypeJSON, value)
		_node.Uint64s = value
	}
	return _node, _spec
}

// MessageWithIntsCreateBulk is the builder for creating many MessageWithInts entities in bulk.
type MessageWithIntsCreateBulk struct {
	config
	builders []*MessageWithIntsCreate
}

// Save creates the MessageWithInts entities in the database.
func (mwicb *MessageWithIntsCreateBulk) Save(ctx context.Context) ([]*MessageWithInts, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mwicb.builders))
	nodes := make([]*MessageWithInts, len(mwicb.builders))
	mutators := make([]Mutator, len(mwicb.builders))
	for i := range mwicb.builders {
		func(i int, root context.Context) {
			builder := mwicb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MessageWithIntsMutation)
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
					_, err = mutators[i+1].Mutate(root, mwicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mwicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mwicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mwicb *MessageWithIntsCreateBulk) SaveX(ctx context.Context) []*MessageWithInts {
	v, err := mwicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mwicb *MessageWithIntsCreateBulk) Exec(ctx context.Context) error {
	_, err := mwicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mwicb *MessageWithIntsCreateBulk) ExecX(ctx context.Context) {
	if err := mwicb.Exec(ctx); err != nil {
		panic(err)
	}
}