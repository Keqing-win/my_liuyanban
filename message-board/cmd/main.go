package main
import (
	"message-board/api"
	"message-board/dao"
)

func main(){
	api.InitEngine()
	dao.InitDB()
}
