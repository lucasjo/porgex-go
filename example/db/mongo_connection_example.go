package main

import (
	"fmt"

	"github.com/lucasjo/porgex-go/db"
)

func main() {

	conn := db.NewConnction()

	fmt.Printf("conn %v\n", conn)
	fmt.Printf("conn %v\n", conn.Info().Database)

}
