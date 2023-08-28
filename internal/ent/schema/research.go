package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"exusiai.dev/roguestats-backend/internal/x/entid"
)

// Research holds the schema definition for the Research entity.
type Research struct {
	ent.Schema
}

// Fields of the Research.
func (Research) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			StorageKey("research_id").
			Unique().
			DefaultFunc(entid.NewGenerator("rsc")).
			Immutable().
			Annotations(
				entgql.OrderField("ID"),
			),
		field.String("name").MaxLen(64),
		field.JSON("schema", map[string]any{}),
	}
}

// Edges of the Research.
func (Research) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("events", Event.Type).
			Required().
			Annotations(entsql.OnDelete(entsql.NoAction)),
	}
}

// Annotations of the Research.
func (Research) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
