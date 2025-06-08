package handler

import (
	"errors"
	"grello-api/api/request"
	"time"

	"grello-api/config"
	"grello-api/database"
	"grello-api/internal/model"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getUserByUsername(username string) (*model.User, error) {
	db := database.DB
	var user model.User
	if err := db.Where(&model.User{Username: username}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// Login
func Login(c *fiber.Ctx) error {
	input := new(request.LoginRequest)

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}

	user, err := getUserByUsername(input.Username)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Internal Server Error", "data": err})
	}
	if user == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid username or password", "data": err})
	}
	if !CheckPasswordHash(input.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid username or password", "data": nil})
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}

func UserData(c *fiber.Ctx) error {
	user := c.Locals("user")
	if user == nil {
		return c.Status(401).JSON(fiber.Map{"status": "error", "message": "No token", "data": nil}) 
	}
	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	return c.Status(200).JSON(fiber.Map{"status": "success", "message": "You trying get other user data", "data": claims}) 
}