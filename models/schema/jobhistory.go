package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// JobHistory holds the schema definition for the JobHistory entity.
type JobHistory struct {
	ent.Schema
}

// Fields of the JobHistory.
func (JobHistory) Fields() []ent.Field {
	return []ent.Field{
		// 日期
		field.Time("date").SchemaType(DateSchemaType).
			Annotations(entgql.OrderField("DATE")),
		// 离入职类型
		field.Enum("job_entry_leave_type").Values("entry", "leave"),
	}
}

func (JobHistory) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("date"),
		index.Fields("create_time"),
		index.Fields("job_entry_leave_type"),
		index.Fields("date", "job_entry_leave_type"),
	}
}

// Edges of the Auth.
func (JobHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("belong_to", User.Type).Ref("jobHistories").Required().Annotations(entgql.Bind()).Unique(),
	}
}

func (JobHistory) Mixin() []ent.Mixin {
	return BasicMixinWith(JobHistory.Type,
		BasicMixinTypeBasePolicy,
		BasicMixinTypeCreateBy,
		BasicMixinTypeCreateTime)
}
