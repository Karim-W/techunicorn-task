// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/karim-w/techunicorn-task/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// UserId holds the value of the "userId" field.
	UserId uuid.UUID `json:"userId,omitempty"`
	// FirstName holds the value of the "FirstName" field.
	FirstName string `json:"FirstName,omitempty"`
	// LastName holds the value of the "LastName" field.
	LastName string `json:"LastName,omitempty"`
	// Email holds the value of the "Email" field.
	Email string `json:"Email,omitempty"`
	// Password holds the value of the "Password" field.
	Password string `json:"Password,omitempty"`
	// Phone holds the value of the "Phone" field.
	Phone string `json:"Phone,omitempty"`
	// Type holds the value of the "Type" field.
	Type string `json:"Type,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Attending holds the value of the Attending edge.
	Attending []*Slot `json:"Attending,omitempty"`
	// Visiting holds the value of the Visiting edge.
	Visiting []*Slot `json:"Visiting,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// AttendingOrErr returns the Attending value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) AttendingOrErr() ([]*Slot, error) {
	if e.loadedTypes[0] {
		return e.Attending, nil
	}
	return nil, &NotLoadedError{edge: "Attending"}
}

// VisitingOrErr returns the Visiting value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) VisitingOrErr() ([]*Slot, error) {
	if e.loadedTypes[1] {
		return e.Visiting, nil
	}
	return nil, &NotLoadedError{edge: "Visiting"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldFirstName, user.FieldLastName, user.FieldEmail, user.FieldPassword, user.FieldPhone, user.FieldType:
			values[i] = new(sql.NullString)
		case user.FieldUserId:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldUserId:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field userId", values[i])
			} else if value != nil {
				u.UserId = *value
			}
		case user.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field FirstName", values[i])
			} else if value.Valid {
				u.FirstName = value.String
			}
		case user.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field LastName", values[i])
			} else if value.Valid {
				u.LastName = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Phone", values[i])
			} else if value.Valid {
				u.Phone = value.String
			}
		case user.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Type", values[i])
			} else if value.Valid {
				u.Type = value.String
			}
		}
	}
	return nil
}

// QueryAttending queries the "Attending" edge of the User entity.
func (u *User) QueryAttending() *SlotQuery {
	return (&UserClient{config: u.config}).QueryAttending(u)
}

// QueryVisiting queries the "Visiting" edge of the User entity.
func (u *User) QueryVisiting() *SlotQuery {
	return (&UserClient{config: u.config}).QueryVisiting(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", userId=")
	builder.WriteString(fmt.Sprintf("%v", u.UserId))
	builder.WriteString(", FirstName=")
	builder.WriteString(u.FirstName)
	builder.WriteString(", LastName=")
	builder.WriteString(u.LastName)
	builder.WriteString(", Email=")
	builder.WriteString(u.Email)
	builder.WriteString(", Password=")
	builder.WriteString(u.Password)
	builder.WriteString(", Phone=")
	builder.WriteString(u.Phone)
	builder.WriteString(", Type=")
	builder.WriteString(u.Type)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}