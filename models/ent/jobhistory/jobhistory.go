// Code generated by entc, DO NOT EDIT.

package jobhistory

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the jobhistory type in the database.
	Label = "job_history"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateByUser holds the string denoting the create_by_user field in the database.
	FieldCreateByUser = "create_by_user"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// FieldJobEntryLeaveType holds the string denoting the job_entry_leave_type field in the database.
	FieldJobEntryLeaveType = "job_entry_leave_type"
	// EdgeCreateBy holds the string denoting the create_by edge name in mutations.
	EdgeCreateBy = "create_by"
	// EdgeBelongTo holds the string denoting the belong_to edge name in mutations.
	EdgeBelongTo = "belong_to"
	// Table holds the table name of the jobhistory in the database.
	Table = "job_histories"
	// CreateByTable is the table that holds the create_by relation/edge.
	CreateByTable = "job_histories"
	// CreateByInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	CreateByInverseTable = "users"
	// CreateByColumn is the table column denoting the create_by relation/edge.
	CreateByColumn = "create_by_user"
	// BelongToTable is the table that holds the belong_to relation/edge.
	BelongToTable = "job_histories"
	// BelongToInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	BelongToInverseTable = "users"
	// BelongToColumn is the table column denoting the belong_to relation/edge.
	BelongToColumn = "user_job_histories"
)

// Columns holds all SQL columns for jobhistory fields.
var Columns = []string{
	FieldID,
	FieldCreateByUser,
	FieldCreateTime,
	FieldDate,
	FieldJobEntryLeaveType,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "job_histories"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_job_histories",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "wing/models/ent/runtime"
//
var (
	Hooks [1]ent.Hook
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
)

// JobEntryLeaveType defines the type for the "job_entry_leave_type" enum field.
type JobEntryLeaveType string

// JobEntryLeaveType values.
const (
	JobEntryLeaveTypeEntry JobEntryLeaveType = "entry"
	JobEntryLeaveTypeLeave JobEntryLeaveType = "leave"
)

func (jelt JobEntryLeaveType) String() string {
	return string(jelt)
}

// JobEntryLeaveTypeValidator is a validator for the "job_entry_leave_type" field enum values. It is called by the builders before save.
func JobEntryLeaveTypeValidator(jelt JobEntryLeaveType) error {
	switch jelt {
	case JobEntryLeaveTypeEntry, JobEntryLeaveTypeLeave:
		return nil
	default:
		return fmt.Errorf("jobhistory: invalid enum value for job_entry_leave_type field: %q", jelt)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (jelt JobEntryLeaveType) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(jelt.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (jelt *JobEntryLeaveType) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*jelt = JobEntryLeaveType(str)
	if err := JobEntryLeaveTypeValidator(*jelt); err != nil {
		return fmt.Errorf("%s is not a valid JobEntryLeaveType", str)
	}
	return nil
}
