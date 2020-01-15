package utils

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Response(ctx *gin.Context, code int, data interface{}, msg string) {
	if data != nil || data == "" {
		ctx.JSON(code, gin.H{
			"code": code,
			"msg":  msg,
			"data": data,
		})
		ctx.Done()
		return
	}

	ctx.JSON(code, gin.H{
		"code": code,
		"msg":  msg,
	})
	ctx.Done()
}

func Response404(ctx *gin.Context) {
	Response(ctx, http.StatusNotFound, nil, "request not found")
	ctx.Abort()
}

func Response200(ctx *gin.Context, data interface{}, msg string) {
	Response(ctx, http.StatusOK, data, msg)
}

func Response401(ctx *gin.Context, msg string) {
	Response(ctx, http.StatusUnauthorized, nil, msg)
	ctx.Abort()
}

func Response400(ctx *gin.Context, msg string) {
	if msg == "" {
		msg = "bad request"
		log.Panicln("response an error...")
	}
	Response(ctx, http.StatusBadRequest, nil, msg)
	ctx.Abort()
}

func Response500(ctx *gin.Context, msg string) {
	if msg == "" {
		msg = "Internal Server error"
		log.Panicln("response an error...")
	}
	Response(ctx, http.StatusInternalServerError, nil, msg)
	ctx.Abort()
}
