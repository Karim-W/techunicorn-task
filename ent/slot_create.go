// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/karim-w/techunicorn-task/ent/slot"
	"github.com/karim-w/techunicorn-task/ent/user"
)

// SlotCreate is the builder for creating a Slot entity.
type SlotCreate struct {
	config
	mutation *SlotMutation
	hooks    []Hook
}

// SetSlotId sets the "slotId" field.
func (sc *SlotCreate) SetSlotId(u uuid.UUID) *SlotCreate {
	sc.mutation.SetSlotId(u)
	return sc
}

// SetPatientName sets the "patientName" field.
func (sc *SlotCreate) SetPatientName(s string) *SlotCreate {
	sc.mutation.SetPatientName(s)
	return sc
}

// SetPatientId sets the "patientId" field.
func (sc *SlotCreate) SetPatientId(s string) *SlotCreate {
	sc.mutation.SetPatientId(s)
	return sc
}

// SetAttendingName sets the "attendingName" field.
func (sc *SlotCreate) SetAttendingName(s string) *SlotCreate {
	sc.mutation.SetAttendingName(s)
	return sc
}

// SetAttendingId sets the "attendingId" field.
func (sc *SlotCreate) SetAttendingId(s string) *SlotCreate {
	sc.mutation.SetAttendingId(s)
	return sc
}

// SetStartTime sets the "startTime" field.
func (sc *SlotCreate) SetStartTime(t time.Time) *SlotCreate {
	sc.mutation.SetStartTime(t)
	return sc
}

// SetEndTime sets the "endTime" field.
func (sc *SlotCreate) SetEndTime(t time.Time) *SlotCreate {
	sc.mutation.SetEndTime(t)
	return sc
}

// SetStatus sets the "status" field.
func (sc *SlotCreate) SetStatus(s string) *SlotCreate {
	sc.mutation.SetStatus(s)
	return sc
}

// SetAttendingDoctorID sets the "AttendingDoctor" edge to the User entity by ID.
func (sc *SlotCreate) SetAttendingDoctorID(id int) *SlotCreate {
	sc.mutation.SetAttendingDoctorID(id)
	return sc
}

// SetNillableAttendingDoctorID sets the "AttendingDoctor" edge to the User entity by ID if the given value is not nil.
func (sc *SlotCreate) SetNillableAttendingDoctorID(id *int) *SlotCreate {
	if id != nil {
		sc = sc.SetAttendingDoctorID(*id)
	}
	return sc
}

// SetAttendingDoctor sets the "AttendingDoctor" edge to the User entity.
func (sc *SlotCreate) SetAttendingDoctor(u *User) *SlotCreate {
	return sc.SetAttendingDoctorID(u.ID)
}

// SetPatientID sets the "patient" edge to the User entity by ID.
func (sc *SlotCreate) SetPatientID(id int) *SlotCreate {
	sc.mutation.SetPatientID(id)
	return sc
}

// SetNillablePatientID sets the "patient" edge to the User entity by ID if the given value is not nil.
func (sc *SlotCreate) SetNillablePatientID(id *int) *SlotCreate {
	if id != nil {
		sc = sc.SetPatientID(*id)
	}
	return sc
}

// SetPatient sets the "patient" edge to the User entity.
func (sc *SlotCreate) SetPatient(u *User) *SlotCreate {
	return sc.SetPatientID(u.ID)
}

// Mutation returns the SlotMutation object of the builder.
func (sc *SlotCreate) Mutation() *SlotMutation {
	return sc.mutation
}

