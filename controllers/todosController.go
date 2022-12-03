package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lilendian0x00/go-mvc/initializers"
	"github.com/lilendian0x00/go-mvc/models"
)

type resp struct {
	Status			string 		`json:"status"`
	Msg				string		`json:"msg"`
}

func TodosIndex(c *fiber.Ctx) error {
	result := initializers.DB.Find(&models.Todo{})
	return c.Status(200).JSON(result)
}

func AddTodo(c *fiber.Ctx) error {
	todo :=  new(models.Todo)
	if err := c.BodyParser(&todo); err != nil {
		c.Status(400).JSON(resp{
			Status: "failure",
			Msg:    "bad post body!",
		})
		return err
	}

	initializers.DB.Create(&todo)

	return c.Status(200).JSON(resp{
		Status: "success",
		Msg:    "todo added successfully!",
	})
}

func RemoveTodo(c *fiber.Ctx) error {
	todo :=  new(models.Todo)
	if err := c.BodyParser(&todo); err != nil {
		c.Status(400).JSON(resp{
			Status: "failure",
			Msg:    "bad post body!",
		})
		return err
	}
	initializers.DB.Delete(&todo)

	return c.Status(200).JSON(resp{
		Status: "success",
		Msg:    "todo deleted successfully!",
	})
}


