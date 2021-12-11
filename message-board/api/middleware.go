package api

import (
	"github.com/gin-gonic/gin"
	"message-board/tool"
)

func auth(ctx *gin.Context){
	username,err:=ctx.Cookie("username")
	if err!=nil{
		tool.RespErrorWithDate(ctx,"请登录后进行操作")
		return
	}
	ctx.Set("username",username)
	ctx.Next()
}
