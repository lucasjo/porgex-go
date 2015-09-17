package main

import (
	"log"

	"github.com/lucasjo/porgex-go/dao"
)

func main() {

	apps := dao.FindAllApps()

	log.Printf("%v\n", apps)

	app := dao.FindIdApp("55ee3a460f5106ab680000ca")

	log.Printf("%v\n", app)

}
