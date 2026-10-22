package http

import "github.com/gofiber/fiber/v3"

func HttpError(c fiber.Ctx, statusCode int, messageError string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"status":        "!OK",
		"message_error": messageError,
	})
}
