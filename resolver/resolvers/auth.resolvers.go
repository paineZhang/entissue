package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"
	ldapadapters "wing/adapters/ldap"
	"wing/models/ent"
	"wing/models/ent/auth"
	"wing/models/ent/privacy"
	"wing/models/ent/user"
	apierror "wing/resolver/error"
	"wing/resolver/generated"
	"wing/resolver/middleware"
	"wing/resolver/model"
)

func (r *authMutationResolver) Login(ctx context.Context, obj *ent.AuthMutation, accountName string, password string) (*model.LoginResponse, error) {
	// TODO 第一个登录的用户，给于所有权限？？？
	//      增加全局的角色，系统管理员和普通用户，系统管理员可以创建、编辑企业架构业务域的相关内容，以及系统设置项。
	//      其他的不可以给，这给第一次系统启动处理用户权限问题。

	exist, err := r.Ldap.Find(accountName)
	if err != nil {
		return nil, apierror.Transform(ctx, err)
	}

	if !exist {
		return nil, apierror.New(ctx, model.ErrorCodeNotExist, "不存在的账户名")
	}

	err = r.Ldap.Authentication(accountName, password)
	if err != nil {

		if ldapadapters.IsAuthUnKnown(err) {
			return nil, apierror.New(ctx, model.ErrorCodeWrongPassword, "密码错误")
		} else {
			return nil, apierror.Transform(ctx, err)
		}
	}

	txModel := ent.FromContext(ctx)

	// 忽略隐私策略
	ctx = privacy.DecisionContext(ctx, privacy.Allow)

	userResult, err := txModel.User.Query().Where(user.AccountNameEQ(accountName)).First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			// LDAP中有，但当前账户系统中没有，同步过来。
			userModel, err := r.Ldap.Get(accountName)
			if err != nil {
				return nil, apierror.Transform(ctx, err)
			}
			userResult, err = txModel.User.Create().SetAccountName(userModel.AccountName).
				SetFamilyName(userModel.FamilyName).SetGivenName(userModel.GivenName).
				SetIntranetWorkEmail(userModel.Mail).SetDisplayName(userModel.DisplayName).
				SetStaffType(user.StaffTypeRegular).
				SetIsOnJob(true).
				Save(ctx)
			if err != nil {
				return nil, apierror.Transform(ctx, err)
			}
		} else {

			return nil, apierror.Transform(ctx, err)
		}
	}

	_, err = txModel.Auth.Create().
		SetAuthType(auth.AuthTypeLogin).
		SetBelongToID(userResult.ID).Save(ctx)
	if err != nil {
		return nil, apierror.Transform(ctx, err)
	}

	loginAt := time.Now()

	token, err := middleware.CreateJWTToken(userResult.ID, userResult.AccountName, loginAt)
	if err != nil {
		return nil, apierror.Transform(ctx, err)
	}

	return &model.LoginResponse{Jwt: token, CurUser: userResult}, nil
}

func (r *authMutationResolver) RefreshToken(ctx context.Context, obj *ent.AuthMutation) (*model.RefreshTokenResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) AuthMutation(ctx context.Context) (*ent.AuthMutation, error) {
	// 忽略认证状态。
	return &ent.AuthMutation{}, nil
}

// AuthMutation returns generated.AuthMutationResolver implementation.
func (r *Resolver) AuthMutation() generated.AuthMutationResolver { return &authMutationResolver{r} }

type authMutationResolver struct{ *Resolver }
