package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Organization holds the schema definition for the Organization entity.
type Organization struct {
	ent.Schema
}

// Fields of the Organization.
func (Organization) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Organization.
func (Organization) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("units", OrgUnit.Type).Annotations(entgql.Bind()),
		edge.To("staffs", User.Type).Annotations(entgql.Bind()),
	}
}

func (Organization) Mixin() []ent.Mixin {
	return BasicMixin(User.Type)
}
