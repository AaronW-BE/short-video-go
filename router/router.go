package router

import (
	"MiniVideo/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRoute() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")
	router.StaticFile("/favicon.ico", "./resources/favicon.ico")

	apiGroup := router.Group("/api", middlewares.Api())
	webGroup := router.Group("/")
	InitApi(apiGroup)
	InitWebGroup(webGroup)

	_ = router.Run(":8000")
}
