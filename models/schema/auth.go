package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Auth holds the schema definition for the Auth entity.
type Auth struct {
	ent.Schema
}

// Fields of the Auth.
func (Auth) Fields() []ent.Field {
	return []ent.Field{
		// 最后认证时间
		field.Time("last_auth_time").Default(time.Now).SchemaType(TimeSchemaType),
		// 认证类型。登录、刷新
		field.Enum("auth_type").Values("login", "refresh"),
	}
}

// Edges of the Auth.
func (Auth) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("belong_to", User.Type).Ref("authHistories").Unique().Annotations(entgql.Bind()),
	}
}
