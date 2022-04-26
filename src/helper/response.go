package helper

import (
	"github.com/afiefafian/go-pos/src/model"
	"github.com/gofiber/fiber/v2"
)

func JsonOK(ctx *fiber.Ctx, data interface{}, message string) error {
	if message == "" {
		message = MsgSuccess
	}
	return ctx.JSON(model.BaseResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}
