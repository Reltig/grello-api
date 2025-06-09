package handler

import (
	"errors"
	"grello-api/api/request"
	"grello-api/api/response"
	"time"

	"grello-api/config"
	"grello-api/database"
	"grello-api/internal/model"
	"grello-api/internal/utils"

	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	SECRET = config.Config("SECRET")
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
		return response.Unauthorized(c, "Invalid username or password", nil)
	}
	if !CheckPasswordHash(input.Password, user.Password) {
		return response.Unauthorized(c, "Invalid username or password", nil)
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["user_id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(SECRET))
	if err != nil {
		return response.InternalServerError(c, "Error while signing token", err.Error())
	}

	return response.Ok(c, "Success login", t)
}

func UserData(c *fiber.Ctx) error {
	auth, err := utils.Auth(c)
	if err != nil {
		return err
	}
	return response.Ok(c, "Your user data", auth)
}
