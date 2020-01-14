package utils

import (
	"github.com/gin-gonic/gin"
	"log"
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
	Response(ctx, 200, nil, "request not found")
}

func Response200(ctx *gin.Context, data interface{}, msg string) {
	Response(ctx, 200, data, msg)
}

func Response401(ctx *gin.Context, msg string) {
	Response(ctx, 401, nil, msg)
}

func Response400(ctx *gin.Context, msg string) {
	if msg == "" {
		msg = "bad request"
		log.Panicln("response an error...")
	}
	Response(ctx, 400, nil, msg)
}

func Response500(ctx *gin.Context, msg string) {
	if msg == "" {
		msg = "Internal Server error"
		log.Panicln("response an error...")
	}
	Response(ctx, 500, nil, msg)
}
