package main

import (
	"fmt"

	"github.com/lucasjo/porgex-go/db/mongo"
)

//	"flag"

func main() {
	mongoConn := mongo.NewMongoDBConn()
	mongoConn.Connect()
	mongoConn.DisConnect()
	fmt.Printf("%+v\n", mongoConn)

}
