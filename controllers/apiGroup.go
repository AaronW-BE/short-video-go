package controllers

import "github.com/gin-gonic/gin"

func InitApi(api *gin.RouterGroup)  {
	api.GET("/index", Index)
	api.POST("/video", Video)
}
