package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Image holds the schema definition for the Image entity.
type Image struct {
	ent.Schema
}

// Fields of the Img.
func (Image) Fields() []ent.Field {
	return []ent.Field{
		field.String("file").
			NotEmpty(),
		field.Enum("type").
			Values(
				"webp",
				"png",
				"jpg",
			),
		field.String("note").
			Optional().
			Nillable(),
		field.Ints("dimentions"),
		field.Uint32("size_bits"),
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

// Edges of the Img.
func (Image) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("release", Release.Type).
			Ref("image").
			Unique().
			Required(),
		edge.From("uploader", User.Type).
			Ref("images").
			Unique().
			Required(),
	}
}

func (Image) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
	}
}
