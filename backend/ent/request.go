// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"mocku/backend/ent/request"
	"mocku/backend/ent/users"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Request is the model entity for the Request schema.
type Request struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RequestQuery when eager-loading is set.
	Edges                   RequestEdges `json:"edges"`
	users_requests_made     *int
	users_requests_received *int
	selectValues            sql.SelectValues
}

// RequestEdges holds the relations/edges for other nodes in the graph.
type RequestEdges struct {
	// Requester holds the value of the requester edge.
	Requester *Users `json:"requester,omitempty"`
	// Receiver holds the value of the receiver edge.
	Receiver *Users `json:"receiver,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// RequesterOrErr returns the Requester value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RequestEdges) RequesterOrErr() (*Users, error) {
	if e.Requester != nil {
		return e.Requester, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: users.Label}
	}
	return nil, &NotLoadedError{edge: "requester"}
}

// ReceiverOrErr returns the Receiver value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RequestEdges) ReceiverOrErr() (*Users, error) {
	if e.Receiver != nil {
		return e.Receiver, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: users.Label}
	}
	return nil, &NotLoadedError{edge: "receiver"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Request) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case request.FieldID:
			values[i] = new(sql.NullInt64)
		case request.FieldType, request.FieldStatus, request.FieldTitle, request.FieldDescription:
			values[i] = new(sql.NullString)
		case request.FieldCreatedAt, request.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case request.ForeignKeys[0]: // users_requests_made
			values[i] = new(sql.NullInt64)
		case request.ForeignKeys[1]: // users_requests_received
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Request fields.
func (r *Request) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case request.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			r.ID = int(value.Int64)
		case request.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				r.Type = value.String
			}
		case request.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				r.Status = value.String
			}
		case request.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				r.Title = value.String
			}
		case request.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				r.Description = value.String
			}
		case request.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case request.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				r.UpdatedAt = value.Time
			}
		case request.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field users_requests_made", value)
			} else if value.Valid {
				r.users_requests_made = new(int)
				*r.users_requests_made = int(value.Int64)
			}
		case request.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field users_requests_received", value)
			} else if value.Valid {
				r.users_requests_received = new(int)
				*r.users_requests_received = int(value.Int64)
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Request.
// This includes values selected through modifiers, order, etc.
func (r *Request) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// QueryRequester queries the "requester" edge of the Request entity.
func (r *Request) QueryRequester() *UsersQuery {
	return NewRequestClient(r.config).QueryRequester(r)
}

// QueryReceiver queries the "receiver" edge of the Request entity.
func (r *Request) QueryReceiver() *UsersQuery {
	return NewRequestClient(r.config).QueryReceiver(r)
}

// Update returns a builder for updating this Request.
// Note that you need to call Request.Unwrap() before calling this method if this Request
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Request) Update() *RequestUpdateOne {
	return NewRequestClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Request entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Request) Unwrap() *Request {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Request is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Request) String() string {
	var builder strings.Builder
	builder.WriteString("Request(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("type=")
	builder.WriteString(r.Type)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(r.Status)
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(r.Title)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(r.Description)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(r.UpdatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Requests is a parsable slice of Request.
type Requests []*Request
