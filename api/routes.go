package api

import "github.com/gofiber/fiber/v2"

var App *fiber.App

func Init() *fiber.App {
	App = fiber.New()
	App.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	return App
}
