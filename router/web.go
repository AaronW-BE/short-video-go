package router

import (
	"MiniVideo/models"
	"MiniVideo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitWebGroup(router *gin.RouterGroup) {
	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"code": 0,
		})
	})

	debugGroup := router.Group("/devtool")
	{
		debugGroup.GET("/", func(context *gin.Context) {
			context.Redirect(http.StatusFound, "./main")
		})
		debugGroup.GET("/main", func(context *gin.Context) {
			context.HTML(http.StatusOK, "devtool/main.tpl", gin.H{
				"title": "开发控制台",
			})
		})
		debugGroup.GET("/cache", func(context *gin.Context) {
			caches := utils.Cache.Items()
			context.HTML(http.StatusOK, "devtool/cache.tpl", gin.H{
				"title":  "开发控制台/Cache",
				"caches": caches,
			})
		})

		debugGroup.GET("/user", func(context *gin.Context) {
			model := &models.User{}
			users := model.FindAll()
			context.HTML(http.StatusOK, "devtool/user.tpl", gin.H{
				"title": "开发控制台/User",
				"users": users,
			})
		})
		debugGroup.GET("/log")
	}

}
