package controllers

import (
	"MiniVideo/services"
	"github.com/gin-gonic/gin"
	"log"
)

func PublishVideo(ctx *gin.Context) {
	var videoService services.Video
	err := ctx.ShouldBind(&videoService)
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
		log.Fatal(err)

		return
	}
	ctx.JSON(200, gin.H{
		"code": 0,
		"data": id,
	})
}

func VideoList(ctx *gin.Context) {
	videos, err := services.List()
	if err != nil {
		ctx.JSON(500, gin.H{
			"code": "obtain video list failed",
		})
		log.Println(err)
		return
	}
	ctx.JSON(200, gin.H{
		"code": 0,
		"data": videos,
	})
}

func FollowingVideo(ctx *gin.Context) {

}
