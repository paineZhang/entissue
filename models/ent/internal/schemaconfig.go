// Code generated by entc, DO NOT EDIT.

package internal

import "context"

// SchemaConfig represents alternative schema names for all tables
// that can be passed at runtime.
type SchemaConfig struct {
	Auth               string // Auth table.
	JobHistory         string // JobHistory table.
	OrgUnit            string // OrgUnit table.
	OrgUnitMember      string // OrgUnitMember table.
	OrgUnitPosition    string // OrgUnitPosition table.
	Organization       string // Organization table.
	OrganizationStaffs string // Organization-staffs->User table.
	Resource           string // Resource table.
	System             string // System table.
	User               string // User table.
}

type schemaCtxKey struct{}

// SchemaConfigFromContext returns a SchemaConfig stored inside a context, or empty if there isn't one.
func SchemaConfigFromContext(ctx context.Context) SchemaConfig {
	config, _ := ctx.Value(schemaCtxKey{}).(SchemaConfig)
	return config
}

// NewSchemaConfigContext returns a new context with the given SchemaConfig attached.
func NewSchemaConfigContext(parent context.Context, config SchemaConfig) context.Context {
	return context.WithValue(parent, schemaCtxKey{}, config)
}