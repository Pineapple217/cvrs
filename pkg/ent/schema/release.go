package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Release holds the schema definition for the Release entity.
type Release struct {
	ent.Schema
}

// Fields of the Release.
func (Release) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
		field.Enum("type").
			Values(
				"album",
				"single",
				"EP",
				"compilation",
				"unknown",
			),
		field.Time("release_date"),
	}
}

// Edges of the Release.
func (Release) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("image", Image.Type).
			Unique(),
		edge.To("tracks", Track.Type),
		edge.From("appearing_artists", Artist.Type).
			Ref("appearing_releases").
			Through("release_appearance", ReleaseAppearance.Type),
	}
}

func (Release) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
	}
}
