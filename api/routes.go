package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/your-moon/go-fiber-starter/models"
	"github.com/your-moon/go-fiber-starter/services"
)

var App *fiber.App

func Init() *fiber.App {
	App = fiber.New()

	App.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	App.Get("user/:id", func(c *fiber.Ctx) error {
		ID, _ := c.ParamsInt("id")

		var user models.User
		if err := services.DB.First(&user, ID).Error; err != nil {
			return c.Status(404).SendString("User not found")
		}

		return c.JSON(user)
	})

	return App
}
