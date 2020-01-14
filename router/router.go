package router

import (
	"MiniVideo/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRoute() {

	router := gin.Default()
	router.LoadHTMLGlob("templates/**/*")

	apiGroup := router.Group("/api", middlewares.Api())
	webGroup := router.Group("/")
	InitApi(apiGroup)
	InitWebGroup(webGroup)

	_ = router.Run(":8000")
}
