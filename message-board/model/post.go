package model
import "time"

type Post struct{
	Id int`json:"id"`
	CommentNum int`json:"comment_num"`
	Txt string `json:"txt"`
	UserName string `json:"user_name"`
	PostTime time.Time `json:"post_time"`
	UpdateTime time.Time `json:"update_time"`


}
type PostDetail struct {
	Post
	Comment []Comment
}
