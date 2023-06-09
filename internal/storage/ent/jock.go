// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"jocer/internal/storage/ent/jock"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Jock is the model entity for the Jock schema.
type Jock struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Content holds the value of the "content" field.
	Content      string `json:"content,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Jock) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case jock.FieldName, jock.FieldContent:
			values[i] = new(sql.NullString)
		case jock.FieldID:
			values[i] = new(uuid.UUID)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Jock fields.
func (j *Jock) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case jock.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				j.ID = *value
			}
		case jock.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				j.Name = value.String
			}
		case jock.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				j.Content = value.String
			}
		default:
			j.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Jock.
// This includes values selected through modifiers, order, etc.
func (j *Jock) Value(name string) (ent.Value, error) {
	return j.selectValues.Get(name)
}

// Update returns a builder for updating this Jock.
// Note that you need to call Jock.Unwrap() before calling this method if this Jock
// was returned from a transaction, and the transaction was committed or rolled back.
func (j *Jock) Update() *JockUpdateOne {
	return NewJockClient(j.config).UpdateOne(j)
}

// Unwrap unwraps the Jock entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (j *Jock) Unwrap() *Jock {
	_tx, ok := j.config.driver.(*txDriver)
	if !ok {
		panic("ent: Jock is not a transactional entity")
	}
	j.config.driver = _tx.drv
	return j
}

// String implements the fmt.Stringer.
func (j *Jock) String() string {
	var builder strings.Builder
	builder.WriteString("Jock(")
	builder.WriteString(fmt.Sprintf("id=%v, ", j.ID))
	builder.WriteString("name=")
	builder.WriteString(j.Name)
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(j.Content)
	builder.WriteByte(')')
	return builder.String()
}

// Jocks is a parsable slice of Jock.
type Jocks []*Jock
