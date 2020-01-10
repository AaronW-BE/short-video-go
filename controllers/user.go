package controllers

import (
	"MiniVideo/utils"
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
