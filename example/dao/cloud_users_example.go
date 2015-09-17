package main

import (
	"log"

	"github.com/lucasjo/porgex-go/dao"
)

func main() {

	users := dao.GetAllCloudUsers()

	log.Printf("%v\n", users[0].Login)

	login := dao.FindLoiginCloudUsers(users[0].Login)

	log.Printf("%v\n", login)
	//55de96460f510687e7000001
	loginId := dao.FindIdCloudUsers("55de96460f510687e7000001")
	log.Printf("%v\n", loginId)

}
