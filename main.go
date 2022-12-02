package main

import (
	"fmt"
	"github.com/lilendian0x00/go-mvc/initializers"
	"github.com/lilendian0x00/go-mvc/utils"
	"log"
)
func init() {
	config, err := utils.LoadEnvConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	initializers.ConnectToDatabase(config.DBSource)
}

func main() {
	fmt.Println("Hello World")
}

