package tool

import (
	"github.com/gin-gonic/gin"
	"message-board/model"
	"net/http"
)

func RespErrorWithDate(ctx *gin.Context,data interface{}){
	ctx.JSON(http.StatusOK,gin.H{
		"info":data,
	})
}
func RespInterError(ctx *gin.Context){
	ctx.JSON(http.StatusAlreadyReported,gin.H{
		"info": "服务器错误",
	})
}
func RespSuccesful(ctx *gin.Context){
	ctx.JSON(http.StatusOK,gin.H{
		"info":"成功",
	})
}
func RespSuccesfulWithDate(ctx *gin.Context, posts []model.Post) {
	ctx.JSON(http.StatusOK,gin.H{
		"info":"成功",
		"data":"data",
	})
}