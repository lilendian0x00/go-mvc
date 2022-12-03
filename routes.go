package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lilendian0x00/go-mvc/controllers"
)

func SetupRoutes(app *fiber.App) {
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
