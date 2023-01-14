package dto

import "github.com/gofiber/fiber/v2"

type ErrorResponse struct {
	Method string
	Path   string
	Msg    error
}

func NewErrorResponse(ctx *fiber.Ctx, msg error) ErrorResponse {
	return ErrorResponse{
		Method: ctx.Method(),
		Path:   ctx.Path(),
		Msg:    msg,
	}
}
