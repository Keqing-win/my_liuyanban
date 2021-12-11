package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/model"
	"message-board/service"
	"message-board/tool"
	"strconv"
	"time"
)

func addComment(ctx *gin.Context){
iUsername,_:=ctx.Get("username")
username:=iUsername.(string)
txt:=ctx.PostForm("txt")
postIdString:=ctx.PostForm("post_id")
postId,err:=strconv.Atoi(postIdString)
if err!=nil{
	fmt.Println("post id to int error",err)
	tool.RespErrorWithDate(ctx,"文章的id有误")
	return
}
	comment:=model.Comment{
		PostId:postId,
		Txt: txt,
		UserName: username,
		CommentTime: time.Now(),

	}

err=service.AddComment(comment)
if err!=nil{
	fmt.Println("add comment err:",err)
	tool.RespInterError(ctx)
	return
}
tool.RespSuccesful(ctx)
}
func deleteComment(ctx *gin.Context){
	iUsername,_:=ctx.Get("username")
	username:=iUsername.(string)
	txt:=ctx.PostForm("txt")
	postIdString:=ctx.PostForm("post_id")
	postId,err:=strconv.Atoi(postIdString)
	if err!=nil{
		fmt.Println("post id to int error")
		tool.RespErrorWithDate(ctx,"文章的id错误")
		return
		}
	comment:=model.Comment{
		PostId:postId,
		Txt: txt,
		UserName: username,
		CommentTime: time.Now(),
	}
	err=service.DeleteComment(comment)
	if err!=nil{
		fmt.Println("delete common error",err)
		tool.RespInterError(ctx)
		return
	}
	}
	func changeComment(ctx *gin.Context){
		iUsername,_:=ctx.Get("username")
		username:=iUsername.(string)
		txt:=ctx.PostForm("txt")
		postIdString:=ctx.PostForm("post_id")
		postId,err:=strconv.Atoi(postIdString)
		if err!=nil{
			fmt.Println("post id to int error")
			tool.RespErrorWithDate(ctx,"文章的id错误")
			return
		}
		comment:=model.Comment{
			PostId:postId,
			Txt: txt,
			UserName: username,
			CommentTime: time.Now(),
		}
		err=service.Changecomment(comment)
		if err!=nil{
			fmt.Println("change comment default:")
			tool.RespErrorWithDate(ctx,"修改失败")

		}
	}




