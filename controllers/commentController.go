package controllers

import (
	"MiniVideo/services"
	"MiniVideo/utils"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
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

func VideoLike(ctx *gin.Context) {
	videoId := ctx.Param("id")
	claims, exists := ctx.Get("claims")
	if !exists {
		utils.Response400(ctx, "未登录")
		return
	}
	claimsObj := claims.(*utils.CustomClaims)

	err := services.LikeVideo(claimsObj.ID, bson.ObjectIdHex(videoId))
	if err != nil {
		utils.Response500(ctx, "点赞失败")
		return
	}
	utils.Response200(ctx, nil, "操作成功")
	return
}
