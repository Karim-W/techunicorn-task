// Code generated by entc, DO NOT EDIT.

package slot

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/karim-w/techunicorn-task/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// SlotId applies equality check predicate on the "slotId" field. It's identical to SlotIdEQ.
func SlotId(v uuid.UUID) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSlotId), v))
	})
}

// PatientName applies equality check predicate on the "patientName" field. It's identical to PatientNameEQ.
func PatientName(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPatientName), v))
	})
}

// PatientId applies equality check predicate on the "patientId" field. It's identical to PatientIdEQ.
func PatientId(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPatientId), v))
	})
}

// AttendingName applies equality check predicate on the "attendingName" field. It's identical to AttendingNameEQ.
func AttendingName(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAttendingName), v))
	})
}

// AttendingId applies equality check predicate on the "attendingId" field. It's identical to AttendingIdEQ.
func AttendingId(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAttendingId), v))
	})
}

// StartTime applies equality check predicate on the "startTime" field. It's identical to StartTimeEQ.
func StartTime(v time.Time) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartTime), v))
	})
}

// EndTime applies equality check predicate on the "endTime" field. It's identical to EndTimeEQ.
func EndTime(v time.Time) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndTime), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// SlotIdEQ applies the EQ predicate on the "slotId" field.
func SlotIdEQ(v uuid.UUID) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSlotId), v))
	})
}

// SlotIdNEQ applies the NEQ predicate on the "slotId" field.
func SlotIdNEQ(v uuid.UUID) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSlotId), v))
	})
}

// SlotIdIn applies the In predicate on the "slotId" field.
func SlotIdIn(vs ...uuid.UUID) predicate.Slot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Slot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSlotId), v...))
	})
}

// SlotIdNotIn applies the NotIn predicate on the "slotId" field.
func SlotIdNotIn(vs ...uuid.UUID) predicate.Slot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Slot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSlotId), v...))
	})
}

// SlotIdGT applies the GT predicate on the "slotId" field.
func SlotIdGT(v uuid.UUID) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSlotId), v))
	})
}

// SlotIdGTE applies the GTE predicate on the "slotId" field.
func SlotIdGTE(v uuid.UUID) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSlotId), v))
	})
}

// SlotIdLT applies the LT predicate on the "slotId" field.
func SlotIdLT(v uuid.UUID) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSlotId), v))
	})
}

// SlotIdLTE applies the LTE predicate on the "slotId" field.
func SlotIdLTE(v uuid.UUID) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSlotId), v))
	})
}

// PatientNameEQ applies the EQ predicate on the "patientName" field.
func PatientNameEQ(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPatientName), v))
	})
}

// PatientNameNEQ applies the NEQ predicate on the "patientName" field.
func PatientNameNEQ(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPatientName), v))
	})
}

// PatientNameIn applies the In predicate on the "patientName" field.
func PatientNameIn(vs ...string) predicate.Slot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Slot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPatientName), v...))
	})
}

// PatientNameNotIn applies the NotIn predicate on the "patientName" field.
func PatientNameNotIn(vs ...string) predicate.Slot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Slot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPatientName), v...))
	})
}

// PatientNameGT applies the GT predicate on the "patientName" field.
func PatientNameGT(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPatientName), v))
	})
}

// PatientNameGTE applies the GTE predicate on the "patientName" field.
func PatientNameGTE(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPatientName), v))
	})
}

// PatientNameLT applies the LT predicate on the "patientName" field.
func PatientNameLT(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPatientName), v))
	})
}

// PatientNameLTE applies the LTE predicate on the "patientName" field.
func PatientNameLTE(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPatientName), v))
	})
}

// PatientNameContains applies the Contains predicate on the "patientName" field.
func PatientNameContains(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPatientName), v))
	})
}

// PatientNameHasPrefix applies the HasPrefix predicate on the "patientName" field.
func PatientNameHasPrefix(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPatientName), v))
	})
}

// PatientNameHasSuffix applies the HasSuffix predicate on the "patientName" field.
func PatientNameHasSuffix(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPatientName), v))
	})
}

// PatientNameEqualFold applies the EqualFold predicate on the "patientName" field.
func PatientNameEqualFold(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPatientName), v))
	})
}

// PatientNameContainsFold applies the ContainsFold predicate on the "patientName" field.
func PatientNameContainsFold(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPatientName), v))
	})
}

// PatientIdEQ applies the EQ predicate on the "patientId" field.
func PatientIdEQ(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPatientId), v))
	})
}

// PatientIdNEQ applies the NEQ predicate on the "patientId" field.
func PatientIdNEQ(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPatientId), v))
	})
}

// PatientIdIn applies the In predicate on the "patientId" field.
func PatientIdIn(vs ...string) predicate.Slot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Slot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPatientId), v...))
	})
}

