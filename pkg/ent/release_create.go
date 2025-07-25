// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Pineapple217/cvrs/pkg/ent/artist"
	"github.com/Pineapple217/cvrs/pkg/ent/image"
	"github.com/Pineapple217/cvrs/pkg/ent/release"
	"github.com/Pineapple217/cvrs/pkg/ent/track"
	"github.com/Pineapple217/cvrs/pkg/pid"
)

// ReleaseCreate is the builder for creating a Release entity.
type ReleaseCreate struct {
	config
	mutation *ReleaseMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (rc *ReleaseCreate) SetName(s string) *ReleaseCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetType sets the "type" field.
func (rc *ReleaseCreate) SetType(r release.Type) *ReleaseCreate {
	rc.mutation.SetType(r)
	return rc
}

// SetReleaseDate sets the "release_date" field.
func (rc *ReleaseCreate) SetReleaseDate(t time.Time) *ReleaseCreate {
	rc.mutation.SetReleaseDate(t)
	return rc
}

// SetID sets the "id" field.
func (rc *ReleaseCreate) SetID(pi pid.ID) *ReleaseCreate {
	rc.mutation.SetID(pi)
	return rc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (rc *ReleaseCreate) SetNillableID(pi *pid.ID) *ReleaseCreate {
	if pi != nil {
		rc.SetID(*pi)
	}
	return rc
}

// SetImageID sets the "image" edge to the Image entity by ID.
func (rc *ReleaseCreate) SetImageID(id pid.ID) *ReleaseCreate {
	rc.mutation.SetImageID(id)
	return rc
}

// SetNillableImageID sets the "image" edge to the Image entity by ID if the given value is not nil.
func (rc *ReleaseCreate) SetNillableImageID(id *pid.ID) *ReleaseCreate {
	if id != nil {
		rc = rc.SetImageID(*id)
	}
	return rc
}

// SetImage sets the "image" edge to the Image entity.
func (rc *ReleaseCreate) SetImage(i *Image) *ReleaseCreate {
	return rc.SetImageID(i.ID)
}

// AddTrackIDs adds the "tracks" edge to the Track entity by IDs.
func (rc *ReleaseCreate) AddTrackIDs(ids ...pid.ID) *ReleaseCreate {
	rc.mutation.AddTrackIDs(ids...)
	return rc
}

// AddTracks adds the "tracks" edges to the Track entity.
func (rc *ReleaseCreate) AddTracks(t ...*Track) *ReleaseCreate {
	ids := make([]pid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return rc.AddTrackIDs(ids...)
}

// AddAppearingArtistIDs adds the "appearing_artists" edge to the Artist entity by IDs.
func (rc *ReleaseCreate) AddAppearingArtistIDs(ids ...pid.ID) *ReleaseCreate {
	rc.mutation.AddAppearingArtistIDs(ids...)
	return rc
}

// AddAppearingArtists adds the "appearing_artists" edges to the Artist entity.
func (rc *ReleaseCreate) AddAppearingArtists(a ...*Artist) *ReleaseCreate {
	ids := make([]pid.ID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return rc.AddAppearingArtistIDs(ids...)
}

// Mutation returns the ReleaseMutation object of the builder.
func (rc *ReleaseCreate) Mutation() *ReleaseMutation {
	return rc.mutation
}

// Save creates the Release in the database.
func (rc *ReleaseCreate) Save(ctx context.Context) (*Release, error) {
	rc.defaults()
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *ReleaseCreate) SaveX(ctx context.Context) *Release {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *ReleaseCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *ReleaseCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *ReleaseCreate) defaults() {
	if _, ok := rc.mutation.ID(); !ok {
		v := release.DefaultID()
		rc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *ReleaseCreate) check() error {
	if _, ok := rc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Release.name"`)}
	}
	if v, ok := rc.mutation.Name(); ok {
		if err := release.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Release.name": %w`, err)}
		}
	}
	if _, ok := rc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Release.type"`)}
	}
	if v, ok := rc.mutation.GetType(); ok {
		if err := release.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Release.type": %w`, err)}
		}
	}
	if _, ok := rc.mutation.ReleaseDate(); !ok {
		return &ValidationError{Name: "release_date", err: errors.New(`ent: missing required field "Release.release_date"`)}
	}
	return nil
}

func (rc *ReleaseCreate) sqlSave(ctx context.Context) (*Release, error) {
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
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = pid.ID(id)
	}
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *ReleaseCreate) createSpec() (*Release, *sqlgraph.CreateSpec) {
	var (
		_node = &Release{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(release.Table, sqlgraph.NewFieldSpec(release.FieldID, field.TypeInt64))
	)
	if id, ok := rc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rc.mutation.Name(); ok {
		_spec.SetField(release.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := rc.mutation.GetType(); ok {
		_spec.SetField(release.FieldType, field.TypeEnum, value)
		_node.Type = value
	}
	if value, ok := rc.mutation.ReleaseDate(); ok {
		_spec.SetField(release.FieldReleaseDate, field.TypeTime, value)
		_node.ReleaseDate = value
	}
	if nodes := rc.mutation.ImageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   release.ImageTable,
			Columns: []string{release.ImageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.TracksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   release.TracksTable,
			Columns: []string{release.TracksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(track.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.AppearingArtistsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   release.AppearingArtistsTable,
			Columns: release.AppearingArtistsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(artist.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &ReleaseAppearanceCreate{config: rc.config, mutation: newReleaseAppearanceMutation(rc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ReleaseCreateBulk is the builder for creating many Release entities in bulk.
type ReleaseCreateBulk struct {
	config
	err      error
	builders []*ReleaseCreate
}

// Save creates the Release entities in the database.
func (rcb *ReleaseCreateBulk) Save(ctx context.Context) ([]*Release, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Release, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ReleaseMutation)
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
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = pid.ID(id)
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
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *ReleaseCreateBulk) SaveX(ctx context.Context) []*Release {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *ReleaseCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *ReleaseCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
