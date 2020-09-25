package router

import (
	"KodiLive/controller"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

// InitRouter create router
func InitRouter() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendStatus(403)
	})

	app.Get("/index.m3u", func(c *fiber.Ctx) error {
		x, err := controller.GetIndex(c.BaseURL())
		if err != nil {
			fmt.Println(err)
		}
		return c.SendString(x)
	})

	app.Get("/channels/*", func(c *fiber.Ctx) error {

		return c.SendString(c.Params("*"))
	})

	app.Get("/tags/*", func(c *fiber.Ctx) error {

		return c.SendString(c.Params("*"))
	})

	return app
}
