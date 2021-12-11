package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/service"
	"message-board/tool"
)

func changePassword(ctx *gin.Context) {
	oldPassword := ctx.PostForm("old_password")
	newPassword := ctx.PostForm("new_password")
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)
	flag,err:=service.IsPasswordCorrect(username,oldPassword)
	if err!=nil{
		fmt.Println("判断正确密码错误：",err)
		tool.RespInterError(ctx)
		return

	}
	if !flag{
		tool.RespErrorWithDate(ctx,"旧密码错误")
		return
	}
	err=service.ChangePassword(username,newPassword)
	if err!=nil{
		fmt.Println("修改密码失败",err)
		tool.RespInterError(ctx)
		return
	}
	tool.RespSuccesful(ctx)
}
