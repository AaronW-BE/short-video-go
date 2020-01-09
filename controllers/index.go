package controllers

import "github.com/gin-gonic/gin"

func Index(context *gin.Context)  {
	context.JSON(200, gin.H{
		"status": 200,
	})
}