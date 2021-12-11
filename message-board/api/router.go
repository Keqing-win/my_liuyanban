package api

import "github.com/gin-gonic/gin"

func InitEngine() {
	engine:=gin.Default()
	engine.POST("/register")
	engine.POST("/login")
	userGroup:=engine.Group("/user")
	{
		userGroup.Use(auth)
		userGroup.POST("/password",changePassword)
	}
	postGroup:=engine.Group("/post")
	{
		postGroup.Use(auth)
		postGroup.POST("/",addPost)//发布新留言
		postGroup.POST("/",post_id)//修改新留言
		postGroup.GET("/",briefPosts)//查看全部留言概况
		postGroup.GET("/",postDetail)//查看一条留言详细信息和其下属评论
	}
	commentGroup:=engine.Group("/comment")
	{
		commentGroup.Use(auth)
		commentGroup.POST("/",addComment)//发布评论
		commentGroup.DELETE("/",changeComment)//修改评论
		commentGroup.DELETE("/",deleteComment)//删除评论

	}
	engine.Run()

}