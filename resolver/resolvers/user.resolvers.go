package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	ldapadapter "wing/adapters/ldap"
	"wing/models/ent"
	"wing/models/ent/jobhistory"
	"wing/models/ent/user"
	"wing/models/hooks"
	apierror "wing/resolver/error"
	generated1 "wing/resolver/generated"
	"wing/resolver/middleware"
	model1 "wing/resolver/model"
	"wing/util"
)

func (r *mutationResolver) UserMutation(ctx context.Context) (*ent.UserMutation, error) {
	return &ent.UserMutation{}, middleware.CheckAuthInfo(ctx, ent.FromContext(ctx))
}

func (r *queryResolver) UserQuery(ctx context.Context) (*ent.UserQuery, error) {
	return &ent.UserQuery{}, middleware.CheckAuthInfo(ctx, ent.FromContext(ctx))
}

func (r *subscriptionResolver) UserAdded(ctx context.Context) (<-chan *model1.UserSubscriptionInfo, error) {
	// 订阅 没有上一级的subscriptionResolver，订阅的授权认证只能在这里。
	if err := middleware.CheckAuthInfo(ctx, ent.FromContext(ctx)); err != nil {
		return nil, err
	}

	// TODO ，增加推送订阅的权限？？？

	ch := make(chan *model1.UserSubscriptionInfo)
	go func() {
		hooks.InstallUserHook(ch, ent.OpCreate, func(ctx context.Context, user *ent.User) error {
			_ = ctx
			ch <- &model1.UserSubscriptionInfo{ID: user.ID}
			return nil
		})

		defer hooks.UninstallUserHook(ch)

		// 直到ctx被cancel
		<-ctx.Done()
	}()
	return ch, nil
}

func (r *subscriptionResolver) UserUpdated(ctx context.Context, filters *ent.UserWhereInput) (<-chan *model1.UserSubscriptionInfo, error) {
	// 订阅 没有上一级的subscriptionResolver，订阅的授权认证只能在这里。
	if err := middleware.CheckAuthInfo(ctx, ent.FromContext(ctx)); err != nil {
		return nil, err
	}

	// 检查订阅的筛选条件实现
	// TODO 单独增加订阅的筛选条件？？ent的UserWhereInput是生成查询谓词的，这里没法直接用，
	//      如果编写代码做检查，太琐碎了。或者按需提供接口。现在看，离入职的状态变更是最需要的。

	// TODO ，增加推送订阅的权限？？？

	ch := make(chan *model1.UserSubscriptionInfo)
	// 创建一个副本在goroutine中用
	go func() {
		hooks.InstallUserHook(ch, ent.OpUpdate|ent.OpDeleteOne, func(ctx context.Context, user *ent.User) error {
			_ = ctx
			if filters == nil {
				ch <- &model1.UserSubscriptionInfo{ID: user.ID}
			} else {
				if filters.ID != nil && *filters.ID == user.ID {
					ch <- &model1.UserSubscriptionInfo{ID: user.ID}
				} else {
					panic("服务未实现筛选的订阅条件")
				}
			}
			return nil

		})

		defer hooks.UninstallUserHook(ch)

		// 直到ctx被cancel
		<-ctx.Done()

	}()
	return ch, nil
}

