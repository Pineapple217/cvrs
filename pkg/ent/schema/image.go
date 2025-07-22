package schema

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	gen "github.com/Pineapple217/cvrs/pkg/ent"
	"github.com/Pineapple217/cvrs/pkg/ent/hook"
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
		field.String("original_name").
			NotEmpty(),
		field.Enum("type").
			Values(
				"WEBP",
				"PNG",
				"JPG",
			),
		field.String("note").
			Optional().
			Nillable(),
		field.Int("dimention_width").
			Range(16, 10_000),
		field.Int("dimention_height").
			Range(16, 10_000),
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
			Unique(),
		edge.From("artist", Artist.Type).
			Ref("image").
			Unique(),
		edge.From("uploader", User.Type).
			Ref("images").
			Unique().
			Required(),
		edge.To("proccesed_image", ProcessedImage.Type),
		edge.To("data", ImageData.Type).
			Unique(),
	}
}

func (Image) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(func(next ent.Mutator) ent.Mutator {
			return hook.ImageFunc(func(ctx context.Context, m *gen.ImageMutation) (ent.Value, error) {
				_, hasRelease := m.ReleaseID()
				_, hasArtist := m.ArtistID()
				if hasRelease == hasArtist {
					if hasRelease {
						return nil, fmt.Errorf("an image must be linked to either a release or an artist, but both were provided")
					} else {
						// return nil, fmt.Errorf("an image must be linked to either a release or an artist, but neither was provided")
					}
				}
				return next.Mutate(ctx, m)
			})
		},
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),
	}
}

func (Image) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
	}
}
