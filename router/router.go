package router

import (
	"KodiLive/controller"
	"fmt"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/logger"
)

// InitRouter create router
func InitRouter(proxy bool) *fiber.App {
	app := fiber.New()

	// app.Use(logger.New())

	// // Or extend your config for customization
	// app.Use(logger.New(logger.Config{
	// 	Format:     "${pid} ${status} - ${method} ${path}\n",
	// 	TimeFormat: "02-Jan-2006",
	// 	TimeZone:   "America/New_York",
	// 	Output:     os.Stdout,
	// }))

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
		return c.SendString(controller.GetList(c.Params("*"), c.BaseURL(), false, proxy))
	})

	app.Get("/tags/*", func(c *fiber.Ctx) error {
		return c.SendString(controller.GetList(c.Params("*"), c.BaseURL(), true, proxy))
	})

	app.Get("/livecam/:name/playlist.m3u8", func(c *fiber.Ctx) error {
		return c.SendString(controller.GetM3u8(c.Params("name"), proxy))
	})

	return app
}
