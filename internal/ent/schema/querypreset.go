package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"exusiai.dev/roguestats-backend/internal/x/entid"
)

// QueryPreset holds the schema definition for the QueryPreset entity.
type QueryPreset struct {
	ent.Schema
}

// Fields of the QueryPreset.
func (QueryPreset) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			StorageKey("query_preset_id").
			Unique().
			DefaultFunc(entid.NewGenerator("qps")).
			Immutable().
			Annotations(
				entgql.OrderField("ID"),
			),
		field.String("name").MaxLen(64),
		field.String("research_id").
			Annotations(
				entgql.Skip(entgql.SkipWhereInput),
			),
		field.JSON("where", map[string]any{}).
			Comment("The filter to apply to the events"),
		field.String("mapping").
			Comment("The mapping expr to apply to the events"),
	}
}

// Edges of the QueryPreset.
func (QueryPreset) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("research", Research.Type).
			Ref("query_presets").
			Field("research_id").
			Required().
			Unique(),
	}
}

// Annotations of the QueryPreset.
func (QueryPreset) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
