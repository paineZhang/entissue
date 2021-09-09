package schema

import (
	"context"
	"time"

	"wing/models/ent/hook"
	"wing/models/util"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

func SetAuthUserIDToCtx(ctx context.Context, userID int) context.Context {
	return context.WithValue(ctx, UserIdContextKey{}, userID)
}

func getAuthUserIdFromCtx(ctx context.Context) *int {
	id := ctx.Value(UserIdContextKey{})
	if id == nil {
		return nil
	}

	uid, ok := ctx.Value(UserIdContextKey{}).(int)
	if !ok {
		panic("类型错误")
	}
	return &uid

}

type BasicMixinType int

const (
	BasicMixinTypeBasePolicy BasicMixinType = iota
	BasicMixinTypeCreateBy
	BasicMixinTypeUpdateBy
	BasicMixinTypeCreateTime
	BasicMixinTypeUpdateTime

	basicMixinTypeCount
)

// 所有的基础混入模型
func BasicMixin(resType interface{}) []ent.Mixin {
	mixins := []ent.Mixin{
		BasePolicyMixin{ResType: util.TToS(resType)},
		CreateByMixin{},
		UpdateByMixin{},
		CreateTimeMixin{},
		UpdateTimeMixin{}}

	if len(mixins) != int(basicMixinTypeCount) {
		panic("类型数量错误")
	}

	return mixins
}

// 根据需要组合
func BasicMixinWith(resType interface{}, types ...BasicMixinType) []ent.Mixin {
	if count := len(types); count == 0 || count > int(basicMixinTypeCount) {
		panic("参数个数错误")
	}

	mixins := make([]ent.Mixin, 0)

	for _, t := range types {
		switch BasicMixinType(t) {
		case BasicMixinTypeBasePolicy:
			mixins = append(mixins, BasePolicyMixin{ResType: util.TToS(resType)})
		case BasicMixinTypeCreateBy:
			mixins = append(mixins, CreateByMixin{})
		case BasicMixinTypeUpdateBy:
			mixins = append(mixins, UpdateByMixin{})
		case BasicMixinTypeCreateTime:
			mixins = append(mixins, CreateTimeMixin{})
		case BasicMixinTypeUpdateTime:
			mixins = append(mixins, UpdateTimeMixin{})
		default:
			panic("类型错误")
		}
	}

	return mixins
}

// 混入的聚合，使得多个schema使用相同的聚合方案时更方便。
// 这尤其体现在全局的隐私策略中（ent目前没有提供全局隐私策略，只提供了全局hook)。

// NOTE : by zhangl 忘了上面注释对MixinAggregation存在的意义的解释了。先把实现封掉

// type MixinAggregation struct {
// 	mixin.Schema
// 	Mixins []ent.Mixin
// }

// func (ins MixinAggregation) Fields() []ent.Field {
// 	var all []ent.Field

// 	for _, e := range ins.Mixins {
// 		if e.Fields() != nil {
// 			all = append(all, e.Fields()...)
// 		}
// 	}

// 	return all
// }

// func (ins MixinAggregation) Edges() []ent.Edge {
// 	var all []ent.Edge

// 	for _, e := range ins.Mixins {
// 		all = append(all, e.Edges()...)
// 	}

// 	return all
// }
// func (ins MixinAggregation) Indexes() []ent.Index {
// 	var all []ent.Index

// 	for _, e := range ins.Mixins {
// 		all = append(all, e.Indexes()...)
// 	}

// 	return all
// }
// func (ins MixinAggregation) Hooks() []ent.Hook {
// 	var all []ent.Hook

// 	for _, e := range ins.Mixins {
// 		all = append(all, e.Hooks()...)
// 	}

// 	return all
// }

// func (ins MixinAggregation) Policy() ent.Policy {

// 	var allPolicy = privacy.Policy{}

// 	for _, e := range ins.Mixins {
// 		if e.Policy() != nil {

// 			allPolicy.Mutation = append(allPolicy.Mutation, e.Policy().(privacy.Policy).Mutation)
// 		}
// 		if e.Policy() != nil {

// 			allPolicy.Query = append(allPolicy.Query, e.Policy().(privacy.Policy).Query)
// 		}
// 	}

// 	if len(allPolicy.Mutation) == 0 && len(allPolicy.Query) == 0 {
// 		return nil
// 	}

// 	return allPolicy
// }

// func (ins MixinAggregation) Annotations() []schema.Annotation {
// 	var all []schema.Annotation

// 	for _, mixin := range ins.Mixins {
// 		all = append(all, mixin.Annotations()...)
// 	}

// 	return all
// }

// ****************************************************************
type CreateByMixin struct{ mixin.Schema }

const createByUserFieldName = "create_by_user"
const createByUserEdgesName = "create_by"

func (CreateByMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int(createByUserFieldName).Optional(),
	}
}