// Save creates the Slot in the database.
func (sc *SlotCreate) Save(ctx context.Context) (*Slot, error) {
	var (
		err  error
		node *Slot
	)
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SlotMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SlotCreate) SaveX(ctx context.Context) *Slot {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SlotCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SlotCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SlotCreate) check() error {
	if _, ok := sc.mutation.SlotId(); !ok {
		return &ValidationError{Name: "slotId", err: errors.New(`ent: missing required field "Slot.slotId"`)}
	}
	if _, ok := sc.mutation.PatientName(); !ok {
		return &ValidationError{Name: "patientName", err: errors.New(`ent: missing required field "Slot.patientName"`)}
	}
	if v, ok := sc.mutation.PatientName(); ok {
		if err := slot.PatientNameValidator(v); err != nil {
			return &ValidationError{Name: "patientName", err: fmt.Errorf(`ent: validator failed for field "Slot.patientName": %w`, err)}
		}
	}
	if _, ok := sc.mutation.PatientId(); !ok {
		return &ValidationError{Name: "patientId", err: errors.New(`ent: missing required field "Slot.patientId"`)}
	}
	if v, ok := sc.mutation.PatientId(); ok {
		if err := slot.PatientIdValidator(v); err != nil {
			return &ValidationError{Name: "patientId", err: fmt.Errorf(`ent: validator failed for field "Slot.patientId": %w`, err)}
		}
	}
	if _, ok := sc.mutation.AttendingName(); !ok {
		return &ValidationError{Name: "attendingName", err: errors.New(`ent: missing required field "Slot.attendingName"`)}
	}
	if v, ok := sc.mutation.AttendingName(); ok {
		if err := slot.AttendingNameValidator(v); err != nil {
			return &ValidationError{Name: "attendingName", err: fmt.Errorf(`ent: validator failed for field "Slot.attendingName": %w`, err)}
		}
	}
	if _, ok := sc.mutation.AttendingId(); !ok {
		return &ValidationError{Name: "attendingId", err: errors.New(`ent: missing required field "Slot.attendingId"`)}
	}
	if v, ok := sc.mutation.AttendingId(); ok {
		if err := slot.AttendingIdValidator(v); err != nil {
			return &ValidationError{Name: "attendingId", err: fmt.Errorf(`ent: validator failed for field "Slot.attendingId": %w`, err)}
		}
	}
	if _, ok := sc.mutation.StartTime(); !ok {
		return &ValidationError{Name: "startTime", err: errors.New(`ent: missing required field "Slot.startTime"`)}
	}
	if _, ok := sc.mutation.EndTime(); !ok {
		return &ValidationError{Name: "endTime", err: errors.New(`ent: missing required field "Slot.endTime"`)}
	}
	if _, ok := sc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Slot.status"`)}
	}
	if v, ok := sc.mutation.Status(); ok {
		if err := slot.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Slot.status": %w`, err)}
		}
	}
	return nil
}

func (sc *SlotCreate) sqlSave(ctx context.Context) (*Slot, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sc *SlotCreate) createSpec() (*Slot, *sqlgraph.CreateSpec) {
	var (
		_node = &Slot{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: slot.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: slot.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.SlotId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: slot.FieldSlotId,
		})
		_node.SlotId = value
	}
	if value, ok := sc.mutation.PatientName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: slot.FieldPatientName,
		})
		_node.PatientName = value
	}
	if value, ok := sc.mutation.PatientId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: slot.FieldPatientId,
		})
		_node.PatientId = value
	}
	if value, ok := sc.mutation.AttendingName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: slot.FieldAttendingName,
		})
		_node.AttendingName = value
	}
	if value, ok := sc.mutation.AttendingId(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: slot.FieldAttendingId,
		})
		_node.AttendingId = value
	}
	if value, ok := sc.mutation.StartTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: slot.FieldStartTime,
		})
		_node.StartTime = value
	}
	if value, ok := sc.mutation.EndTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: slot.FieldEndTime,
		})
		_node.EndTime = value
	}
	if value, ok := sc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: slot.FieldStatus,
		})
		_node.Status = value
	}
	if nodes := sc.mutation.AttendingDoctorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   slot.AttendingDoctorTable,
			Columns: []string{slot.AttendingDoctorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_attending = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.PatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   slot.PatientTable,
			Columns: []string{slot.PatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_visiting = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SlotCreateBulk is the builder for creating many Slot entities in bulk.
type SlotCreateBulk struct {
	config
	builders []*SlotCreate
}

// Save creates the Slot entities in the database.
func (scb *SlotCreateBulk) Save(ctx context.Context) ([]*Slot, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Slot, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SlotMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SlotCreateBulk) SaveX(ctx context.Context) []*Slot {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SlotCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SlotCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
