package hooks

import (
	"context"
	"sync"
	"wing/models/ent"
	"wing/models/ent/hook"
)

type UserHookFunc func(ctx context.Context, user *ent.User) error
type userHook struct {
	op ent.Op
	fc UserHookFunc
}

type userHookHolder struct {
	Mutex     *sync.Mutex
	HookFuncs map[interface{}]userHook
}

var _userHookHolder = &userHookHolder{Mutex: &sync.Mutex{}, HookFuncs: make(map[interface{}]userHook)}

func InstallUserHook(key interface{}, op ent.Op, fc UserHookFunc) {
	_userHookHolder.Mutex.Lock()
	defer _userHookHolder.Mutex.Unlock()

	_, ok := _userHookHolder.HookFuncs[key]
	if ok {
		panic("hook注册key相同")
	}
	_userHookHolder.HookFuncs[key] = userHook{op, fc}
}

func UninstallUserHook(key interface{}) {
	_userHookHolder.Mutex.Lock()
	defer _userHookHolder.Mutex.Unlock()

	_, ok := _userHookHolder.HookFuncs[key]
	if !ok {
		panic("不存在key")
	}
	delete(_userHookHolder.HookFuncs, key)
}

func OnUser(next ent.Mutator) ent.Mutator {
	return hook.UserFunc(
		func(ctx context.Context, m *ent.UserMutation) (ent.Value, error) {

			// 如果想得到结果，例如create时，被新创建出来的ID，需要先执行后续的变更
			// 变更生效且没有错误时，才能认为变更有结果且可用
			value, err := next.Mutate(ctx, m)
			if err != nil {
				return value, err
			}

			user, ok := value.(*ent.User)
			if !ok {
				panic("类型错误")
			}

			_userHookHolder.Mutex.Lock()
			defer _userHookHolder.Mutex.Unlock()
			for _, hook := range _userHookHolder.HookFuncs {
				if (m.Op().Is(ent.OpCreate) && hook.op.Is(ent.OpCreate)) ||
					(m.Op().Is(ent.OpUpdate) && hook.op.Is(ent.OpUpdate)) ||
					(m.Op().Is(ent.OpUpdateOne) && hook.op.Is(ent.OpUpdateOne)) ||
					(m.Op().Is(ent.OpDelete) && hook.op.Is(ent.OpDelete)) ||
					(m.Op().Is(ent.OpDeleteOne) && hook.op.Is(ent.OpDeleteOne)) {

					// TODO ,OpUpdate是条件更新或批量更新，即更新调用的函数不是UpdateOne。所以，这里的值在UpdateOne获取不到。
					err = hook.fc(ctx, user)
					if err != nil {
						return value, err
					}
				}
			}

			return value, err
		},
	)
}
