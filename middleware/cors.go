package middleware

import (
	"github.com/gin-gonic/gin"

	global "changian.com/jubo/startup"
)

func CorsMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {

		context.Writer.Header().Set("Content-Type", "application/json")
		context.Writer.Header().Set("Access-Control-Allow-Origin", global.Config.AccessControlAllowOrigin)
		context.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, DELETE")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, Authorization, accept, origin, Cache-Control, X-CSRF-Token, X-Max, X-Requested-With")
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		context.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(204)
		} else {
			context.Next()
		}
	}
}
