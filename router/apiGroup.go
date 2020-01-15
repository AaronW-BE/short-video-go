package router

import (
	"MiniVideo/controllers"
	"MiniVideo/middlewares"
	"MiniVideo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitApi(api *gin.RouterGroup) {
	api.GET("/index", controllers.Index)

	tg := api.Group("/test")
	{
		tg.GET("/token", middlewares.AuthMiddleware(), func(context *gin.Context) {
			claims, _ := context.Get("claims")

			claimsObj := claims.(*utils.CustomClaims)
			context.String(http.StatusOK, "token test ok, login user id is %s", claimsObj.ID.String())
		})
	}

	vsg := api.Group("/videos")
	{
		vsg.GET("/recommend", controllers.VideoList)      // 推荐视频
		vsg.GET("/following", controllers.FollowingVideo) // 关注人视频
		vsg.POST("/publish", controllers.PublishVideo)    // 发布视频
	}

	vg := api.Group("/video", middlewares.AuthMiddleware())
	{

		vg.POST("/:id/hide", nil) // 隐藏视频

		vg.POST("/:id/show", nil) // 显示视频

		vg.GET("/:id/comments", nil) // 视频id评论

		vg.DELETE("/:id", nil) // 删除视频

		vg.POST("/:id/lock", nil) // 锁定视频，锁定后无法编辑，用户无法查看

		vg.DELETE("/:id/lock", nil) // 取消锁定

		//	=========================================================
		//	互动
		vg.POST("/:id/like", nil)   // 点赞
		vg.DELETE("/:id/like", nil) // 取消点赞
		//
		vg.POST("/:id/comment", nil)        // 评论视频
		vg.DELETE("/:id/comment/:cid", nil) // 删除评论
		//
		vg.POST("/:id/report", nil) // 举报
	}

	ug := api.Group("/user")
	ug.Use(middlewares.AuthMiddleware())
	{
		ug.POST(":id/follow", nil)   // 关注播主
		ug.DELETE(":id/follow", nil) //取消关注

		ug.GET(":id/followers", nil) // 粉丝
		ug.GET(":id/following", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"message": "following",
			})
		}) // 已关注播主
	}

	ag := api.Group("/auth")
	{
		ag.POST("/login/common", controllers.CommonLogin) // 通用登录
		ag.POST("/login/phone", nil)                      // 手机号登录

		ag.POST("/register/common", nil) // 通用注册
		ag.POST("/register/phone", nil)  // 通过手机号注册

		ag.GET("/captcha", controllers.PhoneCaptcha)
	}

}
