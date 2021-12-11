package model

import "time"

type Comment struct{
	Id int `json:"id"`
	PostId int `json:"post_id"`
	Txt string `json:"txt"`
	UserName string `json:"user_name"`
	CommentTime time.Time
}
