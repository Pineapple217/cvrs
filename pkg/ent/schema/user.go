package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("username").
			MinLen(3).
			MaxLen(32).
			Unique(),
		field.Bytes("password").
			NotEmpty(),
		field.Bool("is_admin").
			Default(false),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("images", Image.Type),
	}
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
	}
}
