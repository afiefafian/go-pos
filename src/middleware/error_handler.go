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
		Error:   &model.EmptyObject{},
	}
	return ctx.Status(fiber.StatusNotFound).JSON(response)
}
