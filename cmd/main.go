package main

import (
	"fmt"
	"log"

	"grello-api/api/router"
	"grello-api/config"
	"grello-api/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "App Name",
	})
	// app.Use(cors.New())

	database.ConnectDb()

	router.SetupRoutes(app)

	printRoutes(app)

	port := fmt.Sprintf(":%s", config.Config("APP_PORT"))
	log.Fatal(app.Listen(port))
}

func printRoutes(app *fiber.App) {
	fmt.Println("╔════════════════════════════════════════════╗")
	fmt.Println("║                 ROUTES LIST                 ║")
	fmt.Println("╠══════════╦══════════════════════════════════╣")
	fmt.Println("║ METHOD   ║ PATH                             ║")
	fmt.Println("╠══════════╬══════════════════════════════════╣")
	
	for _, route := range app.GetRoutes() {
		fmt.Printf("║ %-8s ║ %-32s ║\n", route.Method, route.Path)
	}
	
	fmt.Println("╚══════════╩══════════════════════════════════╝")
}