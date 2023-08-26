package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

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
		field.String("name").MaxLen(64),
		field.String("email").Unique().Immutable(),
		field.String("credential").MaxLen(64),
		field.JSON("attributes", map[string]any{}).Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("events", Event.Type).
			Required().
			StorageKey(edge.Column("user_id")).
			Annotations(entsql.OnDelete(entsql.NoAction)),
	}
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// entgql.Mutations(entgql.MutationCreate()),
	}
}
