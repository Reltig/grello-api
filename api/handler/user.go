package handler

import (
	"grello-api/api/request"
	"grello-api/api/response"
	"grello-api/config"
	"grello-api/database"
	"grello-api/internal/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

var (
	BCRYPT_COST = config.ConfigInt("BCRYPT_COST")
)
//TODO: move to utils
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), BCRYPT_COST)
	return string(bytes), err
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var user model.User
	db.Find(&user, id)
	if user.ID == 0 {
		return response.NotFound(c, "No user found with ID", nil)
	}
	return response.Ok(c, "User found", response.User{}.FromModel(&user))
}

func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	request := new(request.CreateUser)
	if err := c.BodyParser(request); err != nil {
		return response.BadRequest(c, "Review your input", err.Error())
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return response.BadRequest(c, "Invalid request body", err.Error())
	}

	hash, err := HashPassword(request.Password)
	if err != nil {
		return response.BadRequest(c, "Couldn't hash password", err.Error())
	}

	user := model.User{
		Username:   request.Username,
		Email:      request.Email,
		Password:   hash,
		FirstName:  request.FirstName,
		SecondName: request.SecondName,
	}
	if err := db.Create(&user).Error; err != nil {
		return response.BadRequest(c, "Couldn't create user", err.Error())
	}

	return response.Ok(c, "Created user", response.User{}.FromModel(&user))
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	request := new(request.UpdateUser)
	if err := c.BodyParser(request); err != nil {
		return response.BadRequest(c, "Review your input", err.Error())
	}
	var user model.User
	db.Find(&user, id)
	if user.ID == 0 {
		return response.NotFound(c, "No user found with ID", nil)
	}
	if request.Password != nil {
		hash, err := HashPassword(*request.Password)
		if err != nil {
			return response.BadRequest(c, "Couldn't hash password", err.Error())
		}
		request.Password = &hash
	}
	if err := db.Model(&user).Updates(request).Error; err != nil {
		return response.InternalServerError(c, "Database error: couldn't update user", nil)
	}
	return response.Ok(c, "User updated", response.User{}.FromModel(&user))
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var user model.User
	db.Find(&user, id)
	if user.ID == 0 {
		return response.NotFound(c, "No user found with ID", nil)
	}
	if err := db.Delete(&user).Error; err != nil {
		return response.InternalServerError(c, "Database error: couldn't delete user", nil)
	}
	return response.Ok(c, "User deleted", response.User{}.FromModel(&user))
}

func GetUserWorkspaces(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DB
	var user model.User
	if err := db.Model(&model.User{}).Preload("Workspaces").Find(&user, id).Error; err != nil {
		return response.InternalServerError(c, "Database error: couldn't get user workspaces", nil)
	}
	return response.Ok(c, "User workspaces", response.Workspace{}.FromModelCollection(user.OwnedWorkspaces))
}
