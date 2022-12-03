package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/lilendian0x00/go-mvc/configs"
	"github.com/lilendian0x00/go-mvc/controllers"
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

func setupRoutes(app *fiber.App) {
	/* \/api\/ */
	api := app.Group("/api")
	v1 := api.Group("/v1", func(ctx *fiber.Ctx) error {
		ctx.Accepts("application/json")
		return ctx.Next()
	})

	v1.Get("/list", controllers.TodosIndex)

	v1.Post("/add", controllers.AddTodo)

	v1.Delete("/remove", controllers.RemoveTodo)
}

func main() {
	app = configs.FiberNewConfig()

	app.Use(cors.New())
	// define the dir for public
	serveStatic(app)

	// setup api routes
	setupRoutes(app)

	log.Printf("[REST] listening at %v\n", Configs.ServerAddress)
	err := app.Listen(Configs.ServerAddress)
	if err != nil {
		log.Fatalln(err)
	}
}

