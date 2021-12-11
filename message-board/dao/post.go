package dao

import "message-board/model"

func InsetPost(post model.Post)error{
	_,err:=dB.Exec("INSERT INTO post(txt,username,post_time,update_time)"+"value(?,?,?,?;",post.Txt,post.UserName,post.PostTime,post.UpdateTime)

	return err
}
func ChangePost(post model.Post)error{
	_,err:=dB.Exec("DELETE FROM post(txt)"+"values(?;",post.Txt)
	return err
	_,err1:=dB.Exec("INSERT INTO post(txt) "+"value(?;", post.Txt)
	return err1
}
func SelectPosts()([]model.Post,error) {
	var posts []model.Post
	row, err := dB.Query("SELECT id,username,txt,post_time,update_time,comment_num FROM post")
	if err != nil {
		return nil, err
	}
	defer row.Close()
	for row.Next() {
		var post model.Post
		err = row.Scan(&post.Id, &post.UserName, &post.Txt, &post.PostTime, &post.UpdateTime, &post.CommentNum)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)


	}
	return posts, nil
}
func SelectPostByPostId(postId int)(model.Post,error){
	var post model.Post
	row:=dB.QueryRow("SELECT id,username,txt,post_time,update_time,comment_num FROM post WHERE id =?")
	if row.Err()!=nil{
		return post,row.Err()
	}
	err:=row.Scan(&post.Id, &post.UserName, &post.Txt, &post.PostTime, &post.UpdateTime, &post.CommentNum)
	if err!=nil{
		return post,err
	}
	return post,nil
}