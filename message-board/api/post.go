package api

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/model"
	"message-board/service"
	"message-board/tool"
	"strconv"
	"time"
)

func addPost(ctx *gin.Context) {
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)
	txt := ctx.PostForm("txt")
	post := model.Post{
		Txt:        txt,
		UserName:   username,
		PostTime:   time.Now(),
		UpdateTime: time.Now(),
	}

	err := service.AddPost(post)
	if err != nil {
		fmt.Println("insert post error", err)
		tool.RespInterError(ctx)
		return

	}
}
func post_id(ctx *gin.Context){
	iUsername, _ := ctx.Get("username")
	username := iUsername.(string)
	txt := ctx.PostForm("txt")
	post := model.Post{
		Txt:        txt,
		UserName:   username,
		PostTime:   time.Now(),
		UpdateTime: time.Now(),
	}
	err := service.ChangePost(post)
	if err != nil {
		fmt.Println("change post error", err)
		tool.RespInterError(ctx)
		return

	}
}
func briefPosts(ctx *gin.Context){
	posts,err:=service.GetPost()
	if err!=nil{
		fmt.Println("get posts err:",err)
		tool.RespInterError(ctx)
		return
	}
	tool.RespSuccesfulWithDate(ctx,posts)
}
func postDetail(ctx *gin.Context){
	postIdString:=ctx.Param("post_id")
	postId,err:=strconv.Atoi(postIdString)
	if err!=nil{
		fmt.Println("post id string to post id int error",err)
		tool.RespErrorWithDate(ctx,"post id string to post id int error")
	}
posts,err:=service.GetPostById(postId)
if err!=nil{
	fmt.Println("get post id by id err",err)
	tool.RespInterError(ctx)
	return
}
comments,err:=service.GetPostComments(postId)
if err!=nil{
	if err!=sql.ErrNoRows{
		fmt.Println("get post comments err:",err)
		tool.RespInterError(ctx)
		return
	}
}
var postDetail model.PostDetail
postDetail.Post=posts
postDetail.Comment=comments
fmt.Println("123")
tool.RespErrorWithDate(ctx,postDetail)
}