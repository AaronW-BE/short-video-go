package router

import (
	"MiniVideo/controllers"
	"MiniVideo/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRoute() {

	router := gin.Default()

	apiGroup := router.Group("/api", middlewares.Api())
	webGroup := router.Group("/")
	controllers.InitApi(apiGroup)
	controllers.InitWebGroup(webGroup)

	_ = router.Run(":8000")
}
