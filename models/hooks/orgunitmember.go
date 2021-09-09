package hooks

import (
	"context"
	"fmt"
	"sync"
	"wing/models/ent"
	"wing/models/ent/hook"
	"wing/models/ent/organization"
	"wing/models/ent/orgunit"

	casbinadapter "wing/adapters/casbin"

	"github.com/golobby/container/v3"
)

type (
	OrgUnitMemberHookFunc func(ctx context.Context, newValue *ent.OrgUnitMember,
		m *ent.OrgUnitMemberMutation) error
	orgUnitMemberHook struct {
		op ent.Op
		fc OrgUnitMemberHookFunc
	}

	orgUnitMemberHookHolder struct {
		Mutex     *sync.Mutex
		HookFuncs map[interface{}]orgUnitMemberHook
	}

	onUpdateOrgUnitMemberKey struct{}
	onCreateOrgUnitMemberKey struct{}
	onDeleteOrgUnitMemberKey struct{}
)

var _orgUnitMemberHookHolder = &orgUnitMemberHookHolder{Mutex: &sync.Mutex{},
	// 初始化固定处理的hooks
	HookFuncs: map[interface{}]orgUnitMemberHook{
		&onUpdateOrgUnitMemberKey{}: {
			ent.OpUpdateOne,
			onUpdateOrgUnitMember,
		},
		&onCreateOrgUnitMemberKey{}: {
			ent.OpCreate,
			onCreateOrgUnitMember,
		},
		&onDeleteOrgUnitMemberKey{}: {
			ent.OpDeleteOne,
			onDeleteOrgUnitMember,
		},
	}}

func InstallOrgUnitMemberHook(key interface{}, op ent.Op, fc OrgUnitMemberHookFunc) {
	_orgUnitMemberHookHolder.Mutex.Lock()
	defer _orgUnitMemberHookHolder.Mutex.Unlock()

	_, ok := _orgUnitMemberHookHolder.HookFuncs[key]
	if ok {
		panic("hook注册key相同")
	}
	_orgUnitMemberHookHolder.HookFuncs[key] = orgUnitMemberHook{op, fc}
}

func UninstallOrgUnitMemberHook(key interface{}) {
	_orgUnitMemberHookHolder.Mutex.Lock()
	defer _orgUnitMemberHookHolder.Mutex.Unlock()

	_, ok := _orgUnitMemberHookHolder.HookFuncs[key]
	if !ok {
		panic("不存在key")
	}
	delete(_orgUnitMemberHookHolder.HookFuncs, key)
}

func OnOrgUnitMember(next ent.Mutator) ent.Mutator {
	return hook.OrgUnitMemberFunc(
		func(ctx context.Context, m *ent.OrgUnitMemberMutation) (ent.Value, error) {

			// 如果想得到结果，例如create时，被新创建出来的ID，需要先执行后续的变更
			// 变更生效且没有错误时，才能认为变更有结果且可用
			value, err := next.Mutate(ctx, m)
			if err != nil {
				return value, err
			}

			_orgUnitMemberHookHolder.Mutex.Lock()
			defer _orgUnitMemberHookHolder.Mutex.Unlock()

			for _, hook := range _orgUnitMemberHookHolder.HookFuncs {
				if (m.Op().Is(ent.OpCreate) && hook.op.Is(ent.OpCreate)) ||
					(m.Op().Is(ent.OpUpdate) && hook.op.Is(ent.OpUpdate)) ||
					(m.Op().Is(ent.OpUpdateOne) && hook.op.Is(ent.OpUpdateOne)) ||
					(m.Op().Is(ent.OpDelete) && hook.op.Is(ent.OpDelete)) ||
					(m.Op().Is(ent.OpDeleteOne) && hook.op.Is(ent.OpDeleteOne)) {

					err = hook.fc(ctx, value.(*ent.OrgUnitMember), m)
					if err != nil {
						return value, err
					}
				}
			}

			return value, err
		},
	)
}

func onCreateOrgUnitMember(ctx context.Context,
	value *ent.OrgUnitMember, m *ent.OrgUnitMemberMutation) error {
	_ = ctx
	_ = m

	var cb *casbinadapter.CasbinAdapter
	container.Resolve(&cb)

	err := addRoleForUserInDomain(cb,
		value.Edges.BelongToOrgUnit.Edges.BelongToOrg.ID,
		value.UserID,
		value.OrgUnitPositionID)
	if err != nil {
		return err
	}
	return nil
}

func onUpdateOrgUnitMember(ctx context.Context,
	newValue *ent.OrgUnitMember, m *ent.OrgUnitMemberMutation) error {

	var cb *casbinadapter.CasbinAdapter
	container.Resolve(&cb)

	oldOrgUnitPositionID, err := m.OldOrgUnitPositionID(ctx)
	if err != nil {
		return err
	}
	oldUserID, err := m.OldUserID(ctx)
	if err != nil {
		return err
	}

	if newValue.UserID != oldUserID || newValue.OrgUnitPositionID != oldOrgUnitPositionID {

		err := deleteRoleForUserInDomain(cb,
			newValue.Edges.BelongToOrgUnit.Edges.BelongToOrg.ID,
			oldUserID,
			oldOrgUnitPositionID)
		if err != nil {
			return err
		}
		err = addRoleForUserInDomain(cb,
			newValue.Edges.BelongToOrgUnit.Edges.BelongToOrg.ID,
			newValue.UserID,
			newValue.OrgUnitPositionID)
		if err != nil {
			return err
		}
	}

	return nil
}

func onDeleteOrgUnitMember(ctx context.Context,
	newValue *ent.OrgUnitMember, m *ent.OrgUnitMemberMutation) error {
	_ = newValue

	var cb *casbinadapter.CasbinAdapter
	container.Resolve(&cb)

	unitId, err := m.OldOrgUnitID(ctx)
	if err != nil {
		return err
	}

	org, err := m.Client().Organization.Query().
		Where(organization.HasUnitsWith(orgunit.ID(unitId))).First(ctx)
	if err != nil {
		return err
	}

	oldOrgUnitPositionID, err := m.OldOrgUnitPositionID(ctx)
	if err != nil {
		return err
	}
	oldUserID, err := m.OldUserID(ctx)
	if err != nil {
		return err
	}
	err = deleteRoleForUserInDomain(cb, org.ID, oldUserID, oldOrgUnitPositionID)
	if err != nil {
		return err
	}
	return nil
}

func addRoleForUserInDomain(cb *casbinadapter.CasbinAdapter, orgId int, userId int, roleId int) error {

	// TODO 把这些hook改成接口后注入。
	success, err := cb.AddRoleForUserInDomain("org", orgId, userId, roleId)

	if err != nil {
		return err
	}

	if !success {
		return fmt.Errorf("添加用户角色失败。%d %d %d", orgId, userId, roleId)
	}
	return nil
}

func deleteRoleForUserInDomain(cb *casbinadapter.CasbinAdapter, orgId int, userId int, roleId int) error {

	success, err := cb.DeleteRoleForUserInDomain("org", orgId, userId, roleId)

	if err != nil {
		return err
	}

	if !success {
		return fmt.Errorf("删除用户角色失败。%d %d %d", orgId, userId, roleId)
	}
	return nil
}
