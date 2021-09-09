package schema

import (
	"wing/models/ent/privacy"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Resource holds the schema definition for the Resource entity.
type Resource struct {
	ent.Schema
}

// Fields of the Resource.
func (Resource) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique().NotEmpty().MaxLen(64),
		field.String("type").Unique().NotEmpty().MaxLen(64),
	}
}

func (Resource) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
		index.Fields("type"),
	}
}

// Edges of the Resource.
func (Resource) Edges() []ent.Edge {
	return nil
}

func (Resource) Policy() ent.Policy {
	// NOTE 资源数据比较特殊。隐私权限上，阻止一切修改，允许一切访问。
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}
