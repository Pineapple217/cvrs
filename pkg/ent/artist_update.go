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
	"github.com/Pineapple217/cvrs/pkg/ent/artist"
	"github.com/Pineapple217/cvrs/pkg/ent/image"
	"github.com/Pineapple217/cvrs/pkg/ent/predicate"
	"github.com/Pineapple217/cvrs/pkg/ent/release"
	"github.com/Pineapple217/cvrs/pkg/ent/track"
	"github.com/Pineapple217/cvrs/pkg/pid"
)

// ArtistUpdate is the builder for updating Artist entities.
type ArtistUpdate struct {
	config
	hooks    []Hook
	mutation *ArtistMutation
}

// Where appends a list predicates to the ArtistUpdate builder.
func (au *ArtistUpdate) Where(ps ...predicate.Artist) *ArtistUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetName sets the "name" field.
func (au *ArtistUpdate) SetName(s string) *ArtistUpdate {
	au.mutation.SetName(s)
	return au
}

// SetNillableName sets the "name" field if the given value is not nil.
func (au *ArtistUpdate) SetNillableName(s *string) *ArtistUpdate {
	if s != nil {
		au.SetName(*s)
	}
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *ArtistUpdate) SetUpdatedAt(t time.Time) *ArtistUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetDeletedAt sets the "deleted_at" field.
func (au *ArtistUpdate) SetDeletedAt(t time.Time) *ArtistUpdate {
	au.mutation.SetDeletedAt(t)
	return au
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (au *ArtistUpdate) SetNillableDeletedAt(t *time.Time) *ArtistUpdate {
	if t != nil {
		au.SetDeletedAt(*t)
	}
	return au
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (au *ArtistUpdate) ClearDeletedAt() *ArtistUpdate {
	au.mutation.ClearDeletedAt()
	return au
}

// AddAppearingTrackIDs adds the "appearing_tracks" edge to the Track entity by IDs.
func (au *ArtistUpdate) AddAppearingTrackIDs(ids ...pid.ID) *ArtistUpdate {
	au.mutation.AddAppearingTrackIDs(ids...)
	return au
}

// AddAppearingTracks adds the "appearing_tracks" edges to the Track entity.
func (au *ArtistUpdate) AddAppearingTracks(t ...*Track) *ArtistUpdate {
	ids := make([]pid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.AddAppearingTrackIDs(ids...)
}

// AddAppearingReleaseIDs adds the "appearing_releases" edge to the Release entity by IDs.
func (au *ArtistUpdate) AddAppearingReleaseIDs(ids ...pid.ID) *ArtistUpdate {
	au.mutation.AddAppearingReleaseIDs(ids...)
	return au
}

// AddAppearingReleases adds the "appearing_releases" edges to the Release entity.
func (au *ArtistUpdate) AddAppearingReleases(r ...*Release) *ArtistUpdate {
	ids := make([]pid.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return au.AddAppearingReleaseIDs(ids...)
}

// SetImageID sets the "image" edge to the Image entity by ID.
func (au *ArtistUpdate) SetImageID(id pid.ID) *ArtistUpdate {
	au.mutation.SetImageID(id)
	return au
}

// SetNillableImageID sets the "image" edge to the Image entity by ID if the given value is not nil.
func (au *ArtistUpdate) SetNillableImageID(id *pid.ID) *ArtistUpdate {
	if id != nil {
		au = au.SetImageID(*id)
	}
	return au
}

// SetImage sets the "image" edge to the Image entity.
func (au *ArtistUpdate) SetImage(i *Image) *ArtistUpdate {
	return au.SetImageID(i.ID)
}

// Mutation returns the ArtistMutation object of the builder.
func (au *ArtistUpdate) Mutation() *ArtistMutation {
	return au.mutation
}

// ClearAppearingTracks clears all "appearing_tracks" edges to the Track entity.
func (au *ArtistUpdate) ClearAppearingTracks() *ArtistUpdate {
	au.mutation.ClearAppearingTracks()
	return au
}

// RemoveAppearingTrackIDs removes the "appearing_tracks" edge to Track entities by IDs.
func (au *ArtistUpdate) RemoveAppearingTrackIDs(ids ...pid.ID) *ArtistUpdate {
	au.mutation.RemoveAppearingTrackIDs(ids...)
	return au
}

// RemoveAppearingTracks removes "appearing_tracks" edges to Track entities.
func (au *ArtistUpdate) RemoveAppearingTracks(t ...*Track) *ArtistUpdate {
	ids := make([]pid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return au.RemoveAppearingTrackIDs(ids...)
}

// ClearAppearingReleases clears all "appearing_releases" edges to the Release entity.
func (au *ArtistUpdate) ClearAppearingReleases() *ArtistUpdate {
	au.mutation.ClearAppearingReleases()
	return au
}

// RemoveAppearingReleaseIDs removes the "appearing_releases" edge to Release entities by IDs.
func (au *ArtistUpdate) RemoveAppearingReleaseIDs(ids ...pid.ID) *ArtistUpdate {
	au.mutation.RemoveAppearingReleaseIDs(ids...)
	return au
}

// RemoveAppearingReleases removes "appearing_releases" edges to Release entities.
func (au *ArtistUpdate) RemoveAppearingReleases(r ...*Release) *ArtistUpdate {
	ids := make([]pid.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return au.RemoveAppearingReleaseIDs(ids...)
}

// ClearImage clears the "image" edge to the Image entity.
func (au *ArtistUpdate) ClearImage() *ArtistUpdate {
	au.mutation.ClearImage()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *ArtistUpdate) Save(ctx context.Context) (int, error) {
	au.defaults()
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *ArtistUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *ArtistUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *ArtistUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *ArtistUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := artist.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *ArtistUpdate) check() error {
	if v, ok := au.mutation.Name(); ok {
		if err := artist.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Artist.name": %w`, err)}
		}
	}
	return nil
}

func (au *ArtistUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(artist.Table, artist.Columns, sqlgraph.NewFieldSpec(artist.FieldID, field.TypeInt64))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(artist.FieldName, field.TypeString, value)
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(artist.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.DeletedAt(); ok {
		_spec.SetField(artist.FieldDeletedAt, field.TypeTime, value)
	}
	if au.mutation.DeletedAtCleared() {
		_spec.ClearField(artist.FieldDeletedAt, field.TypeTime)
	}
	if au.mutation.AppearingTracksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artist.AppearingTracksTable,
			Columns: artist.AppearingTracksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(track.FieldID, field.TypeInt64),
			},
		}
		createE := &TrackAppearanceCreate{config: au.config, mutation: newTrackAppearanceMutation(au.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedAppearingTracksIDs(); len(nodes) > 0 && !au.mutation.AppearingTracksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artist.AppearingTracksTable,
			Columns: artist.AppearingTracksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(track.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TrackAppearanceCreate{config: au.config, mutation: newTrackAppearanceMutation(au.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.AppearingTracksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artist.AppearingTracksTable,
			Columns: artist.AppearingTracksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(track.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TrackAppearanceCreate{config: au.config, mutation: newTrackAppearanceMutation(au.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.AppearingReleasesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artist.AppearingReleasesTable,
			Columns: artist.AppearingReleasesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(release.FieldID, field.TypeInt64),
			},
		}
		createE := &ReleaseAppearanceCreate{config: au.config, mutation: newReleaseAppearanceMutation(au.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedAppearingReleasesIDs(); len(nodes) > 0 && !au.mutation.AppearingReleasesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artist.AppearingReleasesTable,
			Columns: artist.AppearingReleasesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(release.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &ReleaseAppearanceCreate{config: au.config, mutation: newReleaseAppearanceMutation(au.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.AppearingReleasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artist.AppearingReleasesTable,
			Columns: artist.AppearingReleasesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(release.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &ReleaseAppearanceCreate{config: au.config, mutation: newReleaseAppearanceMutation(au.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.ImageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   artist.ImageTable,
			Columns: []string{artist.ImageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.ImageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   artist.ImageTable,
			Columns: []string{artist.ImageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{artist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// ArtistUpdateOne is the builder for updating a single Artist entity.
type ArtistUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ArtistMutation
}

// SetName sets the "name" field.
func (auo *ArtistUpdateOne) SetName(s string) *ArtistUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (auo *ArtistUpdateOne) SetNillableName(s *string) *ArtistUpdateOne {
	if s != nil {
		auo.SetName(*s)
	}
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *ArtistUpdateOne) SetUpdatedAt(t time.Time) *ArtistUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetDeletedAt sets the "deleted_at" field.
func (auo *ArtistUpdateOne) SetDeletedAt(t time.Time) *ArtistUpdateOne {
	auo.mutation.SetDeletedAt(t)
	return auo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (auo *ArtistUpdateOne) SetNillableDeletedAt(t *time.Time) *ArtistUpdateOne {
	if t != nil {
		auo.SetDeletedAt(*t)
	}
	return auo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (auo *ArtistUpdateOne) ClearDeletedAt() *ArtistUpdateOne {
	auo.mutation.ClearDeletedAt()
	return auo
}

// AddAppearingTrackIDs adds the "appearing_tracks" edge to the Track entity by IDs.
func (auo *ArtistUpdateOne) AddAppearingTrackIDs(ids ...pid.ID) *ArtistUpdateOne {
	auo.mutation.AddAppearingTrackIDs(ids...)
	return auo
}

// AddAppearingTracks adds the "appearing_tracks" edges to the Track entity.
func (auo *ArtistUpdateOne) AddAppearingTracks(t ...*Track) *ArtistUpdateOne {
	ids := make([]pid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.AddAppearingTrackIDs(ids...)
}

// AddAppearingReleaseIDs adds the "appearing_releases" edge to the Release entity by IDs.
func (auo *ArtistUpdateOne) AddAppearingReleaseIDs(ids ...pid.ID) *ArtistUpdateOne {
	auo.mutation.AddAppearingReleaseIDs(ids...)
	return auo
}

// AddAppearingReleases adds the "appearing_releases" edges to the Release entity.
func (auo *ArtistUpdateOne) AddAppearingReleases(r ...*Release) *ArtistUpdateOne {
	ids := make([]pid.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return auo.AddAppearingReleaseIDs(ids...)
}

// SetImageID sets the "image" edge to the Image entity by ID.
func (auo *ArtistUpdateOne) SetImageID(id pid.ID) *ArtistUpdateOne {
	auo.mutation.SetImageID(id)
	return auo
}

// SetNillableImageID sets the "image" edge to the Image entity by ID if the given value is not nil.
func (auo *ArtistUpdateOne) SetNillableImageID(id *pid.ID) *ArtistUpdateOne {
	if id != nil {
		auo = auo.SetImageID(*id)
	}
	return auo
}

// SetImage sets the "image" edge to the Image entity.
func (auo *ArtistUpdateOne) SetImage(i *Image) *ArtistUpdateOne {
	return auo.SetImageID(i.ID)
}

// Mutation returns the ArtistMutation object of the builder.
func (auo *ArtistUpdateOne) Mutation() *ArtistMutation {
	return auo.mutation
}

// ClearAppearingTracks clears all "appearing_tracks" edges to the Track entity.
func (auo *ArtistUpdateOne) ClearAppearingTracks() *ArtistUpdateOne {
	auo.mutation.ClearAppearingTracks()
	return auo
}

// RemoveAppearingTrackIDs removes the "appearing_tracks" edge to Track entities by IDs.
func (auo *ArtistUpdateOne) RemoveAppearingTrackIDs(ids ...pid.ID) *ArtistUpdateOne {
	auo.mutation.RemoveAppearingTrackIDs(ids...)
	return auo
}

// RemoveAppearingTracks removes "appearing_tracks" edges to Track entities.
func (auo *ArtistUpdateOne) RemoveAppearingTracks(t ...*Track) *ArtistUpdateOne {
	ids := make([]pid.ID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return auo.RemoveAppearingTrackIDs(ids...)
}

// ClearAppearingReleases clears all "appearing_releases" edges to the Release entity.
func (auo *ArtistUpdateOne) ClearAppearingReleases() *ArtistUpdateOne {
	auo.mutation.ClearAppearingReleases()
	return auo
}

// RemoveAppearingReleaseIDs removes the "appearing_releases" edge to Release entities by IDs.
func (auo *ArtistUpdateOne) RemoveAppearingReleaseIDs(ids ...pid.ID) *ArtistUpdateOne {
	auo.mutation.RemoveAppearingReleaseIDs(ids...)
	return auo
}

// RemoveAppearingReleases removes "appearing_releases" edges to Release entities.
func (auo *ArtistUpdateOne) RemoveAppearingReleases(r ...*Release) *ArtistUpdateOne {
	ids := make([]pid.ID, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return auo.RemoveAppearingReleaseIDs(ids...)
}

// ClearImage clears the "image" edge to the Image entity.
func (auo *ArtistUpdateOne) ClearImage() *ArtistUpdateOne {
	auo.mutation.ClearImage()
	return auo
}

// Where appends a list predicates to the ArtistUpdate builder.
func (auo *ArtistUpdateOne) Where(ps ...predicate.Artist) *ArtistUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *ArtistUpdateOne) Select(field string, fields ...string) *ArtistUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Artist entity.
func (auo *ArtistUpdateOne) Save(ctx context.Context) (*Artist, error) {
	auo.defaults()
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *ArtistUpdateOne) SaveX(ctx context.Context) *Artist {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *ArtistUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *ArtistUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *ArtistUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := artist.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *ArtistUpdateOne) check() error {
	if v, ok := auo.mutation.Name(); ok {
		if err := artist.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Artist.name": %w`, err)}
		}
	}
	return nil
}

func (auo *ArtistUpdateOne) sqlSave(ctx context.Context) (_node *Artist, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(artist.Table, artist.Columns, sqlgraph.NewFieldSpec(artist.FieldID, field.TypeInt64))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Artist.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, artist.FieldID)
		for _, f := range fields {
			if !artist.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != artist.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(artist.FieldName, field.TypeString, value)
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(artist.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.DeletedAt(); ok {
		_spec.SetField(artist.FieldDeletedAt, field.TypeTime, value)
	}
	if auo.mutation.DeletedAtCleared() {
		_spec.ClearField(artist.FieldDeletedAt, field.TypeTime)
	}
	if auo.mutation.AppearingTracksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artist.AppearingTracksTable,
			Columns: artist.AppearingTracksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(track.FieldID, field.TypeInt64),
			},
		}
		createE := &TrackAppearanceCreate{config: auo.config, mutation: newTrackAppearanceMutation(auo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedAppearingTracksIDs(); len(nodes) > 0 && !auo.mutation.AppearingTracksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artist.AppearingTracksTable,
			Columns: artist.AppearingTracksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(track.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TrackAppearanceCreate{config: auo.config, mutation: newTrackAppearanceMutation(auo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.AppearingTracksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artist.AppearingTracksTable,
			Columns: artist.AppearingTracksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(track.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TrackAppearanceCreate{config: auo.config, mutation: newTrackAppearanceMutation(auo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.AppearingReleasesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artist.AppearingReleasesTable,
			Columns: artist.AppearingReleasesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(release.FieldID, field.TypeInt64),
			},
		}
		createE := &ReleaseAppearanceCreate{config: auo.config, mutation: newReleaseAppearanceMutation(auo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedAppearingReleasesIDs(); len(nodes) > 0 && !auo.mutation.AppearingReleasesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artist.AppearingReleasesTable,
			Columns: artist.AppearingReleasesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(release.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &ReleaseAppearanceCreate{config: auo.config, mutation: newReleaseAppearanceMutation(auo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.AppearingReleasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   artist.AppearingReleasesTable,
			Columns: artist.AppearingReleasesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(release.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &ReleaseAppearanceCreate{config: auo.config, mutation: newReleaseAppearanceMutation(auo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.ImageCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   artist.ImageTable,
			Columns: []string{artist.ImageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeInt64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.ImageIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   artist.ImageTable,
			Columns: []string{artist.ImageColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(image.FieldID, field.TypeInt64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Artist{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{artist.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
