package http

import (
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/rahmadfani/wilayah/internal/config"
)

func InitCors(app *fiber.App) {
	allowOrigins := []string{}
	allowOriginsStr := config.ALLOW_ORIGINS

	if allowOriginsStr != "" {
		allowOriginsStr = strings.Replace(allowOriginsStr, " ", "", -1)
		allowOrigins = strings.Split(allowOriginsStr, ",")
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: allowOrigins,
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
		AllowMethods: []string{"OPTIONS", "GET"},
	}))
}
