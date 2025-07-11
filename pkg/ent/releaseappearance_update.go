// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Pineapple217/cvrs/pkg/ent/artist"
	"github.com/Pineapple217/cvrs/pkg/ent/predicate"
	"github.com/Pineapple217/cvrs/pkg/ent/release"
	"github.com/Pineapple217/cvrs/pkg/ent/releaseappearance"
	"github.com/Pineapple217/cvrs/pkg/pid"
)

// ReleaseAppearanceUpdate is the builder for updating ReleaseAppearance entities.
type ReleaseAppearanceUpdate struct {
	config
	hooks    []Hook
	mutation *ReleaseAppearanceMutation
}

// Where appends a list predicates to the ReleaseAppearanceUpdate builder.
func (rau *ReleaseAppearanceUpdate) Where(ps ...predicate.ReleaseAppearance) *ReleaseAppearanceUpdate {
	rau.mutation.Where(ps...)
	return rau
}

// SetReleaseID sets the "release_id" field.
func (rau *ReleaseAppearanceUpdate) SetReleaseID(pi pid.ID) *ReleaseAppearanceUpdate {
	rau.mutation.SetReleaseID(pi)
	return rau
}

// SetNillableReleaseID sets the "release_id" field if the given value is not nil.
func (rau *ReleaseAppearanceUpdate) SetNillableReleaseID(pi *pid.ID) *ReleaseAppearanceUpdate {
	if pi != nil {
		rau.SetReleaseID(*pi)
	}
	return rau
}

// SetArtistID sets the "artist_id" field.
func (rau *ReleaseAppearanceUpdate) SetArtistID(pi pid.ID) *ReleaseAppearanceUpdate {
	rau.mutation.SetArtistID(pi)
	return rau
}

// SetNillableArtistID sets the "artist_id" field if the given value is not nil.
func (rau *ReleaseAppearanceUpdate) SetNillableArtistID(pi *pid.ID) *ReleaseAppearanceUpdate {
	if pi != nil {
		rau.SetArtistID(*pi)
	}
	return rau
}

// SetOrder sets the "order" field.
func (rau *ReleaseAppearanceUpdate) SetOrder(i int) *ReleaseAppearanceUpdate {
	rau.mutation.ResetOrder()
	rau.mutation.SetOrder(i)
	return rau
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (rau *ReleaseAppearanceUpdate) SetNillableOrder(i *int) *ReleaseAppearanceUpdate {
	if i != nil {
		rau.SetOrder(*i)
	}
	return rau
}

// AddOrder adds i to the "order" field.
func (rau *ReleaseAppearanceUpdate) AddOrder(i int) *ReleaseAppearanceUpdate {
	rau.mutation.AddOrder(i)
	return rau
}

// SetArtist sets the "artist" edge to the Artist entity.
func (rau *ReleaseAppearanceUpdate) SetArtist(a *Artist) *ReleaseAppearanceUpdate {
	return rau.SetArtistID(a.ID)
}

// SetRelease sets the "release" edge to the Release entity.
func (rau *ReleaseAppearanceUpdate) SetRelease(r *Release) *ReleaseAppearanceUpdate {
	return rau.SetReleaseID(r.ID)
}

// Mutation returns the ReleaseAppearanceMutation object of the builder.
func (rau *ReleaseAppearanceUpdate) Mutation() *ReleaseAppearanceMutation {
	return rau.mutation
}

// ClearArtist clears the "artist" edge to the Artist entity.
func (rau *ReleaseAppearanceUpdate) ClearArtist() *ReleaseAppearanceUpdate {
	rau.mutation.ClearArtist()
	return rau
}

// ClearRelease clears the "release" edge to the Release entity.
func (rau *ReleaseAppearanceUpdate) ClearRelease() *ReleaseAppearanceUpdate {
	rau.mutation.ClearRelease()
	return rau
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rau *ReleaseAppearanceUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, rau.sqlSave, rau.mutation, rau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rau *ReleaseAppearanceUpdate) SaveX(ctx context.Context) int {
	affected, err := rau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rau *ReleaseAppearanceUpdate) Exec(ctx context.Context) error {
	_, err := rau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rau *ReleaseAppearanceUpdate) ExecX(ctx context.Context) {
	if err := rau.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rau *ReleaseAppearanceUpdate) check() error {
	if v, ok := rau.mutation.Order(); ok {
		if err := releaseappearance.OrderValidator(v); err != nil {
			return &ValidationError{Name: "order", err: fmt.Errorf(`ent: validator failed for field "ReleaseAppearance.order": %w`, err)}
		}
	}
	if rau.mutation.ArtistCleared() && len(rau.mutation.ArtistIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "ReleaseAppearance.artist"`)
	}
	if rau.mutation.ReleaseCleared() && len(rau.mutation.ReleaseIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "ReleaseAppearance.release"`)
	}
	return nil
}

