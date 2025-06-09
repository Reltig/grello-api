package utils

import (
	"grello-api/api/response"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type AuthData struct {
	UserID 	 uint
	Username string
}

func Auth(c *fiber.Ctx) (*AuthData, error) {
	user := c.Locals("user")
	if user == nil {
		return nil, response.Unauthorized(c, "No token", nil)
	}
	token := user.(*jwt.Token)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, response.Unauthorized(c, "No token", nil)
	}
	username := claims["username"].(string)
	userId := uint(claims["user_id"].(float64))
	return &AuthData{
		UserID: userId,
		Username: username,
	}, nil
}