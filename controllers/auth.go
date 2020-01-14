package controllers

import (
	"MiniVideo/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CommonLogin(context *gin.Context) {
	name := context.PostForm("username")
	pwd := context.PostForm("password")
	if name == "aaronwb" && pwd == "123456" {

		//store token cache
		utils.Cache.Set("token", name+pwd, time.Hour)

		context.JSON(http.StatusOK, gin.H{
			"api_name": "common login",
			"data": gin.H{
				"token": name + pwd,
			},
		})
		return
	}
	context.JSON(http.StatusForbidden, gin.H{
		"api_name": "common login",
		"message":  "login failed, invalid username or password",
		"data":     gin.H{},
	})

}

/**
手机号登录
参数：手机号，验证码，校验码
*/
func PhoneLogin() {

}

/**
手机验证码
参数：手机号
返回：校验码
debug: 下发短信内容
*/
func PhoneCaptcha(ctx *gin.Context) {
	phone := ctx.DefaultQuery("phone", "")
	if len(phone) < 11 {
		utils.Response400(ctx, "手机号无效")
		return
	}

	captcha := utils.NewRandomNumber(6)

	utils.Cache.Set("captcha:"+phone, captcha, time.Minute*2)
	utils.Response200(ctx, gin.H{
		"debug":     true,
		"code":      captcha,
		"plain_sms": fmt.Sprintf("您的验证码为 %s ,有效期为：2分钟", captcha),
	}, "验证码发送成功")
}
