package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/lucasjo/porgex-go/config"
)

func main() {
	root, _ := os.Getwd()

	configpath := filepath.Join(root, "config.gcfg")

	sConfig, err := config.NewConfig(configpath)

	if err != nil {
		fmt.Printf("Could not lod config file, error : %v\n", err)
	}

	fmt.Printf("config mongo section %v\n", sConfig.Mongo.Host)

}
