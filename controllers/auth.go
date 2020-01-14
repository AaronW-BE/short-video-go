package controllers

import (
	"MiniVideo/models"
	"MiniVideo/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"time"
)

func CommonLogin(context *gin.Context) {
	name := context.PostForm("username")
	pwd := context.PostForm("password")

	encryptedPwd, err := utils.EncryptStringToPassword(pwd)
	if err != nil {
		utils.Response500(context, "")
		return
	}
	user := &models.User{}
	loginUser, err := user.FindOne(bson.M{
		"username": name,
		"password": encryptedPwd,
	})
	if err != nil {
		utils.Response400(context, "账号或密码错误")
		return
	}

	if loginUser.State == models.STATE_NORMAL {
		// jwt生成token

		utils.Response200(context, nil, "登录成功")
		return
	}
	if loginUser.State == models.STATE_BLOCKED {
		utils.Response401(context, "该用户已禁用，无法登录")
		return
	}
	utils.Response401(context, "登录失败，原因未知")
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
