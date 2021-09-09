package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func customLog() gin.HandlerFunc {

	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		return fmt.Sprintf("[GIN]: %s | %s | %s | %s | %s | %d | %s\n",
			param.TimeStamp.Format(time.RFC3339),
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
		)
	})
}
