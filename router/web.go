package router

import (
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

	debugGroup := router.Group("/debug")
	{
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
		debugGroup.GET("/log")
	}

}
