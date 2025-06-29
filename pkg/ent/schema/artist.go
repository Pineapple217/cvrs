package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Artist holds the schema definition for the Artist entity.
type Artist struct {
	ent.Schema
}

// Fields of the Artist.
func (Artist) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("deleted_at").
			Optional().
			Nillable(),
	}
}

// Edges of the Artist.
func (Artist) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("appearing_tracks", Track.Type).
			Through("track_appearance", TrackAppearance.Type),
		edge.To("appearing_releases", Release.Type).
			Through("release_appearance", ReleaseAppearance.Type),
		edge.To("image", Image.Type).
			Unique(),
	}
}

func (Artist) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
	}
}
