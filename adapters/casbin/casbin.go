package casbinadapter

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	entadapter "github.com/casbin/ent-adapter"
	_ "github.com/go-sql-driver/mysql"

	"github.com/emirpasic/gods/sets/hashset"
)

// TODO 从环境变量或安全文件中去读
const (
	dbAddressDefault      = "127.0.0.1"
	dbPortDefault         = "3306"
	dbUserNameDefault     = "wing"
	dbUserPasswordDefault = "wing"
	dbDatabaseNameDefault = "wing"
)

//go:embed casbin-model.conf
var casbinModel string

type (
	Action string

	Effect        string
	CasbinAdapter struct {
		enforcer *casbin.Enforcer
	}

	Policy struct {
		DomainID   int
		RoleID     int
		ResourceID int
		Action     Action
		Effect     Effect
	}
)

const (
	ActionRead  Action = "read"
	ActionWrite Action = "write"

	EffectAllow Effect = "allow"
	EffectDeny  Effect = "deny"
)

var ActionAll = []Action{ActionRead, ActionWrite}

const (
	prefixToken = ":"
	rolePrefix  = "role" + prefixToken
	subPrefix   = "user" + prefixToken
	objPrefix   = "res" + prefixToken
)

func createDomainId(prefix string, id int) string {
	return fmt.Sprintf("%s%d", prefix+prefixToken, id)
}

func createRoleId(id int) string {
	return fmt.Sprintf("%s%d", rolePrefix, id)
}

func createUserId(id int) string {
	return fmt.Sprintf("%s%d", subPrefix, id)
}

func createResourceId(id int) string {
	return fmt.Sprintf("%s%d", objPrefix, id)
}

