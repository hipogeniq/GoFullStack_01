package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hipogeniq/GoFullStack_01/api/handlers"
)

func setupRoutes(app *fiber.App) {

	app.Get("/", handlers.Home)

}
