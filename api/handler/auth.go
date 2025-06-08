package handler

import (
	"errors"
	"grello-api/api/request"
	"grello-api/api/response"
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
	input := new(request.Login)

	if err := c.BodyParser(&input); err != nil {
		return response.BadRequest(c, "Error on login request", err.Error())
	}

	user, err := getUserByUsername(input.Username)

	if err != nil {
		return response.InternalServerError(c, "DB error", err.Error())
	}
	if user == nil {
		return response.Unauthorized(c, "Invalid username or password", err.Error())
	}
	if !CheckPasswordHash(input.Password, user.Password) {
		return response.Unauthorized(c, "Invalid username or password", nil)
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

	return response.Ok(c, "Success login", t)
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
