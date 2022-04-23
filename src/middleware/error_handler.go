package middleware

import (
	"github.com/afiefafian/go-pos/src/helper"
	"github.com/afiefafian/go-pos/src/model"
	"github.com/gofiber/fiber/v2"
)

// NotFound returns custom 404 error
func RouteNotFound(ctx *fiber.Ctx) error {
	response := model.BaseResponse{
		Success: false,
		Message: helper.MsgErrUrlNotFound(ctx),
	}
	return ctx.Status(fiber.StatusNotFound).JSON(response)
}

// FiberErrorHandler customer error handler for fiber
func FiberErrorHandler(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusInternalServerError).JSON(model.BaseResponse{
		Success: false,
		Message: "INTERNAL SERVER ERROR",
		Error:   err.Error(),
	})
}