// PatientIdNotIn applies the NotIn predicate on the "patientId" field.
func PatientIdNotIn(vs ...string) predicate.Slot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Slot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPatientId), v...))
	})
}

// PatientIdGT applies the GT predicate on the "patientId" field.
func PatientIdGT(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPatientId), v))
	})
}

// PatientIdGTE applies the GTE predicate on the "patientId" field.
func PatientIdGTE(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPatientId), v))
	})
}

// PatientIdLT applies the LT predicate on the "patientId" field.
func PatientIdLT(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPatientId), v))
	})
}

// PatientIdLTE applies the LTE predicate on the "patientId" field.
func PatientIdLTE(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPatientId), v))
	})
}

// PatientIdContains applies the Contains predicate on the "patientId" field.
func PatientIdContains(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPatientId), v))
	})
}

// PatientIdHasPrefix applies the HasPrefix predicate on the "patientId" field.
func PatientIdHasPrefix(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPatientId), v))
	})
}

// PatientIdHasSuffix applies the HasSuffix predicate on the "patientId" field.
func PatientIdHasSuffix(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPatientId), v))
	})
}

// PatientIdEqualFold applies the EqualFold predicate on the "patientId" field.
func PatientIdEqualFold(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPatientId), v))
	})
}

// PatientIdContainsFold applies the ContainsFold predicate on the "patientId" field.
func PatientIdContainsFold(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPatientId), v))
	})
}

// AttendingNameEQ applies the EQ predicate on the "attendingName" field.
func AttendingNameEQ(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAttendingName), v))
	})
}

// AttendingNameNEQ applies the NEQ predicate on the "attendingName" field.
func AttendingNameNEQ(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAttendingName), v))
	})
}

// AttendingNameIn applies the In predicate on the "attendingName" field.
func AttendingNameIn(vs ...string) predicate.Slot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Slot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAttendingName), v...))
	})
}

// AttendingNameNotIn applies the NotIn predicate on the "attendingName" field.
func AttendingNameNotIn(vs ...string) predicate.Slot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Slot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAttendingName), v...))
	})
}

// AttendingNameGT applies the GT predicate on the "attendingName" field.
func AttendingNameGT(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAttendingName), v))
	})
}

// AttendingNameGTE applies the GTE predicate on the "attendingName" field.
func AttendingNameGTE(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAttendingName), v))
	})
}

// AttendingNameLT applies the LT predicate on the "attendingName" field.
func AttendingNameLT(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAttendingName), v))
	})
}

// AttendingNameLTE applies the LTE predicate on the "attendingName" field.
func AttendingNameLTE(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAttendingName), v))
	})
}

// AttendingNameContains applies the Contains predicate on the "attendingName" field.
func AttendingNameContains(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAttendingName), v))
	})
}

// AttendingNameHasPrefix applies the HasPrefix predicate on the "attendingName" field.
func AttendingNameHasPrefix(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAttendingName), v))
	})
}

// AttendingNameHasSuffix applies the HasSuffix predicate on the "attendingName" field.
func AttendingNameHasSuffix(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAttendingName), v))
	})
}

// AttendingNameEqualFold applies the EqualFold predicate on the "attendingName" field.
func AttendingNameEqualFold(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAttendingName), v))
	})
}

// AttendingNameContainsFold applies the ContainsFold predicate on the "attendingName" field.
func AttendingNameContainsFold(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAttendingName), v))
	})
}

// AttendingIdEQ applies the EQ predicate on the "attendingId" field.
func AttendingIdEQ(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAttendingId), v))
	})
}

// AttendingIdNEQ applies the NEQ predicate on the "attendingId" field.
func AttendingIdNEQ(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAttendingId), v))
	})
}

// AttendingIdIn applies the In predicate on the "attendingId" field.
func AttendingIdIn(vs ...string) predicate.Slot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Slot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAttendingId), v...))
	})
}

// AttendingIdNotIn applies the NotIn predicate on the "attendingId" field.
func AttendingIdNotIn(vs ...string) predicate.Slot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Slot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAttendingId), v...))
	})
}

// AttendingIdGT applies the GT predicate on the "attendingId" field.
func AttendingIdGT(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAttendingId), v))
	})
}

// AttendingIdGTE applies the GTE predicate on the "attendingId" field.
func AttendingIdGTE(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAttendingId), v))
	})
}

// AttendingIdLT applies the LT predicate on the "attendingId" field.
func AttendingIdLT(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAttendingId), v))
	})
}

// AttendingIdLTE applies the LTE predicate on the "attendingId" field.
func AttendingIdLTE(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAttendingId), v))
	})
}

// AttendingIdContains applies the Contains predicate on the "attendingId" field.
func AttendingIdContains(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAttendingId), v))
	})
}

