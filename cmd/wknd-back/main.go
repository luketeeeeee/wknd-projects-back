package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New(fiber.Config{
		AppName: "wknd-projects-back",
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hell, world!")
	})

	app.Listen(":3000")
}
