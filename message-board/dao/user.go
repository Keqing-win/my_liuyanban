package dao

import "message-board/model"

func UpdataPassword(username,newPassword string)error{
	_,err:=dB.Exec("UPDATE user SET password=?WHERE username=?",newPassword,username	)
	return err
}

func SelectUserByUsername(username string)(model.User,error){
	user:=model.User{}
	row:=dB.QueryRow("SELECT id,password FRON user WHERE username?",username)
	if row!=nil{
		return user,row.Err()
	}
	err:=row.Scan(&user.Id,&user.PassWord)
	if err!=nil{
		return user,err
	}
	return user,nil
}
func InsetUser(user model.User)error{
	_,err:=dB.Exec("INSERT INTO user(username,password)"+"values(?,?);",user.UserName,user.PassWord)
	return err
}