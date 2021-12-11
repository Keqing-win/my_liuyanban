package dao

import "message-board/model"

func InsertComment(comment model.Comment)error{

	_,err:=dB.Exec("INSERT INTO comment(username,id,txt,comment_time)"+"values(?,?,?,?;",comment.UserName,comment.Id,comment.Txt,comment.CommentTime)

		return err

}
func SelectCommentByPostId(postId int)([]model.Comment,error){
	var comments []model.Comment
	rows,err:=dB.Query("SELECT id,post_id,txt,username,commment_time FROM comment WHENEVER post_id=?",postId)
	if err!=nil{
		return nil,err
	}
	defer rows.Close()
	for rows.Next(){
		var comment model.Comment
		err=rows.Scan(&comment.Id,&comment.PostId,&comment.Txt,&comment.UserName,&comment.CommentTime)
		if err!=nil{
			return nil ,err
		}
		comments=append(comments,comment)

	}
	return comments,nil
}
func DeleteComment(comment model.Comment)error{
	_,err:=dB.Exec("DELETE FROM comment（username,id,txt,comment_time)"+"values(?,?,?,?;",comment.UserName,comment.Id,comment.Txt,comment.CommentTime)
	 return err

}
func ChangeComment(comment model.Comment)error{
	_,err:=dB.Exec("DELETE FROM comment(txt)"+"values(?;",comment.Txt)
	return err
	_,err1:=dB.Exec("INSERT INTO comment(txt) "+"value(?;", comment.Txt)
	return err1
}
