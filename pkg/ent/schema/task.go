package schema

import (
	"encoding/json"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("type").
			Values(
				"scale_img",
			),
		field.Enum("status").
			Values(
				"pending",
				"working",
				"error",
				"done",
			).Default("pending"),
		field.String("error").
			Optional(),
		field.JSON("payload", json.RawMessage{}),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return nil
}

func (Task) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
	}
}
