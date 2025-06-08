package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Pineapple217/cvrs/pkg/pid"
)

// ReleaseAppearance holds the schema definition for the ReleaseAppearance entity.
type ReleaseAppearance struct {
	ent.Schema
}

func (ReleaseAppearance) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("artist_id", "release_id"),
	}
}

// Fields of the ReleaseAppearances.
func (ReleaseAppearance) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("release_id").
			GoType(pid.New()).
			DefaultFunc(pid.New),
		field.Int64("artist_id").
			GoType(pid.New()).
			DefaultFunc(pid.New),
		field.Int("order").
			NonNegative(),
	}
}

// Edges of the ReleaseAppearances.
func (ReleaseAppearance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("artist", Artist.Type).
			Unique().
			Required().
			Field("artist_id"),
		edge.To("release", Release.Type).
			Unique().
			Required().
			Field("release_id"),
	}
}
