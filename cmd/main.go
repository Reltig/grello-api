package main

import (
	"fmt"
	"log"

	"grello-api/api/router"
	"grello-api/config"
	"grello-api/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "App Name",
	})
	app.Use(cors.New())

	database.ConnectDb()

	router.SetupRoutes(app)

	port := fmt.Sprintf(":%s", config.Config("APP_PORT"))
	log.Fatal(app.Listen(port))
}