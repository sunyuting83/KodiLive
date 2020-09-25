package router

import (
	"github.com/gofiber/fiber/v2"
)

// InitRouter create router
func InitRouter() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	return app
}
