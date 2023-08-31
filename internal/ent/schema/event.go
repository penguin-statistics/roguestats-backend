package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"exusiai.dev/roguestats-backend/internal/x/entid"
)

// Event holds the schema definition for the Event entity.
type Event struct {
	ent.Schema
}

// Fields of the Event.
func (Event) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			StorageKey("event_id").
			Unique().
			Immutable().
			DefaultFunc(entid.NewGenerator("evt")).
			Annotations(
				entgql.OrderField("ID"),
			),
		field.Time("created_at").
			Immutable().
			Default(time.Now).
			Annotations(
				entgql.OrderField("CREATED_AT"),
			),
		field.String("user_id").
			Annotations(
				entgql.Skip(entgql.SkipWhereInput),
			),
		field.String("research_id").
			Annotations(
				entgql.Skip(entgql.SkipWhereInput),
			),
		field.String("user_agent").
			Annotations(
				entgql.Skip(entgql.SkipWhereInput),
			),
		field.JSON("content", map[string]any{}),
	}
}

// Edges of the Event.
func (Event) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("events").
			Field("user_id").
			Required().
			Unique(),
		edge.From("research", Research.Type).
			Ref("events").
			Field("research_id").
			Required().
			Unique(),
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

// Annotations of the Event.
func (Event) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		// entgql.Mutations(entgql.MutationCreate()),
	}
}
