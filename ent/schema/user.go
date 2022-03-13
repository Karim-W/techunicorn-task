package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("userId", uuid.New()).Unique().Immutable(),
		field.String("FirstName").Default(""),
		field.String("LastName").Default(""),
		field.String("Email").Unique(),
		field.String("Password").NotEmpty(),
		field.String("Phone").Default(""),
		field.String("Type").NotEmpty(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Attending", Slot.Type),
		edge.To("Visiting", Slot.Type),
	}
}
