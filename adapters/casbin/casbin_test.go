package casbinadapter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRole(t *testing.T) {
	ca, err := New()
	assert.NoError(t, err)

	result, err := ca.clearInDomain("org", 1)
	assert.True(t, result)
	assert.NoError(t, err)

	// 简单的增加删除
	result, err = ca.AddRoleForUserInDomain("org", 1, 1, 1)
	assert.True(t, result)
	assert.NoError(t, err)
	result, err = ca.AddRoleForUserInDomain("org", 1, 1, 1)
	assert.False(t, result)
	assert.NoError(t, err)

	roles := ca.GetRolesInDomain("org", 1)
	assert.Equal(t, len(roles), 1)
	assert.Equal(t, roles[0], 1)

	roles = ca.GetRolesForUserInDomain("org", 1, 1)
	assert.Equal(t, len(roles), 1)
	assert.Equal(t, roles[0], 1)

	result, err = ca.DeleteRoleForUserInDomain("org", 1, 1, 1)
	assert.True(t, result)
	assert.NoError(t, err)
	roles = ca.GetRolesForUserInDomain("org", 1, 1)
	assert.Equal(t, len(roles), 0)

	ca.clearInDomain("org", 1)

	generateGeneralRoleCase(t, ca, 1)

	// 所有角色
	roles = ca.GetRolesInDomain("org", 1)
	assert.Equal(t, len(roles), 2)

	// 用户的角色
	roles = ca.GetRolesForUserInDomain("org", 1, 1)
	assert.Equal(t, len(roles), 2)
	roles = ca.GetRolesForUserInDomain("org", 1, 2)
	assert.Equal(t, len(roles), 1)
	roles = ca.GetRolesForUserInDomain("org", 1, 3)
	assert.Equal(t, len(roles), 1)

	// 角色的用户
	users := ca.GetUsersFromRoleInDomain("org", 1, 1)
	assert.Equal(t, len(users), 2)
	users = ca.GetUsersFromRoleInDomain("org", 1, 2)
	assert.Equal(t, len(users), 2)

	// 删除单独角色
	result, err = ca.DeleteRoleInDomain("org", 1, 1)
	assert.True(t, result)
	assert.NoError(t, err)
	roles = ca.GetRolesInDomain("org", 1)
	assert.Equal(t, len(roles), 1)
	assert.Equal(t, roles[0], 2)

	// 删除所有角色
	result, err = ca.DeleteRolesInDomain("org", 1)
	assert.True(t, result)
	assert.NoError(t, err)
	roles = ca.GetRolesInDomain("org", 1)
	assert.Equal(t, len(roles), 0)

	result, err = ca.clearInDomain("org", 1)
	assert.True(t, result)
	assert.NoError(t, err)

	// 通过用户删除角色
	generateGeneralRoleCase(t, ca, 1)
	generateGeneralRoleCase(t, ca, 2)
	result, err = ca.DeleteRolesForUserInDomain("org", 1, 1)
	assert.True(t, result)
	assert.NoError(t, err)

	roles = ca.GetRolesForUserInDomain("org", 1, 1)
	assert.Equal(t, len(roles), 0)
	roles = ca.GetRolesForUserInDomain("org", 2, 1)
	assert.Equal(t, len(roles), 2)

	// 通过域删角色
	result, err = ca.DeleteRolesInDomain("org", 1)
	assert.True(t, result)
	assert.NoError(t, err)
	roles = ca.GetRolesInDomain("org", 1)
	assert.Equal(t, len(roles), 0)
	roles = ca.GetRolesInDomain("org", 2)
	assert.Equal(t, len(roles), 2)
	result, err = ca.DeleteRolesInDomain("org", 2)
	assert.True(t, result)
	assert.NoError(t, err)
	roles = ca.GetRolesInDomain("org", 2)
	assert.Equal(t, len(roles), 0)

	result, err = ca.clearInDomain("org", 1)
	assert.True(t, result)
	assert.NoError(t, err)

	result, err = ca.clearInDomain("org", 2)
	assert.True(t, result)
	assert.NoError(t, err)
}

func TestPermissions(t *testing.T) {

	ca, err := New()
	assert.NoError(t, err)

	result, err := ca.clearInDomain("org", 1)
	assert.True(t, result)
	assert.NoError(t, err)

	generateGeneralRoleCase(t, ca, 1)

	result, err = ca.AddPolicyForRoleInDomain("org", 1, 1, 1, ActionRead, EffectAllow)
	assert.True(t, result)
	assert.NoError(t, err)
	result, err = ca.AddPolicyForRoleInDomain("org", 1, 2, 1, ActionWrite, EffectAllow)
	assert.True(t, result)
	assert.NoError(t, err)

	result, err = ca.CheckPermission("org", 1, 1, 1, ActionRead)
	assert.True(t, result)
	assert.NoError(t, err)
	result, err = ca.CheckPermission("org", 1, 1, 1, ActionWrite)
	assert.True(t, result)
	assert.NoError(t, err)

	result, err = ca.CheckPermission("org", 1, 2, 1, ActionRead)
	assert.False(t, result)
	assert.NoError(t, err)
	result, err = ca.CheckPermission("org", 1, 2, 1, ActionWrite)
	assert.True(t, result)
	assert.NoError(t, err)

	result, err = ca.CheckPermission("org", 1, 3, 1, ActionRead)
	assert.True(t, result)
	assert.NoError(t, err)
	result, err = ca.CheckPermission("org", 1, 3, 1, ActionWrite)
	assert.False(t, result)
	assert.NoError(t, err)

	result, err = ca.DeletePolicyForRoleInDomain("org", 1, 1, 1, ActionRead, EffectAllow)
	assert.True(t, result)
	assert.NoError(t, err)

	result, err = ca.CheckPermission("org", 1, 1, 1, ActionRead)
	assert.False(t, result)
	assert.NoError(t, err)
	result, err = ca.CheckPermission("org", 1, 3, 1, ActionRead)
	assert.False(t, result)
	assert.NoError(t, err)

	result, err = ca.DeletePoliciesForRoleInDomain("org", 1, 2)
	assert.True(t, result)
	assert.NoError(t, err)

	result, err = ca.CheckPermission("org", 1, 1, 1, ActionWrite)
	assert.False(t, result)
	assert.NoError(t, err)
	result, err = ca.CheckPermission("org", 1, 2, 1, ActionWrite)
	assert.False(t, result)
	assert.NoError(t, err)

	result, err = ca.clearInDomain("org", 1)
	assert.True(t, result)
	assert.NoError(t, err)
}

func generateGeneralRoleCase(t *testing.T, ca *CasbinAdapter, domainId int) {
	result, err := ca.AddRoleForUserInDomain("org", domainId, 1, 1)
	assert.True(t, result)
	assert.NoError(t, err)
	result, err = ca.AddRoleForUserInDomain("org", domainId, 1, 2)
	assert.True(t, result)
	assert.NoError(t, err)
	result, err = ca.AddRoleForUserInDomain("org", domainId, 2, 2)
	assert.True(t, result)
	assert.NoError(t, err)
	result, err = ca.AddRoleForUserInDomain("org", domainId, 3, 1)
	assert.True(t, result)
	assert.NoError(t, err)
}
