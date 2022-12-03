package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/lilendian0x00/go-mvc/configs"
	"github.com/lilendian0x00/go-mvc/initializers"
	"github.com/lilendian0x00/go-mvc/utils"
	"log"
)

var Configs		utils.Config
var app 		*fiber.App

func init() {
	var err error
	// load config file
	Configs, err = utils.LoadEnvConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	initializers.ConnectToDatabase(Configs.DBSource)
	// migrate models
	initializers.SyncDB()
}

func serveStatic(app *fiber.App) {
	app.Static("/", "./views/build")
}

func main() {
	app = configs.FiberNewConfig()

	app.Use(cors.New())
	// define the dir for public
	serveStatic(app)

	// setup api routes
	SetupRoutes(app)

	log.Printf("[REST] listening at %v\n", Configs.ServerAddress)
	err := app.Listen(Configs.ServerAddress)
	if err != nil {
		log.Fatalln(err)
	}
}

