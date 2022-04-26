package rest

import (
	"strconv"

	"github.com/afiefafian/go-pos/src/exception"
	"github.com/afiefafian/go-pos/src/helper"
	"github.com/afiefafian/go-pos/src/middleware"
	"github.com/afiefafian/go-pos/src/model"
	"github.com/afiefafian/go-pos/src/service"
	"github.com/gofiber/fiber/v2"
)

type PaymentController struct {
	paymentService *service.PaymentService
}

func NewPaymentController(paymentService *service.PaymentService) *PaymentController {
	return &PaymentController{paymentService}
}

func (c *PaymentController) Route(app *fiber.App) {
	route := app.Group("/payments")

	route.Get("/", middleware.Protected(), c.findAll)
	route.Get("/:id", middleware.Protected(), c.findByID)
	route.Post("/", c.create)
	route.Put("/:id", c.updateByID)
	route.Delete("/:id", c.deleteByID)
}

func (c *PaymentController) findAll(ctx *fiber.Ctx) error {
	query := helper.NewPaginationQueryFromCtx(ctx)
	payments, pagination, err := c.paymentService.FindAll(query)
	if err != nil {
		panic(err)
	}

	return helper.JsonOK(ctx, map[string]interface{}{
		"payments": payments,
		"meta":     pagination,
	}, "")
}

func (c *PaymentController) findByID(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		panic(exception.EntityNotFound("Payment"))
	}

	var response *model.GetPaymentResponse
	response, err = c.paymentService.GetByID(id)
	if err != nil {
		panic(err)
	}

	return helper.JsonOK(ctx, response, "")
}

func (c *PaymentController) create(ctx *fiber.Ctx) error {
	var request model.CreatePaymentRequest
	if err := ctx.BodyParser(&request); err != nil {
		panic(err)
	}

	response, err := c.paymentService.Create(request)
	if err != nil {
		panic(err)
	}

	return helper.JsonOK(ctx, response, "")
}

func (c *PaymentController) updateByID(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		panic(exception.EntityNotFound("Payment"))
	}

	var request model.UpdatePaymentRequest
	if err := ctx.BodyParser(&request); err != nil {
		panic(err)
	}

	request.ID = id

	err = c.paymentService.UpdateByID(request)
	if err != nil {
		panic(err)
	}

	return helper.JsonOK(ctx, nil, "")
}

func (c *PaymentController) deleteByID(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		panic(exception.EntityNotFound("Payment"))
	}

	err = c.paymentService.DeleteByID(id)
	if err != nil {
		panic(err)
	}

	return helper.JsonOK(ctx, nil, "")
}
