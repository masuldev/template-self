package api

import (
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/masuldev/template-self/internal/api/group"
	"github.com/masuldev/template-self/pkg/log"
)

func Route() *fiber.App {
	app := fiber.New()
	app.Use(recover.New())
	app.Use(log.NewLogger())

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusOK)
	})

	group.AuthRoute(app)

	return app
}
