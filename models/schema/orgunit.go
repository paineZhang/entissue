package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// OrgUnit holds the schema definition for the OrgUnit entity.
type OrgUnit struct {
	ent.Schema
}

// Fields of the OrgUnit.
func (OrgUnit) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("duty").Optional(),
	}
}

// Edges of the OrgUnit.
func (OrgUnit) Edges() []ent.Edge {
	return []ent.Edge{
		// 成员
		edge.To("members", OrgUnitMember.Type).Annotations(entgql.Bind()),
		// 属于当前组织单元的职务
		edge.To("positions", OrgUnitPosition.Type).Annotations(entgql.Bind()),
		// 下层单元和上层单元。上层单元，如果为空，则是顶级单元，一个组织应该只有一个顶级单元。
		edge.To("subUnits", OrgUnit.Type).From("supUnit").Unique().Annotations(entgql.Bind()),

		edge.From("belongToOrg", Organization.Type).Ref("units").Unique().Required().Annotations(entgql.Bind()),
	}
}

func (OrgUnit) Mixin() []ent.Mixin {
	return BasicMixin(User.Type)
}
