package controllers

import "github.com/gin-gonic/gin"

func InitWebGroup(router *gin.RouterGroup)  {
	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"code": 0,
		})
	})
}