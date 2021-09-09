package rules

import (
	"context"
	"wing/models/ent"
	"wing/models/ent/privacy"

	"entgo.io/ent/entql"
	"github.com/golobby/container/v3"
)

func QueryFilterAllRule(resType string) privacy.QueryRule {

	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {

		// NOTE 这里不知道是哪个，暂时按照任何角色有权限，就允许处理
		// TODO 增加具体项目所属资源的处理。
		var pc PermissionChecker
		err := container.Resolve(&pc)
		if err != nil {
			panic(err)
		}
		allow, err := pc.CheckPermission(ctx, resType, ActR)
		if err != nil {
			return privacy.Denyf("%s", err.Error())
		}

		// 需要有所有实体都有id字段，数据id记录不为负，
		// 利用这个特性，当请求者没有相关读权限。过滤掉所有记录，而不是返回错误。
		if !allow {
			f.Where(entql.FieldEQ("id", -1))
		}

		return privacy.Skip
	})
}

func DenyMutationRule(resType string) privacy.MutationRule {
	return privacy.MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {

		var act Act

		if m.Op().Is(ent.OpCreate) {
			act = ActC
		} else if m.Op().Is(ent.OpUpdate) || m.Op().Is(ent.OpUpdateOne) {
			act = ActU
		} else if m.Op().Is(ent.OpDelete) || m.Op().Is(ent.OpDeleteOne) {
			act = ActD
		}

		var pc PermissionChecker
		err := container.Resolve(&pc)
		if err != nil {
			panic(err)
		}

		// NOTE TODO 增加资源所属项目的的处理,这里不知道是哪个(在ctx中增加域信息)，暂时按照任何角色有权限，就允许处理
		allow, err := pc.CheckPermission(ctx, resType, act)
		if err != nil {
			return privacy.Denyf("%s", err.Error())
		}

		if !allow {
			//TODO 增加专门的错误信息
			return privacy.Deny
		}

		return privacy.Skip
	})
}
