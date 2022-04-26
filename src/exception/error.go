package exception

import (
	"github.com/afiefafian/go-pos/src/helper"
	"github.com/gofiber/fiber/v2"
)

var EntityNotFound = func(entity string) error {
	message := helper.MsgErrEntityNotFound(entity)
	err := fiber.ErrNotFound
	err.Message = message
	return err
}

var Unauthorized = func(message string) error {
	err := fiber.ErrUnauthorized
	if message != "" {
		err.Message = message
	}
	return err
}
