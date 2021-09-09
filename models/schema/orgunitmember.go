package schema

import (
	"wing/models/ent/hook"

	"wing/models/hooks"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// OrgUnitMember holds the schema definition for the OrgUnitMember entity.
type OrgUnitMember struct {
	ent.Schema
}

// Fields of the OrgUnitMember.
func (OrgUnitMember) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("is_additional"),
		field.Int("user_id"),
		field.Int("org_unit_id"),
		// 增加这个关系引用，为了是在变更时通过新旧值判断方便
		// ent现在没能生成oldEdges的内容，所以取不到。但是ref后有ID字段就可以了
		field.Int("org_unit_position_id"),
	}
}

// Edges of the OrgUnitMember.
func (OrgUnitMember) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Field("user_id").Unique().Required().Annotations(entgql.Bind()),
		edge.To("position", OrgUnitPosition.Type).
			Field("org_unit_position_id").Unique().Required().Annotations(entgql.Bind()),

		edge.From("belongToOrgUnit", OrgUnit.Type).Ref("members").
			Field("org_unit_id").Unique().Required().Annotations(entgql.Bind()),
	}
}

// Indexes of the Street.
func (OrgUnitMember) Indexes() []ent.Index {
	return []ent.Index{
		// 用户在单元中的任职唯一
		index.Fields("user_id", "org_unit_id").Unique(),
	}
}

func (OrgUnitMember) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(hooks.OnOrgUnitMember,
			ent.OpCreate|ent.OpUpdateOne|ent.OpDeleteOne),

		// ent给不出来批量的UPDATE 和 DELETE 的具体记录，需要拆分成单独的更新和删除
		hook.Reject(ent.OpDelete | ent.OpUpdate),
	}
}

func (OrgUnitMember) Mixin() []ent.Mixin {
	return BasicMixin(User.Type)
}
