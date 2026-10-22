package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/rahmadfani/wilayah/internal/container"

	"github.com/uptrace/bun"
)

func NewHttp(db *bun.DB) *fiber.App {
	ctn := container.NewHttpContainer(db)

	app := fiber.New()

	app.Use(logger.New())

	app.Get("/healthz", healthcheck.New(healthcheck.Config{
		ResponseFormat: healthcheck.FormatJSON,
	}))

	wilayah := app.Group("wilayah")
	ctn.WilayahHandler.PublicRouteWilayah(wilayah)

	return app
}
