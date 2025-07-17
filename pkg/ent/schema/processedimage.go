package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProcessedImage holds the schema definition for the ProcessedImage entity.
type ProcessedImage struct {
	ent.Schema
}

// Fields of the ProcessedImage.
func (ProcessedImage) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").
			Values(
				"WEBP",
				"PNG",
				"JPG",
			),
		field.Int("dimentions").
			Range(16, 3_000),
		field.Uint32("size_bits"),
		field.Bytes("thumb").
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

// Edges of the ProcessedImage.
func (ProcessedImage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("source", Image.Type).
			Ref("proccesed_image").
			Required().
			Unique(),
	}
}

func (ProcessedImage) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
	}
}
