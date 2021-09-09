package middleware

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

const GinContextKey = "GinContextKey"

// 将gin上下文装入前系统上下文中，以便在gqlgen中可以得到gin上下文
func ginContextToContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.WithValue(c.Request.Context(), GinContextKey, c)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

// 从系统上下文得到gin上下文
func ginContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value(GinContextKey)
	if ginContext == nil {
		err := fmt.Errorf("无法丛上下文中得到gin的上下文")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin Context类型错误")
		return nil, err
	}
	return gc, nil
}
