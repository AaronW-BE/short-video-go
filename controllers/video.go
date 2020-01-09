package controllers

import (
	"MiniVideo/services"
	"github.com/gin-gonic/gin"
)

func Video(ctx *gin.Context)  {
	var videoService services.Video
	err := ctx.ShouldBindJSON(&videoService)
	if err != nil {
		ctx.JSON(500, gin.H{
			"code": "parameter failed",
		})
		return
	}
	id, err := videoService.Insert()
	if err != nil {
		ctx.JSON(500, gin.H{
			"code": "insert failed",
		})
		return
	}
	ctx.JSON(200, gin.H{
		"code": 0,
		"data": id,
	})
}