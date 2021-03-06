// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// SlotsColumns holds the columns for the "slots" table.
	SlotsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "slot_id", Type: field.TypeUUID, Unique: true},
		{Name: "patient_name", Type: field.TypeString},
		{Name: "patient_id", Type: field.TypeString},
		{Name: "attending_name", Type: field.TypeString},
		{Name: "attending_id", Type: field.TypeString},
		{Name: "start_time", Type: field.TypeTime},
		{Name: "end_time", Type: field.TypeTime},
		{Name: "status", Type: field.TypeString},
		{Name: "user_attending", Type: field.TypeInt, Nullable: true},
		{Name: "user_visiting", Type: field.TypeInt, Nullable: true},
	}
	// SlotsTable holds the schema information for the "slots" table.
	SlotsTable = &schema.Table{
		Name:       "slots",
		Columns:    SlotsColumns,
		PrimaryKey: []*schema.Column{SlotsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "slots_users_Attending",
				Columns:    []*schema.Column{SlotsColumns[9]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "slots_users_Visiting",
				Columns:    []*schema.Column{SlotsColumns[10]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeUUID, Unique: true},
		{Name: "first_name", Type: field.TypeString, Default: ""},
		{Name: "last_name", Type: field.TypeString, Default: ""},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString, Default: ""},
		{Name: "type", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		SlotsTable,
		UsersTable,
	}
)

func init() {
	SlotsTable.ForeignKeys[0].RefTable = UsersTable
	SlotsTable.ForeignKeys[1].RefTable = UsersTable
}
