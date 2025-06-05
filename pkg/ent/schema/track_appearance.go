package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/Pineapple217/cvrs/pkg/pid"
)

// TrackAppearance holds the schema definition for the TrackAppearance entity.
type TrackAppearance struct {
	ent.Schema
}

func (TrackAppearance) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("artist_id", "track_id"),
	}
}

func (TrackAppearance) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("track_id").
			GoType(pid.New()).
			DefaultFunc(pid.New),
		field.Int64("artist_id").
			GoType(pid.New()).
			DefaultFunc(pid.New),
		field.Int("order").
			Positive(),
	}
}

func (TrackAppearance) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("artist", Artist.Type).
			Unique().
			Required().
			Field("artist_id"),
		edge.To("track", Track.Type).
			Unique().
			Required().
			Field("track_id"),
	}
}
