package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// NOTE 系统设置使用当前对象。
type System struct {
	ent.Schema
}

// 编译器检查接口实现
var _ ent.Interface = (*System)(nil)

func (System) Fields() []ent.Field {

	return []ent.Field{
		field.String("name").Unique().MaxLen(64),
		// field.String("secret_key"),

	}
}

func (System) Mixin() []ent.Mixin {
	return BasicMixin(System.Type)
}
