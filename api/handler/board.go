package handler

import (
	"grello-api/api/request"
	"grello-api/api/response"
	"grello-api/database"
	"grello-api/internal/model"
	"grello-api/internal/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetBoard(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var board model.Board
	db.Find(&board, id)
	if board.ID == 0 {
		return response.NotFound(c, "No board found with ID", nil)
	}
	var workspace model.Workspace
	db.Find(&workspace, board.WorkspaceID)
	if workspace.ID == 0 {
		return response.NotFound(c, "Board don't have workspace", nil)
	}

	auth, err := utils.Auth(c)
	if err != nil {
		return err
	}
	if auth.UserID != workspace.UserID {
		return response.Unauthorized(c, "You don't have permission to watch this board", nil)
	}

	return response.Ok(c, "Board found", response.Board{}.FromModel(&board))
}

func CreateBoard(c *fiber.Ctx) error {
	db := database.DB
	request := new(request.CreateBoard)
	if err := c.BodyParser(request); err != nil {
		return response.BadRequest(c, "Review your input", err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	var workspace model.Workspace
	db.Find(&workspace, request.WorkspaceID)
	if workspace.ID == 0 {
		return response.NotFound(c, "Trying link board with no existing workspace", nil)
	}

	auth, err := utils.Auth(c)
	if err != nil {
		return err
	}

	if auth.UserID != workspace.UserID {
		return response.Unauthorized(c, "You don't have permission to link board with this workspace", nil)
	}

	board := model.Board{
		Name:   	 request.Name,
		Description: request.Description,
		WorkspaceID: request.WorkspaceID,
	}
	if err := db.Create(&board).Error; err != nil {
		return response.BadRequest(c, "Couldn't create board", err.Error())
	}

	return response.Ok(c, "Created board", response.Board{}.FromModel(&board))
}

func UpdateBoard(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	request := new(request.UpdateBoard)
	if err := c.BodyParser(request); err != nil {
		return response.BadRequest(c, "Review your input", err.Error())
	}

	var board model.Board
	db.Find(&board, id)
	if board.ID == 0 {
		return response.NotFound(c, "No board found with ID", nil)
	}

	auth, err := utils.Auth(c)
	if err != nil {
		return err
	}

	var workspace model.Workspace
	db.Find(&workspace, board.WorkspaceID)
	if workspace.ID == 0 {
		return response.NotFound(c, "No workspace found with ID", nil)
	}
	if auth.UserID != workspace.UserID {
		return response.Unauthorized(c, "You don't have permission to update this board", nil)
	}

	if err := db.Model(&board).Updates(request).Error; err != nil {
		return response.InternalServerError(c, "Database error: couldn't update board", nil)
	}
	return response.Ok(c, "Board updated", response.Board{}.FromModel(&board))
}

func DeleteBoard(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var board model.Board
	db.Find(&board, id)
	if board.ID == 0 {
		return response.NotFound(c, "No board found with ID", nil)
	}
	var workspace model.Workspace
	db.Find(&workspace, board.WorkspaceID)
	if workspace.ID == 0 {
		return response.NotFound(c, "No workspace found with ID", nil)
	}
	auth, err := utils.Auth(c)
	if err != nil {
		return err
	}
	if auth.UserID != workspace.UserID {
		return response.Unauthorized(c, "You don't have permission to delete this board", nil)
	}
	if err := db.Delete(&board).Error; err != nil {
		return response.InternalServerError(c, "Database error: couldn't delete board", nil)
	}
	return response.Ok(c, "Board deleted", response.Board{}.FromModel(&board))
}