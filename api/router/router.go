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

	// Board
	board := api.Group("/board")
	board.Get("/:id", handler.GetBoard)
	board.Post("/", handler.CreateBoard)
	board.Patch("/:id", handler.UpdateBoard)
	board.Delete("/:id", handler.DeleteBoard)
	
	// Card group
	cardGroup := api.Group("/card-group")
	cardGroup.Get("/:id", handler.GetCardGroup)
	cardGroup.Post("/", handler.CreateCardGroup)
	cardGroup.Patch("/:id", handler.UpdateCardGroup)
	cardGroup.Delete("/:id", handler.DeleteCardGroup)
}