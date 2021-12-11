package service

import (
	"database/sql"
	"message-board/dao"
	"message-board/model"
)

func ChangePassword(username,newPassword string)error{
	err:=dao.UpdataPassword(username,newPassword)
	return err

}
func IsPasswordCorrect(username,password string)(bool,error){
      user,err:=dao.SelectUserByUsername(username)
	  if err!=nil{
		  if err==sql.ErrNoRows{
			  return false,nil
		  }
		  return false,err
	  }
	  if user.PassWord!=password{
		  return false,err
	  }
	  return true,nil

}
func IsRepeatUserName(username string)(bool,error){
	_,err:=dao.SelectUserByUsername(username)
	if err!=nil{
		if err==sql.ErrNoRows{
			return false,nil
		}
		return false,err

	}
	return true,nil

}

func Regist(user model.User)error{
	err:=dao.InsetUser(user)
	return err
}