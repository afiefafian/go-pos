package rest

import (
	"github.com/afiefafian/go-pos/src/helper"
	"github.com/afiefafian/go-pos/src/middleware"
	"github.com/afiefafian/go-pos/src/model"
	"github.com/afiefafian/go-pos/src/service"
	"github.com/gofiber/fiber/v2"
)

type OrderController struct {
	orderService *service.OrderService
}

func NewOrderController(orderService *service.OrderService) *OrderController {
	return &OrderController{orderService}
}

func (c *OrderController) Route(app *fiber.App) {
	app.Get("orders", middleware.Protected(), c.findAll)
	app.Post("orders/subtotal", middleware.Protected(), c.createSubTotal)
	app.Get("orders/:id", c.findByID)
	app.Post("orders", middleware.Protected(), c.create)
	app.Get("orders/:id/download", c.getOrderPdf)
	app.Get("orders/:id/check-download", middleware.Protected(), c.checkDLOrderPdf)
}

func (c *OrderController) findAll(ctx *fiber.Ctx) error {
	return nil
}

func (c *OrderController) findByID(ctx *fiber.Ctx) error {
	return nil
}

func (c *OrderController) create(ctx *fiber.Ctx) error {
	return nil
}

func (c *OrderController) createSubTotal(ctx *fiber.Ctx) error {
	var productsRequest []model.CreateOrderProductRequest
	if err := ctx.BodyParser(&productsRequest); err != nil {
		panic(err)
	}
	request := model.CreateSubTotalRequest{
		Products: productsRequest,
	}

	response, err := c.orderService.CheckSubTotal(request)
	if err != nil {
		panic(err)
	}

	return helper.JsonOK(ctx, response, "")
}

func (c *OrderController) getOrderPdf(ctx *fiber.Ctx) error {
	return nil
}

func (c *OrderController) checkDLOrderPdf(ctx *fiber.Ctx) error {
	response := model.CheckOrderDownloadResponse{
		IsDownload: true,
	}
	return helper.JsonOK(ctx, response, "")
}
