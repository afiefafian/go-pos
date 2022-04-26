package middleware

import (
	"github.com/afiefafian/go-pos/src/helper"
	"github.com/afiefafian/go-pos/src/model"
	"github.com/go-playground/validator/v10"
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

// FiberErrorHandler custom error handler for fiber
func FiberErrorHandler(ctx *fiber.Ctx, err error) error {
	newErr := fiber.ErrInternalServerError

	// Retrieve the custom status code if it's an fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		newErr.Code = e.Code
		newErr.Message = e.Error()
	}

	if e, ok := err.(error); ok {
		newErr = fiber.ErrBadRequest
		newErr.Message = e.Error()
	}

	if e, ok := err.(validator.ValidationErrors); ok {
		errDetail := helper.FormatValidationError(e)

		return ctx.Status(fiber.StatusBadRequest).JSON(model.BaseResponse{
			Success: false,
			Message: helper.MsgErrValidation,
			Error:   errDetail,
		})
	}

	return ctx.Status(newErr.Code).JSON(model.BaseResponse{
		Success: false,
		Message: newErr.Error(),
		Error:   &model.EmptyObject{},
	})
}
