package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/masuldev/template-self/internal/controller/dto"
	"github.com/masuldev/template-self/pkg/log"
)

type Controller interface {
	CreateCertificate(c *fiber.Ctx) error
}

type authController struct {
}

func NewAuthController() Controller {
	return &authController{}
}

func (ac *authController) CreateCertificate(c *fiber.Ctx) error {
	var input *dto.CreateCertificate
	err := c.BodyParser(&input)
	if err != nil {
		log.Error("Example Err: Invalid Parameter")
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"example": "example",
	})
}
