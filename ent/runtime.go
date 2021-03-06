// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/karim-w/techunicorn-task/ent/schema"
	"github.com/karim-w/techunicorn-task/ent/slot"
	"github.com/karim-w/techunicorn-task/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	slotFields := schema.Slot{}.Fields()
	_ = slotFields
	// slotDescPatientName is the schema descriptor for patientName field.
	slotDescPatientName := slotFields[1].Descriptor()
	// slot.PatientNameValidator is a validator for the "patientName" field. It is called by the builders before save.
	slot.PatientNameValidator = slotDescPatientName.Validators[0].(func(string) error)
	// slotDescPatientId is the schema descriptor for patientId field.
	slotDescPatientId := slotFields[2].Descriptor()
	// slot.PatientIdValidator is a validator for the "patientId" field. It is called by the builders before save.
	slot.PatientIdValidator = slotDescPatientId.Validators[0].(func(string) error)
	// slotDescAttendingName is the schema descriptor for attendingName field.
	slotDescAttendingName := slotFields[3].Descriptor()
	// slot.AttendingNameValidator is a validator for the "attendingName" field. It is called by the builders before save.
	slot.AttendingNameValidator = slotDescAttendingName.Validators[0].(func(string) error)
	// slotDescAttendingId is the schema descriptor for attendingId field.
	slotDescAttendingId := slotFields[4].Descriptor()
	// slot.AttendingIdValidator is a validator for the "attendingId" field. It is called by the builders before save.
	slot.AttendingIdValidator = slotDescAttendingId.Validators[0].(func(string) error)
	// slotDescStatus is the schema descriptor for status field.
	slotDescStatus := slotFields[7].Descriptor()
	// slot.StatusValidator is a validator for the "status" field. It is called by the builders before save.
	slot.StatusValidator = slotDescStatus.Validators[0].(func(string) error)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescFirstName is the schema descriptor for FirstName field.
	userDescFirstName := userFields[1].Descriptor()
	// user.DefaultFirstName holds the default value on creation for the FirstName field.
	user.DefaultFirstName = userDescFirstName.Default.(string)
	// userDescLastName is the schema descriptor for LastName field.
	userDescLastName := userFields[2].Descriptor()
	// user.DefaultLastName holds the default value on creation for the LastName field.
	user.DefaultLastName = userDescLastName.Default.(string)
	// userDescPassword is the schema descriptor for Password field.
	userDescPassword := userFields[4].Descriptor()
	// user.PasswordValidator is a validator for the "Password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescPhone is the schema descriptor for Phone field.
	userDescPhone := userFields[5].Descriptor()
	// user.DefaultPhone holds the default value on creation for the Phone field.
	user.DefaultPhone = userDescPhone.Default.(string)
	// userDescType is the schema descriptor for Type field.
	userDescType := userFields[6].Descriptor()
	// user.TypeValidator is a validator for the "Type" field. It is called by the builders before save.
	user.TypeValidator = userDescType.Validators[0].(func(string) error)
}