func New() (cb *CasbinAdapter, err error) {

	cb = &CasbinAdapter{}

	m, err := model.NewModelFromString(casbinModel)
	if err != nil {
		return nil, err
	}

	adapter, err := entadapter.NewAdapter("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		dbUserNameDefault, dbUserPasswordDefault, dbAddressDefault, dbPortDefault, dbDatabaseNameDefault))
	if err != nil {
		return nil, err
	}

	cb.enforcer, err = casbin.NewEnforcer(m, adapter)
	if err != nil {
		return nil, err
	}

	cb.enforcer.LoadPolicy()

	return cb, nil
}

func (ins *CasbinAdapter) AddRoleForUserInDomain(domainPrefix string, domainId int,
	userId int, roleId int) (bool, error) {
	return ins.enforcer.AddRoleForUserInDomain(createUserId(userId),
		createRoleId(roleId), createDomainId(domainPrefix, domainId))
}

func (ins *CasbinAdapter) DeleteRoleForUserInDomain(domainPrefix string, domainId int,
	userId int, roleId int) (bool, error) {
	return ins.enforcer.DeleteRoleForUserInDomain(createUserId(userId),
		createRoleId(roleId), createDomainId(domainPrefix, domainId))
}

func (ins *CasbinAdapter) DeleteRolesForUserInDomain(domainPrefix string, domainId int, userId int) (bool, error) {
	return ins.enforcer.RemoveFilteredGroupingPolicy(0,
		createUserId(userId), "", createDomainId(domainPrefix, domainId))
}

func (ins *CasbinAdapter) DeleteRolesForUser(userId int) (bool, error) {
	return ins.enforcer.DeleteUser(createUserId(userId))
}

func (ins *CasbinAdapter) GetRolesForUserInDomain(domainPrefix string, domainId int, userId int) []int {

	roles := ins.enforcer.GetRolesForUserInDomain(createUserId(userId), createDomainId(domainPrefix, domainId))
	roleIds := make([]int, len(roles))
	for idx, role := range roles {
		roleId, err := strconv.Atoi(strings.Split(role, prefixToken)[1])
		if err != nil {
			panic("解析错误")
		}
		roleIds[idx] = roleId
	}

	return roleIds
}

func (ins *CasbinAdapter) GetUsersFromRoleInDomain(domainPrefix string, domainId int, roleId int) []int {
	users, err := ins.enforcer.GetRoleManager().GetUsers(createRoleId(roleId), createDomainId(domainPrefix, domainId))
	if err != nil {
		return []int{}
	}

	userIds := make([]int, len(users))
	for idx, user := range users {
		userId, err := strconv.Atoi(strings.Split(user, prefixToken)[1])
		if err != nil {
			panic("解析错误")
		}
		userIds[idx] = userId
	}
	return userIds
}

func (ins *CasbinAdapter) DeleteRoleInDomain(domainPrefix string, domainId int, roleId int) (bool, error) {
	_, err := ins.enforcer.RemoveFilteredGroupingPolicy(1, createRoleId(roleId), createDomainId(domainPrefix, domainId))

	if err != nil {
		return false, err
	}

	// 可能还没有产生权限策略，会返回false
	_, err = ins.enforcer.RemoveFilteredPolicy(0, createRoleId(roleId), createDomainId(domainPrefix, domainId))
	if err != nil {
		return false, err
	}

	return true, nil
}

func (ins *CasbinAdapter) DeleteRolesInDomain(domainPrefix string, domainId int) (bool, error) {
	_, err := ins.enforcer.RemoveFilteredGroupingPolicy(2, createDomainId(domainPrefix, domainId))
	if err != nil {
		return false, err
	}

	// 可能还没有产生权限策略，会返回false
	_, err = ins.enforcer.RemoveFilteredPolicy(1, createDomainId(domainPrefix, domainId))
	if err != nil {
		return false, err
	}

	return true, nil
}

func (ins *CasbinAdapter) GetRolesInDomain(domainPrefix string, domainId int) []int {

	set := hashset.New()

	// NOTE 这个方法会给出g的所有行的role结果，所以有重复的。通过集合去重
	policies := ins.enforcer.GetFilteredGroupingPolicy(2, createDomainId(domainPrefix, domainId))

	for _, policy := range policies {
		roleId, err := strconv.Atoi(strings.Split(policy[1], prefixToken)[1])
		if err != nil {
			panic("解析错误")
		}
		set.Add(roleId)
	}

	roleIds := make([]int, set.Size())

	for idx, roleId := range set.Values() {
		roleIds[idx] = roleId.(int)
	}

	return roleIds
}

func (ins *CasbinAdapter) AddPolicyForRoleInDomain(domainPrefix string, domainId int,
	roleId int, resourceId int, action Action, effect Effect) (bool, error) {
	return ins.enforcer.AddPermissionForUser(createRoleId(roleId),
		createDomainId(domainPrefix, domainId), createResourceId(resourceId), string(action), string(effect))
}

func (ins *CasbinAdapter) DeletePolicyForRoleInDomain(domainPrefix string, domainId int,
	roleId int, resourceId int, action Action, effect Effect) (bool, error) {
	return ins.enforcer.DeletePermissionForUser(createRoleId(roleId),
		createDomainId(domainPrefix, domainId), createResourceId(resourceId), string(action), string(effect))
}

func (ins *CasbinAdapter) DeletePoliciesForRoleInDomain(domainPrefix string, domainId int, roleId int) (bool, error) {
	return ins.enforcer.RemoveFilteredPolicy(0, createRoleId(roleId), createDomainId(domainPrefix, domainId))
}

func (ins *CasbinAdapter) DeletePoliciesForResourceInDomain(domainPrefix string, domainId int, resourceId int) (bool, error) {
	return ins.enforcer.RemoveFilteredPolicy(1, createDomainId(domainPrefix, domainId), createResourceId(resourceId))
}

func (ins *CasbinAdapter) DeletePoliciesInDomain(domainPrefix string, domainId int) (bool, error) {
	return ins.enforcer.RemoveFilteredPolicy(1, createDomainId(domainPrefix, domainId))
}

func (ins *CasbinAdapter) GetPoliciesForRoleInDomain(domainPrefix string, domainId int, roleId int) []Policy {
	ps := ins.enforcer.GetPermissionsForUserInDomain(createRoleId(roleId), createDomainId(domainPrefix, domainId))

	return convertPolicy(ps)
}

func (ins *CasbinAdapter) GetPoliciesForUserInDomain(domainPrefix string, domainId int, userId int) []Policy {
	ps := ins.enforcer.GetFilteredPolicy(0, createUserId(userId), createDomainId(domainPrefix, domainId))

	return convertPolicy(ps)

}

func (ins *CasbinAdapter) GetPoliciesForResourceInDomain(domainPrefix string, domainId int, resourceId int) []Policy {

	ps := ins.enforcer.GetFilteredPolicy(1, createDomainId(domainPrefix, domainId), createResourceId(resourceId))

	return convertPolicy(ps)
}

func (ins *CasbinAdapter) UpdatePoliciesForRoleInDomain(domainPrefix string, domainId int,
	roleId int, polices []Policy) error {

	pps := make([][]string, len(polices))

	domain := createDomainId(domainPrefix, domainId)
	role := createRoleId(roleId)

	for idx, p := range polices {
		pps[idx] = make([]string, 5)

		pps[idx][0] = domain
		pps[idx][1] = role
		pps[idx][2] = createResourceId(p.ResourceID)
		pps[idx][3] = string(p.Action)
		pps[idx][4] = string(p.Effect)
	}

	success, err := ins.enforcer.UpdateFilteredPolicies(pps, 0, role, domain)
	if err != nil {
		return err
	}

	if !success {
		return fmt.Errorf("casbin更新策略失败")
	}

	return nil
}

func convertPolicy(ps [][]string) []Policy {
	ret := make([]Policy, 0)

	for _, p := range ps {
		roleID, err := strconv.Atoi(strings.Split(p[0], prefixToken)[1])
		if err != nil {
			panic(err)
		}

		domainID, err := strconv.Atoi(strings.Split(p[1], prefixToken)[1])
		if err != nil {
			panic(err)
		}
		resourceID, err := strconv.Atoi(strings.Split(p[2], prefixToken)[1])
		if err != nil {
			panic(err)
		}

		ret = append(ret, Policy{
			RoleID:     roleID,
			DomainID:   domainID,
			ResourceID: resourceID,
			Action:     Action(p[3]),
			Effect:     Effect(p[4]),
		})
	}

	return ret
}

func (ins *CasbinAdapter) CheckPermission(domainPrefix string, domainId int, userId int,
	resourceId int, action Action) (bool, error) {
	return ins.enforcer.Enforce(createUserId(userId),
		createDomainId(domainPrefix, domainId), createResourceId(resourceId), string(action))
}

//  clearInDomain clear all。for unittest
func (ins *CasbinAdapter) clearInDomain(domainPrefix string, domainId int) (bool, error) {
	// TODO 类似这种连续调用操作都不是原子的，但是没找到批处理的接口实现。
	_, err := ins.DeletePoliciesInDomain(domainPrefix, domainId)
	if err != nil {
		return false, err
	}
	_, err = ins.DeleteRolesInDomain(domainPrefix, domainId)
	if err != nil {
		return false, err
	}
	return true, nil
}
