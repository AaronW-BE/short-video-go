package middlewares

import (
	"MiniVideo/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		log.Print("auth middleware processing...")

		authorization := context.GetHeader("Authorization")

		if authorization == "" {
			utils.Response401(context, "请求失败，需要登录")
			return
		}

		// 查询 token
		storeToken, err := utils.Cache.Get("token:" + authorization)

		if !err || storeToken == nil {
			utils.Response500(context, "token无效")
			return
		}

		// 解析token
		jwtInstance := utils.NewJWT()
		claims, err2 := jwtInstance.ParseToken(authorization)
		if err2 != nil {
			utils.Response401(context, "token无效，授权失败")
			return
		}
		if claims.ID == storeToken {
			context.Set("claims", claims)
			context.Next()
			return
		}
		utils.Response401(context, "")
		return
	}
}
