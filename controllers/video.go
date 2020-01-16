package controllers

import (
	"MiniVideo/models"
	"MiniVideo/services"
	"MiniVideo/utils"
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

/**
推荐视频
已登录：每请求一次返回五条视频（后期加入推荐算法）
未登录：每请求一次返回十条视频
*/
func RecommendVideo(ctx *gin.Context) {
	claims, exists := ctx.Get("claims")
	isLogin := false
	if exists {
		isLogin = true
	}

	var videos []models.Video
	if isLogin {
		log.Println("已登录视频推荐")
		claimsObj := claims.(*utils.CustomClaims)

		videos = services.FetchRecommendVideos(claimsObj.ID)
	} else {
		videos, err := services.FetchHotVideos()
		if err != nil {
			log.Println(err)
			utils.Response500(ctx, err.Error())
			return
		}
		utils.Response200(ctx, videos, "请求成功")
		return
	}
	utils.Response200(ctx, videos, "请求视频成功")
}

func FollowingVideo(ctx *gin.Context) {
	// 获取所有关注人，视频按时间排序
}
