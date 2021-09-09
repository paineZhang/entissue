package schema

import (
	"wing/models/ent/hook"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"wing/models/hooks"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// 编译器检查接口实现
var _ ent.Interface = (*User)(nil)

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		// 账户名。与LDAP同步
		field.String("account_name").NotEmpty().MaxLen(64).Unique().Annotations(entgql.OrderField("ACCOUNT_NAME")),
		// 人员类型：正式、协力（派遣）、实习、外部
		field.Enum("staff_type").Values("regular", "dispatching", "intern", "external"),
		// 是否在职
		field.Bool("is_on_job").Default(false),
		// 姓。与LDAP同步
		field.String("family_name").MaxLen(20).NotEmpty().Annotations(entgql.OrderField("FAMILY_NAME")),
		// 名。与LDAP同步
		field.String("given_name").MaxLen(20).NotEmpty(),
		// 显示名称，这通常不代表姓和名的组合结果。与LDAP同步
		field.String("display_name").MaxLen(20).NotEmpty().Annotations(entgql.OrderField("DISPLAY_NAME")),
		// 出生日期
		field.Time("birthday").Optional().Nillable().SchemaType(DateSchemaType),
		// 身份证号
		field.String("id_number").Unique().MaxLen(18).Optional().Nillable(),
		// 性别
		field.Enum("sex").Values("male", "female").Optional().Nillable(),
		// 常用电话
		field.String("phone_number").Optional().Unique().MaxLen(20).Nillable(),
		// 常驻地址
		field.String("address").Optional().MaxLen(255).Nillable(),
		// 员工编号
		field.String("staff_id").Optional().MaxLen(64).Unique(),

		// 个人常用邮箱
		field.String("personal_email").Optional().MaxLen(64).Unique(),
		// 内网工作邮箱。与LDAP同步
		field.String("intranet_work_email").NotEmpty().MaxLen(64).Unique(),
		// 外网工作邮箱
		field.String("extranet_work_email").Optional().MaxLen(64).Unique(),

		// 用户头像图片地址
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("account_name"),
		index.Fields("staff_type"),
		index.Fields("family_name", "given_name"),
		index.Fields("family_name"),
		index.Fields("display_name"),
		index.Fields("is_on_job"),
		index.Fields("sex"),
		index.Fields("phone_number"),
		index.Fields("birthday"),
		index.Fields("id_number"),
		index.Fields("staff_id"),
		index.Fields("personal_email"),
		index.Fields("intranet_work_email"),
		index.Fields("extranet_work_email"),
		index.Fields("intranet_work_email", "extranet_work_email").Unique(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{

		edge.To("authHistories", Auth.Type).Annotations(entgql.Bind()),
		edge.To("jobHistories", JobHistory.Type).Annotations(entgql.Bind()),

		// NOTE mixin中的create by和update by，如果不加from是O2O的关系，但需要的是M2O的关系
		edge.From("creates", User.Type).Ref("create_by").Annotations(entgql.Bind()),
		edge.From("updates", User.Type).Ref("update_by").Annotations(entgql.Bind()),

		edge.From("belongToOrgUnitMembers", OrgUnitMember.Type).Ref("user").Annotations(entgql.Bind()),

		edge.From("belongToOg", Organization.Type).Ref("staffs").Annotations(entgql.Bind()),
	}
}

func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(hooks.OnUser, ent.OpCreate|ent.OpUpdateOne|ent.OpUpdate),
		hook.Reject(ent.OpDelete | ent.OpDeleteOne),
	}
}

func (User) Mixin() []ent.Mixin {
	return BasicMixin(User.Type)
}
