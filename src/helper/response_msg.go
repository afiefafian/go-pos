package helper

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

const MsgSuccess = "Success"
const MsgErrValidation = "Validation Error"

func MsgErrEntityNotFound(entity string) string {
	return fmt.Sprintf("%s Not Found", entity)
}

func MsgErrUrlNotFound(ctx *fiber.Ctx) string {
	return fmt.Sprintf("Cannot %s %s", ctx.Method(), ctx.OriginalURL())
}