// func (CreateByMixin) Indexes() []ent.Index {
// 	return []ent.Index{
// 		index.Fields(createByUserFieldName),
// 	}
// }

func (CreateByMixin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To(createByUserEdgesName, User.Type).
			Field(createByUserFieldName).
			Unique().
			Annotations(entgql.Bind()),
	}
}

func addCreateBy(next ent.Mutator) ent.Mutator {
	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
		// 第一个账户的创建创建者是自己，没有认证信息。
		// 所以，如果出错，并且没有设置该字段，则认为出错。

		clear := m.FieldCleared(createByUserFieldName)
		if !clear {

			// 存在某些定级资源不会有创建者。例如第一个系统实例和账户
			// 暂时不在这里做判断了。
			userId := getAuthUserIdFromCtx(ctx)

			if userId != nil {
				err := m.SetField(createByUserFieldName, *userId)
				if err != nil {
					return nil, err
				}
			}
		}

		return next.Mutate(ctx, m)
	})
}

func (CreateByMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(addCreateBy, ent.OpCreate),
	}
}

// ****************************************************************
type UpdateByMixin struct{ mixin.Schema }

const updateByUserFieldName = "update_by_user"
const updateByUserEdgeName = "update_by"

func (UpdateByMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int(updateByUserFieldName).
			Optional(),
	}
}

// func (UpdateByMixin) Indexes() []ent.Index {
// 	return []ent.Index{
// 		index.Fields(updateByUserFieldName),
// 	}
// }

func (UpdateByMixin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To(updateByUserEdgeName, User.Type).
			Field(updateByUserFieldName).
			Unique().
			Annotations(entgql.Bind()),
	}
}

func addUpdateBy(next ent.Mutator) ent.Mutator {
	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {

		if m.Op().Is(ent.OpUpdate) || m.Op().Is(ent.OpUpdateOne) {
			clear := m.FieldCleared(updateByUserFieldName)
			if !clear {

				userId := getAuthUserIdFromCtx(ctx)

				if userId != nil {
					err := m.SetField(updateByUserFieldName, *userId)
					if err != nil {
						return nil, err
					}
				}
			}
		}

		return next.Mutate(ctx, m)
	})
}

func (UpdateByMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(addUpdateBy, ent.OpUpdate|ent.OpUpdateOne),
	}
}

// ****************************************************************
type CreateTimeMixin struct{ mixin.Schema }

func (CreateTimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("create_time").
			Default(time.Now).
			Immutable().
			Annotations(entgql.OrderField("CREATE_TIME")).
			SchemaType(TimeSchemaType),
	}
}

func (CreateTimeMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("create_time"),
	}
}

// ****************************************************************
type UpdateTimeMixin struct{ mixin.Schema }

func (UpdateTimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("update_time").
			Default(time.Now).
			UpdateDefault(time.Now).
			Optional().
			Nillable().
			Annotations(entgql.OrderField("UPDATE_TIME")).
			SchemaType(TimeSchemaType),
	}
}

func (UpdateTimeMixin) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("update_time"),
	}
}

//*************************************************

// 基础隐私策略：
// 检查认证信息，阻止未认证，允许所有admin.
// 其他权限

type BasePolicyMixin struct {
	mixin.Schema
	ResType string
}

func (ins BasePolicyMixin) Policy() ent.Policy {
	return nil

	// TODO 增加rule的实现，并实现当前隐私策略。

	// return privacy.Policy{
	// 	Mutation: privacy.MutationPolicy{
	// 		rule.DenyMutationRule(ins.ResType),
	// 	},
	// 	Query: privacy.QueryPolicy{
	// 		rule.QueryFilterAllRule(ins.ResType),
	// 	},
	// }
}
