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

// Metric holds the schema definition for the Metric entity.
type Metric struct {
	ent.Schema
}

// Fields of the Metric.
func (Metric) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			StorageKey("Metric_id").
			Unique().
			DefaultFunc(entid.NewGenerator("rsc")).
			Immutable().
			Annotations(
				entgql.OrderField("ID"),
			),
		field.String("name").MaxLen(64),
		field.JSON("filter", map[string]any{}).
			Comment("The jsonpb filter to apply to the events"),
		field.String("mapping").
			Comment("The mapping expr to apply to the events"),
	}
}

// Edges of the Metric.
func (Metric) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("events", Event.Type).
			// Required().
			Annotations(
				entsql.OnDelete(entsql.NoAction),
				entgql.Skip(entgql.SkipType),
			),
	}
}

// Annotations of the Metric.
func (Metric) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
