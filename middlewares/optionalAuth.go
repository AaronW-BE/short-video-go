package middlewares

import (
	"MiniVideo/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func OptionalAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		log.Print("optional auth middleware processing...")

		authorization := context.GetHeader("Authorization")
		if authorization != "" {
			storeToken, exists := utils.Cache.Get("token:" + authorization)
			if exists {
				jwtInstance := utils.NewJWT()
				claims, err2 := jwtInstance.ParseToken(authorization)
				if err2 == nil {
					if claims.ID == storeToken {
						log.Println("user logged")

						context.Set("claims", claims)
					}
				}
			}
		}
		context.Next()
	}
}
