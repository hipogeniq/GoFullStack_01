package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hipogeniq/GoFullStack_01/api/database"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Div Rhino!")
	})

	app.Listen(":3333")
}
