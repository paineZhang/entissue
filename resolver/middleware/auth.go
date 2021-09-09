package middleware

import (
	"context"
	"fmt"

	"wing/resolver/model"

	"wing/models/ent"
	"wing/models/schema"

	"github.com/gin-gonic/gin"

	apierror "wing/resolver/error"
)

type authInfoKey struct{}

type AuthInfo struct {
	UserId      int
	AccountName string
}

// 将jwt token 单独传递进ctx，因为每次调用CheckAuthInfo都会有额外的计算。
func ginTransferJWTTokenToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 忽略错误。认为token为空。
		authInfo, _ := getAuthInfo(c.Request)
		if authInfo != nil {

			ctx := context.WithValue(c.Request.Context(), authInfoKey{}, authInfo)
			// userid在model中用于自动添加create by, update by以及操作日志信息中的操作对象信息
			ctx = schema.SetAuthUserIDToCtx(ctx, authInfo.UserId)

			c.Request = c.Request.WithContext(ctx)
		}

		c.Next()
	}
}

// 检查认证信息有效性，并将token传递进上下文中
func CheckAuthInfo(ctx context.Context, modelClient *ent.Client) error {

	authInfo, ok := ctx.Value(authInfoKey{}).(*AuthInfo)
	if !ok || authInfo == nil {
		return apierror.New(ctx, model.ErrorCodeNoAuth, "未认证")
	}

	usr, err := modelClient.User.Get(ctx, authInfo.UserId)
	if err != nil {
		return apierror.Transform(ctx, err)
	}

	if usr == nil {
		return apierror.New(ctx,
			model.ErrorCodeNotExist, fmt.Sprintf("认证账户不存在。用户ID：%d", authInfo.UserId))
	}

	if !usr.IsOnJob {
		return apierror.New(ctx, model.ErrorCodeAccountNotAllowed, "用户非在职状态，访问被阻止")
	}

	// TODO 增加token过期刷新机制。
	// 但需要增加刷新token的过程。防止token长时间有效
	// 因为token已经解开了。说明token加解密没问题。暂时就认为认证成功。
	return nil

}

func GetAuthInfo(ctx context.Context) (authInfo *AuthInfo) {
	authInfo, ok := ctx.Value(authInfoKey{}).(*AuthInfo)
	if !ok {
		return nil
	}
	return authInfo
}
