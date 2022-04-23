package config

import (
	"github.com/afiefafian/go-pos/src/middleware"
	"github.com/gofiber/fiber/v2"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: middleware.FiberErrorHandler,
	}
}
