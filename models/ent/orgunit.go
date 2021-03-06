// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"
	"wing/models/ent/organization"
	"wing/models/ent/orgunit"
	"wing/models/ent/user"

	"entgo.io/ent/dialect/sql"
)

// OrgUnit is the model entity for the OrgUnit schema.
type OrgUnit struct {
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
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrgUnitQuery when eager-loading is set.
	Edges              OrgUnitEdges `json:"edges"`
	org_unit_sub_units *int
	organization_units *int
}

// OrgUnitEdges holds the relations/edges for other nodes in the graph.
type OrgUnitEdges struct {
	// CreateBy holds the value of the create_by edge.
	CreateBy *User `json:"create_by,omitempty"`
	// UpdateBy holds the value of the update_by edge.
	UpdateBy *User `json:"update_by,omitempty"`
	// Members holds the value of the members edge.
	Members []*OrgUnitMember `json:"members,omitempty"`
	// Positions holds the value of the positions edge.
	Positions []*OrgUnitPosition `json:"positions,omitempty"`
	// SupUnit holds the value of the supUnit edge.
	SupUnit *OrgUnit `json:"supUnit,omitempty"`
	// SubUnits holds the value of the subUnits edge.
	SubUnits []*OrgUnit `json:"subUnits,omitempty"`
	// BelongToOrg holds the value of the belongToOrg edge.
	BelongToOrg *Organization `json:"belongToOrg,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [7]bool
}

// CreateByOrErr returns the CreateBy value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrgUnitEdges) CreateByOrErr() (*User, error) {
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
func (e OrgUnitEdges) UpdateByOrErr() (*User, error) {
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

// MembersOrErr returns the Members value or an error if the edge
// was not loaded in eager-loading.
func (e OrgUnitEdges) MembersOrErr() ([]*OrgUnitMember, error) {
	if e.loadedTypes[2] {
		return e.Members, nil
	}
	return nil, &NotLoadedError{edge: "members"}
}

// PositionsOrErr returns the Positions value or an error if the edge
// was not loaded in eager-loading.
func (e OrgUnitEdges) PositionsOrErr() ([]*OrgUnitPosition, error) {
	if e.loadedTypes[3] {
		return e.Positions, nil
	}
	return nil, &NotLoadedError{edge: "positions"}
}

// SupUnitOrErr returns the SupUnit value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrgUnitEdges) SupUnitOrErr() (*OrgUnit, error) {
	if e.loadedTypes[4] {
		if e.SupUnit == nil {
			// The edge supUnit was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: orgunit.Label}
		}
		return e.SupUnit, nil
	}
	return nil, &NotLoadedError{edge: "supUnit"}
}

// SubUnitsOrErr returns the SubUnits value or an error if the edge
// was not loaded in eager-loading.
func (e OrgUnitEdges) SubUnitsOrErr() ([]*OrgUnit, error) {
	if e.loadedTypes[5] {
		return e.SubUnits, nil
	}
	return nil, &NotLoadedError{edge: "subUnits"}
}

// BelongToOrgOrErr returns the BelongToOrg value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrgUnitEdges) BelongToOrgOrErr() (*Organization, error) {
	if e.loadedTypes[6] {
		if e.BelongToOrg == nil {
			// The edge belongToOrg was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: organization.Label}
		}
		return e.BelongToOrg, nil
	}
	return nil, &NotLoadedError{edge: "belongToOrg"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrgUnit) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case orgunit.FieldID, orgunit.FieldCreateByUser, orgunit.FieldUpdateByUser:
			values[i] = new(sql.NullInt64)
		case orgunit.FieldName, orgunit.FieldDuty:
			values[i] = new(sql.NullString)
		case orgunit.FieldCreateTime, orgunit.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case orgunit.ForeignKeys[0]: // org_unit_sub_units
			values[i] = new(sql.NullInt64)
		case orgunit.ForeignKeys[1]: // organization_units
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type OrgUnit", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrgUnit fields.
func (ou *OrgUnit) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orgunit.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ou.ID = int(value.Int64)
		case orgunit.FieldCreateByUser:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field create_by_user", values[i])
			} else if value.Valid {
				ou.CreateByUser = int(value.Int64)
			}
		case orgunit.FieldUpdateByUser:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field update_by_user", values[i])
			} else if value.Valid {
				ou.UpdateByUser = int(value.Int64)
			}
		case orgunit.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				ou.CreateTime = value.Time
			}
		case orgunit.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				ou.UpdateTime = new(time.Time)
				*ou.UpdateTime = value.Time
			}
		case orgunit.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ou.Name = value.String
			}
		case orgunit.FieldDuty:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field duty", values[i])
			} else if value.Valid {
				ou.Duty = value.String
			}
		case orgunit.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field org_unit_sub_units", value)
			} else if value.Valid {
				ou.org_unit_sub_units = new(int)
				*ou.org_unit_sub_units = int(value.Int64)
			}
		case orgunit.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field organization_units", value)
			} else if value.Valid {
				ou.organization_units = new(int)
				*ou.organization_units = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryCreateBy queries the "create_by" edge of the OrgUnit entity.
func (ou *OrgUnit) QueryCreateBy() *UserQuery {
	return (&OrgUnitClient{config: ou.config}).QueryCreateBy(ou)
}

// QueryUpdateBy queries the "update_by" edge of the OrgUnit entity.
func (ou *OrgUnit) QueryUpdateBy() *UserQuery {
	return (&OrgUnitClient{config: ou.config}).QueryUpdateBy(ou)
}

// QueryMembers queries the "members" edge of the OrgUnit entity.
func (ou *OrgUnit) QueryMembers() *OrgUnitMemberQuery {
	return (&OrgUnitClient{config: ou.config}).QueryMembers(ou)
}

// QueryPositions queries the "positions" edge of the OrgUnit entity.
func (ou *OrgUnit) QueryPositions() *OrgUnitPositionQuery {
	return (&OrgUnitClient{config: ou.config}).QueryPositions(ou)
}

// QuerySupUnit queries the "supUnit" edge of the OrgUnit entity.
func (ou *OrgUnit) QuerySupUnit() *OrgUnitQuery {
	return (&OrgUnitClient{config: ou.config}).QuerySupUnit(ou)
}

// QuerySubUnits queries the "subUnits" edge of the OrgUnit entity.
func (ou *OrgUnit) QuerySubUnits() *OrgUnitQuery {
	return (&OrgUnitClient{config: ou.config}).QuerySubUnits(ou)
}

// QueryBelongToOrg queries the "belongToOrg" edge of the OrgUnit entity.
func (ou *OrgUnit) QueryBelongToOrg() *OrganizationQuery {
	return (&OrgUnitClient{config: ou.config}).QueryBelongToOrg(ou)
}

// Update returns a builder for updating this OrgUnit.
// Note that you need to call OrgUnit.Unwrap() before calling this method if this OrgUnit
// was returned from a transaction, and the transaction was committed or rolled back.
func (ou *OrgUnit) Update() *OrgUnitUpdateOne {
	return (&OrgUnitClient{config: ou.config}).UpdateOne(ou)
}

// Unwrap unwraps the OrgUnit entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ou *OrgUnit) Unwrap() *OrgUnit {
	tx, ok := ou.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrgUnit is not a transactional entity")
	}
	ou.config.driver = tx.drv
	return ou
}

// String implements the fmt.Stringer.
func (ou *OrgUnit) String() string {
	var builder strings.Builder
	builder.WriteString("OrgUnit(")
	builder.WriteString(fmt.Sprintf("id=%v", ou.ID))
	builder.WriteString(", create_by_user=")
	builder.WriteString(fmt.Sprintf("%v", ou.CreateByUser))
	builder.WriteString(", update_by_user=")
	builder.WriteString(fmt.Sprintf("%v", ou.UpdateByUser))
	builder.WriteString(", create_time=")
	builder.WriteString(ou.CreateTime.Format(time.ANSIC))
	if v := ou.UpdateTime; v != nil {
		builder.WriteString(", update_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", name=")
	builder.WriteString(ou.Name)
	builder.WriteString(", duty=")
	builder.WriteString(ou.Duty)
	builder.WriteByte(')')
	return builder.String()
}

// OrgUnits is a parsable slice of OrgUnit.
type OrgUnits []*OrgUnit

func (ou OrgUnits) config(cfg config) {
	for _i := range ou {
		ou[_i].config = cfg
	}
}
