package middlewares

import "github.com/gin-gonic/gin"

func Api() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("content-type", "application/json")
		context.Next()
	}
}
