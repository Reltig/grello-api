package handler

import (
	"grello-api/api/request"
	"grello-api/api/response"
	"grello-api/database"
	"grello-api/internal/model"
	"grello-api/internal/utils"
	"grello-api/pkg/collections"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetWorkspace(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var workspace model.Workspace
	db.Find(&workspace, id)
	if workspace.ID == 0 {
		return response.NotFound(c, "No workspace found with ID", nil)
	}

	auth, err := utils.Auth(c)
	if err != nil {
		return err
	}
	if auth.UserID != workspace.OwnerID && collections.Any(workspace.Users, func(user *model.User) bool { return user.ID == auth.UserID }) {
		return response.Unauthorized(c, "You don't have permission to watch this workspace", nil)
	}

	return response.Ok(c, "Workspace found", response.Workspace{}.FromModel(&workspace))
}

func CreateWorkspace(c *fiber.Ctx) error {
	db := database.DB
	request := new(request.CreateWorkspace)
	if err := c.BodyParser(request); err != nil {
		return response.BadRequest(c, "Review your input", err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	auth, err := utils.Auth(c)
	if err != nil {
		return err
	}

	workspace := model.Workspace{
		Name:        request.Name,
		Description: request.Description,
		OwnerID:     auth.UserID,
	}
	if err := db.Create(&workspace).Error; err != nil {
		return response.BadRequest(c, "Couldn't create workspace", err.Error())
	}

	return response.Ok(c, "Created workspace", response.Workspace{}.FromModel(&workspace))
}

func UpdateWorkspace(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	request := new(request.UpdateWorkspace)
	if err := c.BodyParser(request); err != nil {
		return response.BadRequest(c, "Review your input", err.Error())
	}

	auth, err := utils.Auth(c)
	if err != nil {
		return err
	}

	var workspace model.Workspace
	db.Find(&workspace, id)
	if workspace.ID == 0 {
		return response.NotFound(c, "No workspace found with ID", nil)
	}
	if auth.UserID != workspace.OwnerID {
		return response.Unauthorized(c, "You don't have permission to update this workspace", nil)
	}

	if err := db.Model(&workspace).Updates(request).Error; err != nil {
		return response.InternalServerError(c, "Database error: couldn't update workspace", nil)
	}
	return response.Ok(c, "Workspace updated", response.Workspace{}.FromModel(&workspace))
}

func DeleteWorkspace(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var workspace model.Workspace
	db.Find(&workspace, id)
	if workspace.ID == 0 {
		return response.NotFound(c, "No workspace found with ID", nil)
	}
	auth, err := utils.Auth(c)
	if err != nil {
		return err
	}
	if auth.UserID != workspace.OwnerID {
		return response.Unauthorized(c, "You don't have permission to delete this workspace", nil)
	}
	if err := db.Delete(&workspace).Error; err != nil {
		return response.InternalServerError(c, "Database error: couldn't delete workspace", err.Error())
	}
	return response.Ok(c, "Workspace deleted", response.Workspace{}.FromModel(&workspace))
}
