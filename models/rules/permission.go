package rules

import (
	"context"

	casbinadapter "wing/adapters/casbin"
	"wing/resolver/middleware"

	"github.com/golobby/container/v3"
)

// 权限动作类型
type Act string

// 权限动作枚举
const (
	ActC Act = "C"
	ActR Act = "R"
	ActU Act = "U"
	ActD Act = "D"
)

type PermissionChecker struct {
}

func (c PermissionChecker) CheckPermission(ctx context.Context,
	resType string, act Act) (bool, error) {
	// TODO: 逻辑不对，等域的处理加进来再补全
	var casbin *casbinadapter.CasbinAdapter
	err := container.Transient(&casbin)
	if err != nil {
		panic(err)
	}
	var cbact casbinadapter.Action

	if act == ActC || act == ActU || act == ActD {
		cbact = casbinadapter.ActionWrite
	} else if act == ActR {
		cbact = casbinadapter.ActionRead
	} else {
		panic("操作枚举类型不对")
	}

	authInfo := middleware.GetAuthInfo(ctx)
	if authInfo == nil {
		return false, nil
	}

	// TODO 逻辑不对，得改
	return casbin.CheckPermission("org", 1, authInfo.UserId, 1, cbact)
}
