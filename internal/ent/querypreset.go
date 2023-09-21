// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"exusiai.dev/roguestats-backend/internal/ent/querypreset"
	"exusiai.dev/roguestats-backend/internal/ent/research"
	"exusiai.dev/roguestats-backend/internal/model"
)

// QueryPreset is the model entity for the QueryPreset schema.
type QueryPreset struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// ResearchID holds the value of the "research_id" field.
	ResearchID string `json:"research_id,omitempty"`
	// The filter to apply to the events
	Where map[string]interface{} `json:"where,omitempty"`
	// The mapping expr to apply to the events
	Mapping string `json:"mapping,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the QueryPresetQuery when eager-loading is set.
	Edges        QueryPresetEdges `json:"edges"`
	selectValues sql.SelectValues

	GroupCountResult *model.GroupCountResult `json:"static,omitempty"`
}

// QueryPresetEdges holds the relations/edges for other nodes in the graph.
type QueryPresetEdges struct {
	// Research holds the value of the research edge.
	Research *Research `json:"research,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ResearchOrErr returns the Research value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e QueryPresetEdges) ResearchOrErr() (*Research, error) {
	if e.loadedTypes[0] {
		if e.Research == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: research.Label}
		}
		return e.Research, nil
	}
	return nil, &NotLoadedError{edge: "research"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*QueryPreset) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case querypreset.FieldWhere:
			values[i] = new([]byte)
		case querypreset.FieldID, querypreset.FieldName, querypreset.FieldResearchID, querypreset.FieldMapping:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the QueryPreset fields.
func (qp *QueryPreset) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case querypreset.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				qp.ID = value.String
			}
		case querypreset.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				qp.Name = value.String
			}
		case querypreset.FieldResearchID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field research_id", values[i])
			} else if value.Valid {
				qp.ResearchID = value.String
			}
		case querypreset.FieldWhere:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field where", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &qp.Where); err != nil {
					return fmt.Errorf("unmarshal field where: %w", err)
				}
			}
		case querypreset.FieldMapping:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mapping", values[i])
			} else if value.Valid {
				qp.Mapping = value.String
			}
		default:
			qp.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the QueryPreset.
// This includes values selected through modifiers, order, etc.
func (qp *QueryPreset) Value(name string) (ent.Value, error) {
	return qp.selectValues.Get(name)
}

// QueryResearch queries the "research" edge of the QueryPreset entity.
func (qp *QueryPreset) QueryResearch() *ResearchQuery {
	return NewQueryPresetClient(qp.config).QueryResearch(qp)
}

// Update returns a builder for updating this QueryPreset.
// Note that you need to call QueryPreset.Unwrap() before calling this method if this QueryPreset
// was returned from a transaction, and the transaction was committed or rolled back.
func (qp *QueryPreset) Update() *QueryPresetUpdateOne {
	return NewQueryPresetClient(qp.config).UpdateOne(qp)
}

// Unwrap unwraps the QueryPreset entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (qp *QueryPreset) Unwrap() *QueryPreset {
	_tx, ok := qp.config.driver.(*txDriver)
	if !ok {
		panic("ent: QueryPreset is not a transactional entity")
	}
	qp.config.driver = _tx.drv
	return qp
}

// String implements the fmt.Stringer.
func (qp *QueryPreset) String() string {
	var builder strings.Builder
	builder.WriteString("QueryPreset(")
	builder.WriteString(fmt.Sprintf("id=%v, ", qp.ID))
	builder.WriteString("name=")
	builder.WriteString(qp.Name)
	builder.WriteString(", ")
	builder.WriteString("research_id=")
	builder.WriteString(qp.ResearchID)
	builder.WriteString(", ")
	builder.WriteString("where=")
	builder.WriteString(fmt.Sprintf("%v", qp.Where))
	builder.WriteString(", ")
	builder.WriteString("mapping=")
	builder.WriteString(qp.Mapping)
	builder.WriteByte(')')
	return builder.String()
}

// QueryPresets is a parsable slice of QueryPreset.
type QueryPresets []*QueryPreset
