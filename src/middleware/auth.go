package middleware

import (
	"github.com/afiefafian/go-pos/src/config"
	"github.com/afiefafian/go-pos/src/exception"
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

// Protected protect routes
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		AuthScheme:   "JWT",
		SigningKey:   []byte(config.SECRET),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	return exception.Unauthorized("")
}
