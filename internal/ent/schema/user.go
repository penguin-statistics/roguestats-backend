package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	"exusiai.dev/roguestats-backend/internal/ent/schema/directives"
	"exusiai.dev/roguestats-backend/internal/x/entid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			StorageKey("user_id").
			Unique().
			DefaultFunc(entid.NewGenerator("usr")).
			Immutable(),
		field.String("name").
			MaxLen(64).
			Annotations(
				entgql.Skip(entgql.SkipWhereInput),
			),
		field.String("email").
			Unique().
			Immutable().
			Annotations(
				entgql.Skip(entgql.SkipWhereInput),
				entgql.Directives(directives.Private("id")),
			),
		field.String("credential").
			MaxLen(64).
			Annotations(
				entgql.Skip(entgql.SkipAll),
			),
		field.JSON("attributes", map[string]any{}).
			Optional().
			Annotations(
				entgql.Directives(directives.Private("id")),
			),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("events", Event.Type).
			// Required().
			Annotations(
				entsql.OnDelete(entsql.NoAction),
				entgql.Skip(entgql.SkipType),
			),
	}
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// entgql.Mutations(entgql.MutationCreate()),
	}
}