func (rau *ReleaseAppearanceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := rau.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(releaseappearance.Table, releaseappearance.Columns, sqlgraph.NewFieldSpec(releaseappearance.FieldArtistID, field.TypeInt64), sqlgraph.NewFieldSpec(releaseappearance.FieldReleaseID, field.TypeInt64))
	if ps := rau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rau.mutation.Order(); ok {
		_spec.SetField(releaseappearance.FieldOrder, field.TypeInt, value)
	}
	if value, ok := rau.mutation.AddedOrder(); ok {
		_spec.AddField(releaseappearance.FieldOrder, field.TypeInt, value)
	}
	if rau.mutation.ArtistCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   releaseappearance.ArtistTable,
			Columns: []string{releaseappearance.ArtistColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artist.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rau.mutation.ArtistIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   releaseappearance.ArtistTable,
			Columns: []string{releaseappearance.ArtistColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artist.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rau.mutation.ReleaseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   releaseappearance.ReleaseTable,
			Columns: []string{releaseappearance.ReleaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(release.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rau.mutation.ReleaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   releaseappearance.ReleaseTable,
			Columns: []string{releaseappearance.ReleaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(release.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{releaseappearance.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rau.mutation.done = true
	return n, nil
}

// ReleaseAppearanceUpdateOne is the builder for updating a single ReleaseAppearance entity.
type ReleaseAppearanceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ReleaseAppearanceMutation
}

// SetReleaseID sets the "release_id" field.
func (rauo *ReleaseAppearanceUpdateOne) SetReleaseID(pi pid.ID) *ReleaseAppearanceUpdateOne {
	rauo.mutation.SetReleaseID(pi)
	return rauo
}

// SetNillableReleaseID sets the "release_id" field if the given value is not nil.
func (rauo *ReleaseAppearanceUpdateOne) SetNillableReleaseID(pi *pid.ID) *ReleaseAppearanceUpdateOne {
	if pi != nil {
		rauo.SetReleaseID(*pi)
	}
	return rauo
}

// SetArtistID sets the "artist_id" field.
func (rauo *ReleaseAppearanceUpdateOne) SetArtistID(pi pid.ID) *ReleaseAppearanceUpdateOne {
	rauo.mutation.SetArtistID(pi)
	return rauo
}

// SetNillableArtistID sets the "artist_id" field if the given value is not nil.
func (rauo *ReleaseAppearanceUpdateOne) SetNillableArtistID(pi *pid.ID) *ReleaseAppearanceUpdateOne {
	if pi != nil {
		rauo.SetArtistID(*pi)
	}
	return rauo
}

// SetOrder sets the "order" field.
func (rauo *ReleaseAppearanceUpdateOne) SetOrder(i int) *ReleaseAppearanceUpdateOne {
	rauo.mutation.ResetOrder()
	rauo.mutation.SetOrder(i)
	return rauo
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (rauo *ReleaseAppearanceUpdateOne) SetNillableOrder(i *int) *ReleaseAppearanceUpdateOne {
	if i != nil {
		rauo.SetOrder(*i)
	}
	return rauo
}

// AddOrder adds i to the "order" field.
func (rauo *ReleaseAppearanceUpdateOne) AddOrder(i int) *ReleaseAppearanceUpdateOne {
	rauo.mutation.AddOrder(i)
	return rauo
}

// SetArtist sets the "artist" edge to the Artist entity.
func (rauo *ReleaseAppearanceUpdateOne) SetArtist(a *Artist) *ReleaseAppearanceUpdateOne {
	return rauo.SetArtistID(a.ID)
}

// SetRelease sets the "release" edge to the Release entity.
func (rauo *ReleaseAppearanceUpdateOne) SetRelease(r *Release) *ReleaseAppearanceUpdateOne {
	return rauo.SetReleaseID(r.ID)
}

// Mutation returns the ReleaseAppearanceMutation object of the builder.
func (rauo *ReleaseAppearanceUpdateOne) Mutation() *ReleaseAppearanceMutation {
	return rauo.mutation
}

// ClearArtist clears the "artist" edge to the Artist entity.
func (rauo *ReleaseAppearanceUpdateOne) ClearArtist() *ReleaseAppearanceUpdateOne {
	rauo.mutation.ClearArtist()
	return rauo
}

// ClearRelease clears the "release" edge to the Release entity.
func (rauo *ReleaseAppearanceUpdateOne) ClearRelease() *ReleaseAppearanceUpdateOne {
	rauo.mutation.ClearRelease()
	return rauo
}

// Where appends a list predicates to the ReleaseAppearanceUpdate builder.
func (rauo *ReleaseAppearanceUpdateOne) Where(ps ...predicate.ReleaseAppearance) *ReleaseAppearanceUpdateOne {
	rauo.mutation.Where(ps...)
	return rauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rauo *ReleaseAppearanceUpdateOne) Select(field string, fields ...string) *ReleaseAppearanceUpdateOne {
	rauo.fields = append([]string{field}, fields...)
	return rauo
}

// Save executes the query and returns the updated ReleaseAppearance entity.
func (rauo *ReleaseAppearanceUpdateOne) Save(ctx context.Context) (*ReleaseAppearance, error) {
	return withHooks(ctx, rauo.sqlSave, rauo.mutation, rauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rauo *ReleaseAppearanceUpdateOne) SaveX(ctx context.Context) *ReleaseAppearance {
	node, err := rauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rauo *ReleaseAppearanceUpdateOne) Exec(ctx context.Context) error {
	_, err := rauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rauo *ReleaseAppearanceUpdateOne) ExecX(ctx context.Context) {
	if err := rauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rauo *ReleaseAppearanceUpdateOne) check() error {
	if v, ok := rauo.mutation.Order(); ok {
		if err := releaseappearance.OrderValidator(v); err != nil {
			return &ValidationError{Name: "order", err: fmt.Errorf(`ent: validator failed for field "ReleaseAppearance.order": %w`, err)}
		}
	}
	if rauo.mutation.ArtistCleared() && len(rauo.mutation.ArtistIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "ReleaseAppearance.artist"`)
	}
	if rauo.mutation.ReleaseCleared() && len(rauo.mutation.ReleaseIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "ReleaseAppearance.release"`)
	}
	return nil
}

func (rauo *ReleaseAppearanceUpdateOne) sqlSave(ctx context.Context) (_node *ReleaseAppearance, err error) {
	if err := rauo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(releaseappearance.Table, releaseappearance.Columns, sqlgraph.NewFieldSpec(releaseappearance.FieldArtistID, field.TypeInt64), sqlgraph.NewFieldSpec(releaseappearance.FieldReleaseID, field.TypeInt64))
	if id, ok := rauo.mutation.ArtistID(); !ok {
		return nil, &ValidationError{Name: "artist_id", err: errors.New(`ent: missing "ReleaseAppearance.artist_id" for update`)}
	} else {
		_spec.Node.CompositeID[0].Value = id
	}
	if id, ok := rauo.mutation.ReleaseID(); !ok {
		return nil, &ValidationError{Name: "release_id", err: errors.New(`ent: missing "ReleaseAppearance.release_id" for update`)}
	} else {
		_spec.Node.CompositeID[1].Value = id
	}
	if fields := rauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, len(fields))
		for i, f := range fields {
			if !releaseappearance.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			_spec.Node.Columns[i] = f
		}
	}
	if ps := rauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rauo.mutation.Order(); ok {
		_spec.SetField(releaseappearance.FieldOrder, field.TypeInt, value)
	}
	if value, ok := rauo.mutation.AddedOrder(); ok {
		_spec.AddField(releaseappearance.FieldOrder, field.TypeInt, value)
	}
	if rauo.mutation.ArtistCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   releaseappearance.ArtistTable,
			Columns: []string{releaseappearance.ArtistColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artist.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rauo.mutation.ArtistIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   releaseappearance.ArtistTable,
			Columns: []string{releaseappearance.ArtistColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artist.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rauo.mutation.ReleaseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   releaseappearance.ReleaseTable,
			Columns: []string{releaseappearance.ReleaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(release.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rauo.mutation.ReleaseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   releaseappearance.ReleaseTable,
			Columns: []string{releaseappearance.ReleaseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(release.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ReleaseAppearance{config: rauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{releaseappearance.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rauo.mutation.done = true
	return _node, nil
}
