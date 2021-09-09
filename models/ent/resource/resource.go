// Code generated by entc, DO NOT EDIT.

package resource

import (
	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the resource type in the database.
	Label = "resource"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// Table holds the table name of the resource in the database.
	Table = "resources"
)

// Columns holds all SQL columns for resource fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldType,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
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
	Hooks  [1]ent.Hook
	Policy ent.Policy
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// TypeValidator is a validator for the "type" field. It is called by the builders before save.
	TypeValidator func(string) error
)