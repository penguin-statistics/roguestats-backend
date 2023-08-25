package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").StorageKey("event_id").Unique().Immutable(),
		field.String("created_at").Immutable(),
		field.String("user_agent"),
		field.JSON("content", map[string]any{}),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("events").Unique(),
		edge.From("research", Research.Type).Ref("events").Unique(),
	}
}

// Indexes of the Event.
func (Event) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
		index.Fields("user_agent"),
		index.Edges("user"),
		index.Edges("research"),
	}
}
