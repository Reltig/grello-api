package router

import (
	"grello-api/api/handler"
	"grello-api/api/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Auth
	auth := app.Group("/auth")
	auth.Post("/login", handler.Login)
	auth.Get("/user-data", middleware.Protected(), handler.UserData)

	// Api
	api := app.Group("/api", middleware.Protected())

	// User
	user := api.Group("/user")
	user.Get("/:id", handler.GetUser)
	user.Post("/", handler.CreateUser)
	user.Patch("/:id", handler.UpdateUser)
	user.Delete("/:id", handler.DeleteUser)
	user.Get("/:id/workspaces", handler.GetUserWorkspaces)

	// Workspace
	workspace := api.Group("/workspace")
	workspace.Get("/:id", handler.GetWorkspace)
	workspace.Post("/", handler.CreateWorkspace)
	workspace.Patch("/:id", handler.UpdateWorkspace)
	workspace.Delete("/:id", handler.DeleteWorkspace)
	
}