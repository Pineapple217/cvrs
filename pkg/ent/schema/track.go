package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Track holds the schema definition for the Track entity.
type Track struct {
	ent.Schema
}

// Fields of the Track.
func (Track) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			NotEmpty(),
		field.Int("position").
			Positive(),
	}
}

// Edges of the Track.
func (Track) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("appearing_artists", Artist.Type).
			Ref("appearing_tracks").
			Through("appearance", TrackAppearance.Type),
		edge.From("release", Release.Type).
			Ref("tracks").
			Unique().
			Required(),
	}
}

func (Track) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
	}
}
