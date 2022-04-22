package rest

import "github.com/gofiber/fiber/v2"

type OrderController struct {
}

func NewOrderController() *OrderController {
	return &OrderController{}
}

func (c *OrderController) Route(app *fiber.App) {
	app.Get("orders", c.findAll)
	app.Get("orders/:id", c.findByID)
	app.Post("orders", c.create)
	app.Post("orders/subtotal", c.createSubTotal)
	app.Get("orders/:id/download", c.getOrderPdf)
	app.Get("orders/:id/check-download", c.checkDLOrderPdf)
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
	return nil
}

func (c *OrderController) getOrderPdf(ctx *fiber.Ctx) error {
	return nil
}

func (c *OrderController) checkDLOrderPdf(ctx *fiber.Ctx) error {
	return nil
}
