package http

import "github.com/gofiber/fiber/v3"

func HttpSuccess(c fiber.Ctx, statusCode int, data any) error {
	return c.JSON(fiber.Map{
		"status": "OK",
		"data":   data,
	})
}
