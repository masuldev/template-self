package group

import (
	"github.com/masuldev/template-self/internal/controller/http"
)

func AuthRoute(app *fiber.App) {
	controller := http.NewAuthController()
	authGroup := app.Group("/api/v1")

	authGroup.Post("auth", controller.CreateCertificate)
}
