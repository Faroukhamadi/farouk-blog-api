// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/Faroukhamadi/farouk-blog-api/ent/content"
	"github.com/Faroukhamadi/farouk-blog-api/ent/post"
)

// Content is the model entity for the Content schema.
type Content struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ContentQuery when eager-loading is set.
	Edges        ContentEdges `json:"edges"`
	post_content *int
}

// ContentEdges holds the relations/edges for other nodes in the graph.
type ContentEdges struct {
	// Post holds the value of the post edge.
	Post *Post `json:"post,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PostOrErr returns the Post value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ContentEdges) PostOrErr() (*Post, error) {
	if e.loadedTypes[0] {
		if e.Post == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: post.Label}
		}
		return e.Post, nil
	}
	return nil, &NotLoadedError{edge: "post"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Content) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case content.FieldID:
			values[i] = new(sql.NullInt64)
		case content.FieldTitle, content.FieldText:
			values[i] = new(sql.NullString)
		case content.FieldCreatedAt, content.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case content.ForeignKeys[0]: // post_content
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Content", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Content fields.
func (c *Content) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case content.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case content.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				c.CreatedAt = value.Time
			}
		case content.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				c.UpdatedAt = value.Time
			}
		case content.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				c.Title = value.String
			}
		case content.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				c.Text = value.String
			}
		case content.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field post_content", value)
			} else if value.Valid {
				c.post_content = new(int)
				*c.post_content = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryPost queries the "post" edge of the Content entity.
func (c *Content) QueryPost() *PostQuery {
	return (&ContentClient{config: c.config}).QueryPost(c)
}

// Update returns a builder for updating this Content.
// Note that you need to call Content.Unwrap() before calling this method if this Content
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Content) Update() *ContentUpdateOne {
	return (&ContentClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Content entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Content) Unwrap() *Content {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Content is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Content) String() string {
	var builder strings.Builder
	builder.WriteString("Content(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("created_at=")
	builder.WriteString(c.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(c.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(c.Title)
	builder.WriteString(", ")
	builder.WriteString("text=")
	builder.WriteString(c.Text)
	builder.WriteByte(')')
	return builder.String()
}

// Contents is a parsable slice of Content.
type Contents []*Content

func (c Contents) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}