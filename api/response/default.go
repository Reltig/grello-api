package response

import "github.com/gofiber/fiber/v2"

const (
	SUCCESS_STATUS = "success"
	ERROR_STATUS   = "error"
)

func defaultSuccessResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  SUCCESS_STATUS,
		"message": message,
		"data":    data,
	})
}

func defaultErrorResponse(c *fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(fiber.Map{
		"status":  ERROR_STATUS,
		"message": message,
		"data":    data,
	})
}

func Ok(c *fiber.Ctx, message string, data interface{}) error {
	return defaultSuccessResponse(c, 200, message, data)
}

func Created(c *fiber.Ctx, message string, data interface{}) error {
	return defaultSuccessResponse(c, 201, message, data)
}

func BadRequest(c *fiber.Ctx, message string, data interface{}) error {
	return defaultErrorResponse(c, 400, message, data)
}

func Unauthorized(c *fiber.Ctx, message string, data interface{}) error {
	return defaultErrorResponse(c, 401, message, data)
}

func Forbidden(c *fiber.Ctx, message string, data interface{}) error {
	return defaultErrorResponse(c, 403, message, data)
}

func NotFound(c *fiber.Ctx, message string, data interface{}) error {
	return defaultErrorResponse(c, 404, message, data)
}

func InternalServerError(c *fiber.Ctx, message string, data interface{}) error {
	return defaultErrorResponse(c, 500, message, data)
}