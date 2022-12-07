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
	var Todos []models.Todo
	_ = initializers.DB.Find(&Todos)
	return c.Status(200).JSON(Todos)
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

func UpdateTodo(c *fiber.Ctx) error {
	todo :=  new(models.Todo)
	if err := c.BodyParser(&todo); err != nil {
		c.Status(400).JSON(resp{
			Status: "failure",
			Msg:    "bad post body!",
		})
		return err
	}

	initializers.DB.Model(&todo).Updates(todo)

	return c.Status(200).JSON(resp{
		Status: "success",
		Msg:    "todo successfully updated",
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