// AttendingIdHasPrefix applies the HasPrefix predicate on the "attendingId" field.
func AttendingIdHasPrefix(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAttendingId), v))
	})
}

// AttendingIdHasSuffix applies the HasSuffix predicate on the "attendingId" field.
func AttendingIdHasSuffix(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAttendingId), v))
	})
}

// AttendingIdEqualFold applies the EqualFold predicate on the "attendingId" field.
func AttendingIdEqualFold(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAttendingId), v))
	})
}

// AttendingIdContainsFold applies the ContainsFold predicate on the "attendingId" field.
func AttendingIdContainsFold(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAttendingId), v))
	})
}

// StartTimeEQ applies the EQ predicate on the "startTime" field.
func StartTimeEQ(v time.Time) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStartTime), v))
	})
}

// StartTimeNEQ applies the NEQ predicate on the "startTime" field.
func StartTimeNEQ(v time.Time) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStartTime), v))
	})
}

// StartTimeIn applies the In predicate on the "startTime" field.
func StartTimeIn(vs ...time.Time) predicate.Slot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Slot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStartTime), v...))
	})
}

// StartTimeNotIn applies the NotIn predicate on the "startTime" field.
func StartTimeNotIn(vs ...time.Time) predicate.Slot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Slot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStartTime), v...))
	})
}

// StartTimeGT applies the GT predicate on the "startTime" field.
func StartTimeGT(v time.Time) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStartTime), v))
	})
}

// StartTimeGTE applies the GTE predicate on the "startTime" field.
func StartTimeGTE(v time.Time) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStartTime), v))
	})
}

// StartTimeLT applies the LT predicate on the "startTime" field.
func StartTimeLT(v time.Time) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStartTime), v))
	})
}

// StartTimeLTE applies the LTE predicate on the "startTime" field.
func StartTimeLTE(v time.Time) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStartTime), v))
	})
}

// EndTimeEQ applies the EQ predicate on the "endTime" field.
func EndTimeEQ(v time.Time) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldEndTime), v))
	})
}

// EndTimeNEQ applies the NEQ predicate on the "endTime" field.
func EndTimeNEQ(v time.Time) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldEndTime), v))
	})
}

// EndTimeIn applies the In predicate on the "endTime" field.
func EndTimeIn(vs ...time.Time) predicate.Slot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Slot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldEndTime), v...))
	})
}

// EndTimeNotIn applies the NotIn predicate on the "endTime" field.
func EndTimeNotIn(vs ...time.Time) predicate.Slot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Slot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldEndTime), v...))
	})
}

// EndTimeGT applies the GT predicate on the "endTime" field.
func EndTimeGT(v time.Time) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldEndTime), v))
	})
}

// EndTimeGTE applies the GTE predicate on the "endTime" field.
func EndTimeGTE(v time.Time) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldEndTime), v))
	})
}

// EndTimeLT applies the LT predicate on the "endTime" field.
func EndTimeLT(v time.Time) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldEndTime), v))
	})
}

// EndTimeLTE applies the LTE predicate on the "endTime" field.
func EndTimeLTE(v time.Time) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldEndTime), v))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...string) predicate.Slot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Slot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...string) predicate.Slot {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Slot(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// StatusContains applies the Contains predicate on the "status" field.
func StatusContains(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldStatus), v))
	})
}

// StatusHasPrefix applies the HasPrefix predicate on the "status" field.
func StatusHasPrefix(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldStatus), v))
	})
}

// StatusHasSuffix applies the HasSuffix predicate on the "status" field.
func StatusHasSuffix(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldStatus), v))
	})
}

// StatusEqualFold applies the EqualFold predicate on the "status" field.
func StatusEqualFold(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldStatus), v))
	})
}

// StatusContainsFold applies the ContainsFold predicate on the "status" field.
func StatusContainsFold(v string) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldStatus), v))
	})
}

// HasAttendingDoctor applies the HasEdge predicate on the "AttendingDoctor" edge.
func HasAttendingDoctor() predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AttendingDoctorTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AttendingDoctorTable, AttendingDoctorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAttendingDoctorWith applies the HasEdge predicate on the "AttendingDoctor" edge with a given conditions (other predicates).
func HasAttendingDoctorWith(preds ...predicate.User) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AttendingDoctorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AttendingDoctorTable, AttendingDoctorColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPatient applies the HasEdge predicate on the "patient" edge.
func HasPatient() predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PatientTable, PatientColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPatientWith applies the HasEdge predicate on the "patient" edge with a given conditions (other predicates).
func HasPatientWith(preds ...predicate.User) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PatientTable, PatientColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Slot) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Slot) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Slot) predicate.Slot {
	return predicate.Slot(func(s *sql.Selector) {
		p(s.Not())
	})
}
