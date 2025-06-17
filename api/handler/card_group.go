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

func GetCardGroup(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var cardGroup model.CardGroup
	db.Find(&cardGroup, id)
	if cardGroup.ID == 0 {
		return response.NotFound(c, "No card group found with ID", nil)
	}
	return response.Ok(c, "Card group found", response.CardGroup{}.FromModel(&cardGroup))
}

func CreateCardGroup(c *fiber.Ctx) error {
	db := database.DB
	request := new(request.CreateCardGroup)
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
	var board model.Board
	db.Find(&board, request.BoardID).Preload("Users")
	if board.ID == 0 {
		return response.NotFound(c, "No board found with ID", nil)
	}
	if collections.Any(board.Users, func(user *model.User) bool { return auth.UserID == user.ID }) {
		return response.Unauthorized(c, "You can't create card group in this board", nil)
	}

	cardGroup := model.CardGroup{
		Name:    request.Name,
		BoardID: request.BoardID,
	}
	if err := db.Create(&cardGroup).Error; err != nil {
		return response.BadRequest(c, "Couldn't create card group", err.Error())
	}

	return response.Ok(c, "Created card group", response.CardGroup{}.FromModel(&cardGroup))
}

func UpdateCardGroup(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	request := new(request.UpdateBoard)
	if err := c.BodyParser(request); err != nil {
		return response.BadRequest(c, "Review your input", err.Error())
	}
	var cardGroup model.CardGroup
	db.Find(&cardGroup, id)
	if cardGroup.ID == 0 {
		return response.NotFound(c, "No card group found with ID", nil)
	}

	auth, err := utils.Auth(c)
	if err != nil {
		return err
	}
	var board model.Board
	db.Find(&board, cardGroup.BoardID).Preload("Users")
	if board.ID == 0 {
		return response.NotFound(c, "No board found with ID", nil)
	}
	if collections.Any(board.Users, func(user *model.User) bool { return auth.UserID == user.ID }) {
		return response.Unauthorized(c, "You can't update card group in this board", nil)
	}

	if err := db.Model(&cardGroup).Updates(request).Error; err != nil {
		return response.InternalServerError(c, "Database error: couldn't update card group", nil)
	}
	return response.Ok(c, "Card group updated", response.CardGroup{}.FromModel(&cardGroup))
}

func DeleteCardGroup(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var cardGroup model.CardGroup
	db.Find(&cardGroup, id)
	if cardGroup.ID == 0 {
		return response.NotFound(c, "No card group found with ID", nil)
	}

	auth, err := utils.Auth(c)
	if err != nil {
		return err
	}
	var board model.Board
	db.Find(&board, cardGroup.BoardID).Preload("Users")
	if board.ID == 0 {
		return response.NotFound(c, "No board found with ID", nil)
	}
	if collections.Any(board.Users, func(user *model.User) bool { return auth.UserID == user.ID }) {
		return response.Unauthorized(c, "You can't delete card group in this board", nil)
	}

	if err := db.Delete(&cardGroup).Error; err != nil {
		return response.InternalServerError(c, "Database error: couldn't delete card group", nil)
	}
	return response.Ok(c, "Card group deleted", response.CardGroup{}.FromModel(&cardGroup))
}
