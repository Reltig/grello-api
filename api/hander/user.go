package hander

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
	if (user.ID == 0 ) {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No product found with ID", "data": nil})
	}
	return c.JSON(fiber.Map{
		"status": "success", 
		"message": "Product found", 
		"data": response.UserResponse{}.FromModel(&user),
	})
}

func CreateUser(c *fiber.Ctx) error {
	db := database.DB
	request := new(request.CreateUserReqeust)
	if err := c.BodyParser(request); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "errors": err.Error()})
	}

	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request body", "errors": err.Error()})
	}

	hash, err := hashPassword(request.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "errors": err.Error()})
	}

	user := model.User{
		Username: request.Username,
		Email: request.Email,
		Password: hash,
		FirstName: request.FirstName,
		SecondName: request.SecondName,
	}
	if err := db.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "errors": err.Error()})
	}

	return c.JSON(fiber.Map{
		"status": "success", 
		"message": "Created user", 
		"data": response.UserResponse{}.FromModel(&user),
	})
}