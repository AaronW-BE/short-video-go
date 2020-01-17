package controllers

import (
	"MiniVideo/services"
	"MiniVideo/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func GetVideoComments(ctx *gin.Context) {
	videoId := ctx.Param("id")

	utils.Response200(ctx, videoId, "ok")
}

func PublishComment(ctx *gin.Context) {
	videoId := ctx.Param("id")
	content := ctx.PostForm("content")

	log.Println("video id", videoId)

	if content == "" {
		utils.Response400(ctx, "请输入评论内容")
		return
	}
	claims, exists := ctx.Get("claims")
	if !exists {
		utils.Response400(ctx, "未登录")
		return
	}
	claimsObj := claims.(*utils.CustomClaims)

	err := services.PublishVideoComment(videoId, content, claimsObj.ID)
	if err != nil {
		log.Println(err)

		utils.Response400(ctx, "发布失败")
		return
	}
	utils.Response200(ctx, nil, "发布成功")
}
