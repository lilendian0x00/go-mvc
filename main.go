package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lilendian0x00/go-mvc/configs"
	"github.com/lilendian0x00/go-mvc/initializers"
	"github.com/lilendian0x00/go-mvc/utils"
	"log"
)

var Configs		utils.Config
var App 		*fiber.App

func init() {
	var err error
	Configs, err = utils.LoadEnvConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	initializers.ConnectToDatabase(Configs.DBSource)
}

func main() {
	App = configs.FiberNewConfig()

	log.Printf("[REST] listening at %v\n", Configs.ServerAddress)
	err := App.Listen(Configs.ServerAddress)
	if err != nil {
		log.Fatalln(err)
	}
}

