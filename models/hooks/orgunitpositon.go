package hooks

import (
	"context"
	"fmt"
	"sync"
	casbinadapter "wing/adapters/casbin"
	"wing/models/ent"
	"wing/models/ent/hook"
	"wing/models/ent/organization"
	"wing/models/ent/orgunit"

	"github.com/golobby/container/v3"
)

// TODO 等go发布带有泛型的版本，这里需要重构

type (
	OrgUnitPositionHookFunc func(ctx context.Context, m *ent.OrgUnitPositionMutation) error
	orgUnitPositionHook     struct {
		op ent.Op
		fc OrgUnitPositionHookFunc
	}

	orgUnitPositionHookHolder struct {
		Mutex     *sync.Mutex
		HookFuncs map[interface{}]orgUnitPositionHook
	}

	onDeleteOrgUnitPositionKey struct{}
)

var _orgUnitPositionHookHolder = &orgUnitPositionHookHolder{Mutex: &sync.Mutex{},
	HookFuncs: map[interface{}]orgUnitPositionHook{
		&onDeleteOrgUnitPositionKey{}: {
			ent.OpDeleteOne,
			onDeleteOrgUnitPosition,
		},
	}}

func InstallOrgUnitPositionHook(key interface{}, op ent.Op, fc OrgUnitPositionHookFunc) {
	_orgUnitPositionHookHolder.Mutex.Lock()
	defer _orgUnitPositionHookHolder.Mutex.Unlock()

	_, ok := _orgUnitPositionHookHolder.HookFuncs[key]
	if ok {
		panic("hook注册key相同")
	}
	_orgUnitPositionHookHolder.HookFuncs[key] = orgUnitPositionHook{op, fc}
}

func UninstallOrgUnitPositionHook(key interface{}) {
	_orgUnitPositionHookHolder.Mutex.Lock()
	defer _orgUnitPositionHookHolder.Mutex.Unlock()

	_, ok := _orgUnitPositionHookHolder.HookFuncs[key]
	if !ok {
		panic("不存在key")
	}
	delete(_orgUnitPositionHookHolder.HookFuncs, key)
}

func OnOrgUnitPosition(next ent.Mutator) ent.Mutator {
	return hook.OrgUnitPositionFunc(
		func(ctx context.Context, m *ent.OrgUnitPositionMutation) (ent.Value, error) {

			// 如果想得到结果，例如create时，被新创建出来的ID，需要先执行后续的变更
			// 变更生效且没有错误时，才能认为变更有结果且可用
			value, err := next.Mutate(ctx, m)
			if err != nil {
				return value, err
			}

			_orgUnitPositionHookHolder.Mutex.Lock()
			defer _orgUnitPositionHookHolder.Mutex.Unlock()
			for _, hook := range _orgUnitPositionHookHolder.HookFuncs {
				if (m.Op().Is(ent.OpCreate) && hook.op.Is(ent.OpCreate)) ||
					(m.Op().Is(ent.OpUpdate) && hook.op.Is(ent.OpUpdate)) ||
					(m.Op().Is(ent.OpUpdateOne) && hook.op.Is(ent.OpUpdateOne)) ||
					(m.Op().Is(ent.OpDelete) && hook.op.Is(ent.OpDelete)) ||
					(m.Op().Is(ent.OpDeleteOne) && hook.op.Is(ent.OpDeleteOne)) {

					err = hook.fc(ctx, m)
					if err != nil {
						return value, err
					}
				}
			}

			return value, err
		},
	)
}

func onDeleteOrgUnitPosition(ctx context.Context, m *ent.OrgUnitPositionMutation) error {

	var cb *casbinadapter.CasbinAdapter
	container.Resolve(&cb)

	id, exists := m.ID()
	if !exists {
		return fmt.Errorf("不存在的删除记录")
	}

	unitId, err := m.OldOrgUnitID(ctx)
	if err != nil {
		return err
	}

	org, err := m.Client().Organization.Query().
		Where(organization.HasUnitsWith(orgunit.ID(unitId))).First(ctx)
	if err != nil {
		return err
	}

	success, err := cb.DeleteRoleInDomain("org", org.ID, id)

	if err != nil {
		return err
	}

	if !success {
		return fmt.Errorf("删除组织角色失败败。org::%d role::%d", org.ID, id)
	}
	return nil
}
