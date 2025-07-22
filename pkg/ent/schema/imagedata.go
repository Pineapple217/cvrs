package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ImageData holds the schema definition for the ImageData entity.
type ImageData struct {
	ent.Schema
}

// Fields of the ImageData.
func (ImageData) Fields() []ent.Field {
	return []ent.Field{
		field.Int("avr_r"),
		field.Int("avr_g"),
		field.Int("avr_b"),
		field.Int("avg_brightness"),
		field.Int("avg_saturation"),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the ImageData.
func (ImageData) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("image", Image.Type).
			Ref("data").
			Required(),
	}
}
