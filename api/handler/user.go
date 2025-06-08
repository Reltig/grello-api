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

func hashPassword(password string) (string, error) {
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

	hash, err := hashPassword(request.Password)
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
	if request.Username != nil {
		user.Username = *request.Username
	}
	if request.FirstName != nil {
		user.FirstName = request.FirstName
	}
	if request.SecondName != nil {
		user.SecondName = request.SecondName
	}
	if request.Email != nil {
		user.Email = *request.Email
	}
	if request.Password != nil {
		hash, err := hashPassword(*request.Password)
		if err != nil {
			return response.BadRequest(c, "Couldn't hash password", err.Error())
		}
		user.Password = hash
	}
	if err := db.Save(&user).Error; err != nil {
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
