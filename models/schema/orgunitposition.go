package schema

import (
	"fmt"
	"wing/models/ent/hook"
	"wing/models/hooks"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// OrgUnitPosition holds the schema definition for the OrgUnitPosition entity.
type OrgUnitPosition struct {
	ent.Schema
}

// Fields of the OrgUnitPosition.
func (OrgUnitPosition) Fields() []ent.Field {
	return []ent.Field{
		// 名称
		field.String("name"),
		// 职责描述
		field.String("duty").Optional(),
		// 等级。不能为负值。值越小，等级越高。这主要用于职务的权限继承关系、显示排序上等使用。
		field.Int("level").Validate(func(value int) error {
			if value < 0 {
				return fmt.Errorf("组织单元职位等级不能为负值")
			}
			return nil
		}),

		// 增加这个关系引用，为了是在变更时通过新旧值判断方便
		// ent现在没能生成oldEdges的内容，所以取不到。但是ref后有ID字段就可以了
		field.Int("org_unit_id"),
	}
}

// Edges of the OrgUnitPosition.
func (OrgUnitPosition) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("belongToOrgUnitMembers", OrgUnitMember.Type).
			Ref("position").Annotations(entgql.Bind()),

		edge.From("belongToOrgUnit", OrgUnit.Type).Ref("positions").Field("org_unit_id").
			Unique().Required().Annotations(entgql.Bind()),
	}
}

func (OrgUnitPosition) Indexes() []ent.Index {
	return []ent.Index{
		// 职位在组织单元中唯一
		index.Fields("name").
			Edges("belongToOrgUnit").
			Unique(),

		// level 在单元中唯一。
		// NOTE org_unit_id是后加的对belongToOrgUnit edge的引用。
		//      所以，这里也可以用Edges("belongToOrgUnit").的写法。是一样的这里是留个例子可以做参考。
		index.Fields("level", "org_unit_id").Unique(),
	}
}

func (OrgUnitPosition) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(hooks.OnOrgUnitPosition, ent.OpDeleteOne),

		// ent给不出来批量的UPDATE 和 DELETE 的具体记录，需要拆分成单独的更新和删除
		hook.Reject(ent.OpDelete),
	}
}

func (OrgUnitPosition) Mixin() []ent.Mixin {
	return BasicMixin(User.Type)
}
