package router

import (
	"grello-api/api/hander"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	// User
	user := api.Group("user")
	user.Get("/:id", hander.GetUser)
	user.Post("/", hander.CreateUser)
}