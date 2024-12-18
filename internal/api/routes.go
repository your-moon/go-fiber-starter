package api

import (
	"github.com/gofiber/fiber/v2"

	"github.com/your-moon/go-fiber-starter/internal/models"
	"github.com/your-moon/go-fiber-starter/internal/services"
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

	App.Post("/user", func(c *fiber.Ctx) error {
		var params models.User
		if err := c.BodyParser(&params); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		if err := services.DB.Create(&params).Error; err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.JSON(params)
	})

	return App
}
