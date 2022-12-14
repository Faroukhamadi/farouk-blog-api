// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Faroukhamadi/farouk-blog-api/ent/content"
	"github.com/Faroukhamadi/farouk-blog-api/ent/post"
	"github.com/Faroukhamadi/farouk-blog-api/ent/predicate"
)

// ContentUpdate is the builder for updating Content entities.
type ContentUpdate struct {
	config
	hooks    []Hook
	mutation *ContentMutation
}

// Where appends a list predicates to the ContentUpdate builder.
func (cu *ContentUpdate) Where(ps ...predicate.Content) *ContentUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *ContentUpdate) SetUpdatedAt(t time.Time) *ContentUpdate {
	cu.mutation.SetUpdatedAt(t)
	return cu
}

// SetTitle sets the "title" field.
func (cu *ContentUpdate) SetTitle(s string) *ContentUpdate {
	cu.mutation.SetTitle(s)
	return cu
}

// SetText sets the "text" field.
func (cu *ContentUpdate) SetText(s string) *ContentUpdate {
	cu.mutation.SetText(s)
	return cu
}

// SetPostID sets the "post" edge to the Post entity by ID.
func (cu *ContentUpdate) SetPostID(id int) *ContentUpdate {
	cu.mutation.SetPostID(id)
	return cu
}

// SetNillablePostID sets the "post" edge to the Post entity by ID if the given value is not nil.
func (cu *ContentUpdate) SetNillablePostID(id *int) *ContentUpdate {
	if id != nil {
		cu = cu.SetPostID(*id)
	}
	return cu
}

// SetPost sets the "post" edge to the Post entity.
func (cu *ContentUpdate) SetPost(p *Post) *ContentUpdate {
	return cu.SetPostID(p.ID)
}

// Mutation returns the ContentMutation object of the builder.
func (cu *ContentUpdate) Mutation() *ContentMutation {
	return cu.mutation
}

// ClearPost clears the "post" edge to the Post entity.
func (cu *ContentUpdate) ClearPost() *ContentUpdate {
	cu.mutation.ClearPost()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ContentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cu.defaults()
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ContentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ContentUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ContentUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ContentUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *ContentUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := content.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

func (cu *ContentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   content.Table,
			Columns: content.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: content.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: content.FieldUpdatedAt,
		})
	}
	if value, ok := cu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: content.FieldTitle,
		})
	}
	if value, ok := cu.mutation.Text(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: content.FieldText,
		})
	}
	if cu.mutation.PostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   content.PostTable,
			Columns: []string{content.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   content.PostTable,
			Columns: []string{content.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{content.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ContentUpdateOne is the builder for updating a single Content entity.
type ContentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ContentMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *ContentUpdateOne) SetUpdatedAt(t time.Time) *ContentUpdateOne {
	cuo.mutation.SetUpdatedAt(t)
	return cuo
}

// SetTitle sets the "title" field.
func (cuo *ContentUpdateOne) SetTitle(s string) *ContentUpdateOne {
	cuo.mutation.SetTitle(s)
	return cuo
}

// SetText sets the "text" field.
func (cuo *ContentUpdateOne) SetText(s string) *ContentUpdateOne {
	cuo.mutation.SetText(s)
	return cuo
}

// SetPostID sets the "post" edge to the Post entity by ID.
func (cuo *ContentUpdateOne) SetPostID(id int) *ContentUpdateOne {
	cuo.mutation.SetPostID(id)
	return cuo
}

// SetNillablePostID sets the "post" edge to the Post entity by ID if the given value is not nil.
func (cuo *ContentUpdateOne) SetNillablePostID(id *int) *ContentUpdateOne {
	if id != nil {
		cuo = cuo.SetPostID(*id)
	}
	return cuo
}

// SetPost sets the "post" edge to the Post entity.
func (cuo *ContentUpdateOne) SetPost(p *Post) *ContentUpdateOne {
	return cuo.SetPostID(p.ID)
}

// Mutation returns the ContentMutation object of the builder.
func (cuo *ContentUpdateOne) Mutation() *ContentMutation {
	return cuo.mutation
}

// ClearPost clears the "post" edge to the Post entity.
func (cuo *ContentUpdateOne) ClearPost() *ContentUpdateOne {
	cuo.mutation.ClearPost()
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ContentUpdateOne) Select(field string, fields ...string) *ContentUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Content entity.
func (cuo *ContentUpdateOne) Save(ctx context.Context) (*Content, error) {
	var (
		err  error
		node *Content
	)
	cuo.defaults()
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ContentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Content)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ContentMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ContentUpdateOne) SaveX(ctx context.Context) *Content {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ContentUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ContentUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *ContentUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := content.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

func (cuo *ContentUpdateOne) sqlSave(ctx context.Context) (_node *Content, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   content.Table,
			Columns: content.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: content.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Content.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, content.FieldID)
		for _, f := range fields {
			if !content.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != content.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: content.FieldUpdatedAt,
		})
	}
	if value, ok := cuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: content.FieldTitle,
		})
	}
	if value, ok := cuo.mutation.Text(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: content.FieldText,
		})
	}
	if cuo.mutation.PostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   content.PostTable,
			Columns: []string{content.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   content.PostTable,
			Columns: []string{content.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: post.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Content{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{content.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
