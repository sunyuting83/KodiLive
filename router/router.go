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
		return c.SendString(controller.GetList(c.Params("*"), c.BaseURL(), false))
	})

	app.Get("/tags/*", func(c *fiber.Ctx) error {
		return c.SendString(controller.GetList(c.Params("*"), c.BaseURL(), true))
	})

	app.Get("/livecam/:name/playlist.m3u8", func(c *fiber.Ctx) error {
		return c.SendString(controller.GetM3u8(c.Params("name")))
	})

	return app
}
