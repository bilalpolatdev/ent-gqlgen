// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"gqlgen-ent/ent/jobprogress"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// JobProgress is the model entity for the JobProgress schema.
type JobProgress struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// One holds the value of the "One" field.
	One int `json:"One,omitempty"`
	// Two holds the value of the "Two" field.
	Two int `json:"Two,omitempty"`
	// Three holds the value of the "Three" field.
	Three int `json:"Three,omitempty"`
	// Four holds the value of the "Four" field.
	Four int `json:"Four,omitempty"`
	// Five holds the value of the "Five" field.
	Five int `json:"Five,omitempty"`
	// Six holds the value of the "Six" field.
	Six int `json:"Six,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the JobProgressQuery when eager-loading is set.
	Edges        JobProgressEdges `json:"edges"`
	selectValues sql.SelectValues
}

// JobProgressEdges holds the relations/edges for other nodes in the graph.
type JobProgressEdges struct {
	// Progress holds the value of the progress edge.
	Progress []*JobDetail `json:"progress,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ProgressOrErr returns the Progress value or an error if the edge
// was not loaded in eager-loading.
func (e JobProgressEdges) ProgressOrErr() ([]*JobDetail, error) {
	if e.loadedTypes[0] {
		return e.Progress, nil
	}
	return nil, &NotLoadedError{edge: "progress"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*JobProgress) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case jobprogress.FieldID, jobprogress.FieldOne, jobprogress.FieldTwo, jobprogress.FieldThree, jobprogress.FieldFour, jobprogress.FieldFive, jobprogress.FieldSix:
			values[i] = new(sql.NullInt64)
		case jobprogress.FieldCreatedAt, jobprogress.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the JobProgress fields.
func (jp *JobProgress) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case jobprogress.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			jp.ID = int(value.Int64)
		case jobprogress.FieldOne:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field One", values[i])
			} else if value.Valid {
				jp.One = int(value.Int64)
			}
		case jobprogress.FieldTwo:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field Two", values[i])
			} else if value.Valid {
				jp.Two = int(value.Int64)
			}
		case jobprogress.FieldThree:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field Three", values[i])
			} else if value.Valid {
				jp.Three = int(value.Int64)
			}
		case jobprogress.FieldFour:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field Four", values[i])
			} else if value.Valid {
				jp.Four = int(value.Int64)
			}
		case jobprogress.FieldFive:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field Five", values[i])
			} else if value.Valid {
				jp.Five = int(value.Int64)
			}
		case jobprogress.FieldSix:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field Six", values[i])
			} else if value.Valid {
				jp.Six = int(value.Int64)
			}
		case jobprogress.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				jp.CreatedAt = value.Time
			}
		case jobprogress.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				jp.UpdatedAt = value.Time
			}
		default:
			jp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the JobProgress.
// This includes values selected through modifiers, order, etc.
func (jp *JobProgress) Value(name string) (ent.Value, error) {
	return jp.selectValues.Get(name)
}

// QueryProgress queries the "progress" edge of the JobProgress entity.
func (jp *JobProgress) QueryProgress() *JobDetailQuery {
	return NewJobProgressClient(jp.config).QueryProgress(jp)
}

// Update returns a builder for updating this JobProgress.
// Note that you need to call JobProgress.Unwrap() before calling this method if this JobProgress
// was returned from a transaction, and the transaction was committed or rolled back.
func (jp *JobProgress) Update() *JobProgressUpdateOne {
	return NewJobProgressClient(jp.config).UpdateOne(jp)
}

// Unwrap unwraps the JobProgress entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (jp *JobProgress) Unwrap() *JobProgress {
	_tx, ok := jp.config.driver.(*txDriver)
	if !ok {
		panic("ent: JobProgress is not a transactional entity")
	}
	jp.config.driver = _tx.drv
	return jp
}

// String implements the fmt.Stringer.
func (jp *JobProgress) String() string {
	var builder strings.Builder
	builder.WriteString("JobProgress(")
	builder.WriteString(fmt.Sprintf("id=%v, ", jp.ID))
	builder.WriteString("One=")
	builder.WriteString(fmt.Sprintf("%v", jp.One))
	builder.WriteString(", ")
	builder.WriteString("Two=")
	builder.WriteString(fmt.Sprintf("%v", jp.Two))
	builder.WriteString(", ")
	builder.WriteString("Three=")
	builder.WriteString(fmt.Sprintf("%v", jp.Three))
	builder.WriteString(", ")
	builder.WriteString("Four=")
	builder.WriteString(fmt.Sprintf("%v", jp.Four))
	builder.WriteString(", ")
	builder.WriteString("Five=")
	builder.WriteString(fmt.Sprintf("%v", jp.Five))
	builder.WriteString(", ")
	builder.WriteString("Six=")
	builder.WriteString(fmt.Sprintf("%v", jp.Six))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(jp.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(jp.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// JobProgresses is a parsable slice of JobProgress.
type JobProgresses []*JobProgress