func (r *userMutationResolver) Create(ctx context.Context, obj *ent.UserMutation, input ent.CreateUserInput, password string) (*ent.User, error) {
	model := ent.FromContext(ctx)

	// 检查LDAP中的账户名是否有重复
	find, err := r.Ldap.Find(input.AccountName)
	if err != nil {
		return nil, apierror.Transform(ctx, err)
	}

	if find {
		return nil, apierror.New(ctx, model1.ErrorCodeAlreadyExist, "系统存在相同账户名")
	}

	// 检查当前账户系统中账户名是否有重复
	exsit, err := model.User.Query().Where(user.AccountNameEQ(input.AccountName)).All(ctx)
	if err != nil {
		return nil, apierror.Transform(ctx, err)
	}
	if len(exsit) != 0 {
		return nil, apierror.New(ctx, model1.ErrorCodeAlreadyExist, "ldap服务存在相同账户名")
	}

	c := model.User.Create().SetInput(input)

	// 处理在职状态为真，增加工作履历
	if input.IsOnJob != nil && *input.IsOnJob {
		jh, err := model.JobHistory.Create().SetJobEntryLeaveType(jobhistory.JobEntryLeaveTypeEntry).Save(ctx)
		if err != nil {
			return nil, apierror.Transform(ctx, err)
		}
		c = c.AddJobHistories(jh)
	}

	// 当前系统中添加账户
	usr, err := c.Save(ctx)
	if err != nil {
		return nil, apierror.Transform(ctx, err)
	}

	// LDAP 中添加账户
	// NOTE ldap没有回滚机制，在所有都成功后再操作
	err = r.Ldap.Add(&ldapadapter.PersonModel{
		AccountName: input.AccountName,
		Password:    password,
		DisplayName: input.DisplayName,
		FamilyName:  input.FamilyName,
		GivenName:   input.GivenName,
		Mail:        input.IntranetWorkEmail,
	})

	if err != nil {
		return nil, apierror.Transform(ctx, err)
	}

	return usr, nil
}

func (r *userMutationResolver) Update(ctx context.Context, obj *ent.UserMutation, userID int, input ent.UpdateUserInput) (bool, error) {
	model := ent.FromContext(ctx)

	u := model.User.Update().SetInput(input)

	// 检查在职状态变更。处理工作履历
	// TODO 添加、变更中对工作履历的处理放到hook中？
	if input.IsOnJob != nil {
		usr, err := model.User.Get(ctx, userID)
		if err != nil {
			return false, apierror.Transform(ctx, err)
		}

		if usr.IsOnJob != *input.IsOnJob {
			jh, err := model.JobHistory.Create().
				SetJobEntryLeaveType(jobhistory.JobEntryLeaveTypeEntry).Save(ctx)
			if err != nil {
				return false, apierror.Transform(ctx, err)
			}

			u = u.AddJobHistories(jh)

			// 如果更新为离职，生成一个随机密码，并修改。
			// TODO 需要配合jwt token的刷新需求，例如，离职后，如果相关用户再与系统交互，那么需要刷新token。
			//      jwt中可以保存用户密码，在刷新时，重新认证。这样就会导致失败。并提示重新登录。
			//		那么需要涉及一个如何判断是否需要立即刷新token的机制。
			if !*input.IsOnJob {
				err = r.Ldap.ResetPassword(usr.AccountName, util.NewRandomString(24))
				if err != nil {
					return false, apierror.Transform(ctx, err)
				}
			}
		}
	}

	err := u.Exec(ctx)
	if err != nil {
		return false, apierror.Transform(ctx, err)
	}
	return true, nil
}

func (r *userMutationResolver) ChangePassword(ctx context.Context, obj *ent.UserMutation, oldPassword string, newPassword string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userMutationResolver) ResetPassword(ctx context.Context, obj *ent.UserMutation, accountID *int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userQueryResolver) Users(ctx context.Context, obj *ent.UserQuery, after *ent.Cursor, first *int, before *ent.Cursor, last *int, filters *ent.UserWhereInput, orders []*ent.UserOrder) (*ent.UserConnection, error) {
	orderOptions := make([]ent.UserPaginateOption, 0)

	for _, order := range orders {
		orderOptions = append(orderOptions, ent.WithUserOrder(order))
	}
	orderOptions = append(orderOptions, ent.WithUserFilter(filters.Filter))

	cnn, err := ent.FromContext(ctx).User.Query().
		Paginate(ctx, after, first, before, last, orderOptions...)

	if err != nil {
		return nil, apierror.Transform(ctx, err)
	}

	return cnn, nil
}

// UserMutation returns generated1.UserMutationResolver implementation.
func (r *Resolver) UserMutation() generated1.UserMutationResolver { return &userMutationResolver{r} }

// UserQuery returns generated1.UserQueryResolver implementation.
func (r *Resolver) UserQuery() generated1.UserQueryResolver { return &userQueryResolver{r} }

type userMutationResolver struct{ *Resolver }
type userQueryResolver struct{ *Resolver }
