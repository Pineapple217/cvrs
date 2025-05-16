package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"github.com/Pineapple217/cvrs/pkg/pid"
)

type IDMixin struct {
	mixin.Schema
}

func (IDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			GoType(pid.New()).
			DefaultFunc(pid.New).
			Immutable().
			Unique(),
	}
}

func (IDMixin) Edges() []ent.Edge { return nil }
