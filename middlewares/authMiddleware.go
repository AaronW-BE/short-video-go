package middlewares

import (
	"MiniVideo/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		println("")
		log.Print("auth middleware processing...")

		authorization := context.GetHeader("Authorization")

		log.Println("token:", authorization)

		if authorization == "" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"message": "权限不足,请求失败",
			})
			return
		}

		// 查询 token
		storeToken, err := utils.Cache.Get("token")

		log.Println("stored token", storeToken)
		if err {
			context.JSON(http.StatusInternalServerError, nil)
			return
		}

		if storeToken == nil {
			context.JSON(http.StatusUnauthorized, nil)
			return
		}

		if storeToken == authorization {
			// 解析用户
			context.Set("user_id", "001")
			context.Next()
		}

		context.JSON(http.StatusUnauthorized, nil)
		return
	}
}
