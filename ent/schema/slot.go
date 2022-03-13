package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Slot holds the schema definition for the Slot entity.
type Slot struct {
	ent.Schema
}

// Fields of the Slot.
func (Slot) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("slotId", uuid.New()).Unique().Immutable(),
		field.String("patientName").NotEmpty(),
		field.String("patientId").NotEmpty(),
		field.String("attendingName").NotEmpty(),
		field.String("attendingId").NotEmpty(),
		field.Time("startTime"),
		field.Time("endTime"),
		field.String("status").NotEmpty(),
	}
}

// Edges of the Slot.
func (Slot) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("AttendingDoctor", User.Type).
			Ref("Attending").
			Unique(),
		edge.From("patient", User.Type).
			Ref("Visiting").
			Unique(),
	}
}
