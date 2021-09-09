// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"wing/models/ent/orgunit"
	"wing/models/ent/orgunitposition"
	"wing/models/ent/user"

	"entgo.io/ent/dialect/sql"
)

// OrgUnitPosition is the model entity for the OrgUnitPosition schema.
type OrgUnitPosition struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateByUser holds the value of the "create_by_user" field.
	CreateByUser int `json:"create_by_user,omitempty"`
	// UpdateByUser holds the value of the "update_by_user" field.
	UpdateByUser int `json:"update_by_user,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Duty holds the value of the "duty" field.
	Duty string `json:"duty,omitempty"`
	// Level holds the value of the "level" field.
	Level int `json:"level,omitempty"`
	// OrgUnitID holds the value of the "org_unit_id" field.
	OrgUnitID int `json:"org_unit_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrgUnitPositionQuery when eager-loading is set.
	Edges OrgUnitPositionEdges `json:"edges"`
}

// OrgUnitPositionEdges holds the relations/edges for other nodes in the graph.
type OrgUnitPositionEdges struct {
	// CreateBy holds the value of the create_by edge.
	CreateBy *User `json:"create_by,omitempty"`
	// UpdateBy holds the value of the update_by edge.
	UpdateBy *User `json:"update_by,omitempty"`
	// BelongToOrgUnitMembers holds the value of the belongToOrgUnitMembers edge.
	BelongToOrgUnitMembers []*OrgUnitMember `json:"belongToOrgUnitMembers,omitempty"`
	// BelongToOrgUnit holds the value of the belongToOrgUnit edge.
	BelongToOrgUnit *OrgUnit `json:"belongToOrgUnit,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// CreateByOrErr returns the CreateBy value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrgUnitPositionEdges) CreateByOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.CreateBy == nil {
			// The edge create_by was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.CreateBy, nil
	}
	return nil, &NotLoadedError{edge: "create_by"}
}

// UpdateByOrErr returns the UpdateBy value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrgUnitPositionEdges) UpdateByOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.UpdateBy == nil {
			// The edge update_by was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.UpdateBy, nil
	}
	return nil, &NotLoadedError{edge: "update_by"}
}

// BelongToOrgUnitMembersOrErr returns the BelongToOrgUnitMembers value or an error if the edge
// was not loaded in eager-loading.
func (e OrgUnitPositionEdges) BelongToOrgUnitMembersOrErr() ([]*OrgUnitMember, error) {
	if e.loadedTypes[2] {
		return e.BelongToOrgUnitMembers, nil
	}
	return nil, &NotLoadedError{edge: "belongToOrgUnitMembers"}
}

// BelongToOrgUnitOrErr returns the BelongToOrgUnit value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrgUnitPositionEdges) BelongToOrgUnitOrErr() (*OrgUnit, error) {
	if e.loadedTypes[3] {
		if e.BelongToOrgUnit == nil {
			// The edge belongToOrgUnit was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: orgunit.Label}
		}
		return e.BelongToOrgUnit, nil
	}
	return nil, &NotLoadedError{edge: "belongToOrgUnit"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrgUnitPosition) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case orgunitposition.FieldID, orgunitposition.FieldCreateByUser, orgunitposition.FieldUpdateByUser, orgunitposition.FieldLevel, orgunitposition.FieldOrgUnitID:
			values[i] = new(sql.NullInt64)
		case orgunitposition.FieldName, orgunitposition.FieldDuty:
			values[i] = new(sql.NullString)
		case orgunitposition.FieldCreateTime, orgunitposition.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OrgUnitPosition", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrgUnitPosition fields.
func (oup *OrgUnitPosition) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orgunitposition.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			oup.ID = int(value.Int64)
		case orgunitposition.FieldCreateByUser:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_by_user", values[i])
			} else if value.Valid {
				oup.CreateByUser = int(value.Int64)
			}
		case orgunitposition.FieldUpdateByUser:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_by_user", values[i])
			} else if value.Valid {
				oup.UpdateByUser = int(value.Int64)
			}
		case orgunitposition.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				oup.CreateTime = value.Time
			}
		case orgunitposition.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				oup.UpdateTime = new(time.Time)
				*oup.UpdateTime = value.Time
			}
		case orgunitposition.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				oup.Name = value.String
			}
		case orgunitposition.FieldDuty:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field duty", values[i])
			} else if value.Valid {
				oup.Duty = value.String
			}
		case orgunitposition.FieldLevel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field level", values[i])
			} else if value.Valid {
				oup.Level = int(value.Int64)
			}
		case orgunitposition.FieldOrgUnitID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field org_unit_id", values[i])
			} else if value.Valid {
				oup.OrgUnitID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryCreateBy queries the "create_by" edge of the OrgUnitPosition entity.
func (oup *OrgUnitPosition) QueryCreateBy() *UserQuery {
	return (&OrgUnitPositionClient{config: oup.config}).QueryCreateBy(oup)
}

// QueryUpdateBy queries the "update_by" edge of the OrgUnitPosition entity.
func (oup *OrgUnitPosition) QueryUpdateBy() *UserQuery {
	return (&OrgUnitPositionClient{config: oup.config}).QueryUpdateBy(oup)
}

// QueryBelongToOrgUnitMembers queries the "belongToOrgUnitMembers" edge of the OrgUnitPosition entity.
func (oup *OrgUnitPosition) QueryBelongToOrgUnitMembers() *OrgUnitMemberQuery {
	return (&OrgUnitPositionClient{config: oup.config}).QueryBelongToOrgUnitMembers(oup)
}

// QueryBelongToOrgUnit queries the "belongToOrgUnit" edge of the OrgUnitPosition entity.
func (oup *OrgUnitPosition) QueryBelongToOrgUnit() *OrgUnitQuery {
	return (&OrgUnitPositionClient{config: oup.config}).QueryBelongToOrgUnit(oup)
}

// Update returns a builder for updating this OrgUnitPosition.
// Note that you need to call OrgUnitPosition.Unwrap() before calling this method if this OrgUnitPosition
// was returned from a transaction, and the transaction was committed or rolled back.
func (oup *OrgUnitPosition) Update() *OrgUnitPositionUpdateOne {
	return (&OrgUnitPositionClient{config: oup.config}).UpdateOne(oup)
}

// Unwrap unwraps the OrgUnitPosition entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oup *OrgUnitPosition) Unwrap() *OrgUnitPosition {
	tx, ok := oup.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrgUnitPosition is not a transactional entity")
	}
	oup.config.driver = tx.drv
	return oup
}

// String implements the fmt.Stringer.
func (oup *OrgUnitPosition) String() string {
	var builder strings.Builder
	builder.WriteString("OrgUnitPosition(")
	builder.WriteString(fmt.Sprintf("id=%v", oup.ID))
	builder.WriteString(", create_by_user=")
	builder.WriteString(fmt.Sprintf("%v", oup.CreateByUser))
	builder.WriteString(", update_by_user=")
	builder.WriteString(fmt.Sprintf("%v", oup.UpdateByUser))
	builder.WriteString(", create_time=")
	builder.WriteString(oup.CreateTime.Format(time.ANSIC))
	if v := oup.UpdateTime; v != nil {
		builder.WriteString(", update_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", name=")
	builder.WriteString(oup.Name)
	builder.WriteString(", duty=")
	builder.WriteString(oup.Duty)
	builder.WriteString(", level=")
	builder.WriteString(fmt.Sprintf("%v", oup.Level))
	builder.WriteString(", org_unit_id=")
	builder.WriteString(fmt.Sprintf("%v", oup.OrgUnitID))
	builder.WriteByte(')')
	return builder.String()
}

// OrgUnitPositions is a parsable slice of OrgUnitPosition.
type OrgUnitPositions []*OrgUnitPosition

func (oup OrgUnitPositions) config(cfg config) {
	for _i := range oup {
		oup[_i].config = cfg
	}
}