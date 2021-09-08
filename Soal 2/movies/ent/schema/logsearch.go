package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// LogSearch holds the schema definition for the LogSearch entity.
type LogSearch struct {
	ent.Schema
}

// Fields of the LogSearch.
func (LogSearch) Fields() []ent.Field {
	return []ent.Field{
		field.String("keyword"),
		field.String("page"),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the LogSearch.
func (LogSearch) Edges() []ent.Edge {
	return nil
}
