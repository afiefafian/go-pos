package rest

import "github.com/gofiber/fiber/v2"

type PaymentController struct {
}

func NewPaymentController() *PaymentController {
	return &PaymentController{}
}

func (c *PaymentController) Route(app *fiber.App) {
	app.Get("payments", c.findAll)
	app.Get("payments/:id", c.findByID)
	app.Post("payments", c.create)
	app.Put("payments/:id", c.updateByID)
	app.Delete("payments/:id", c.deleteByID)
}

func (c *PaymentController) findAll(ctx *fiber.Ctx) error {
	return nil
}

func (c *PaymentController) findByID(ctx *fiber.Ctx) error {
	return nil
}

func (c *PaymentController) create(ctx *fiber.Ctx) error {
	return nil
}

func (c *PaymentController) updateByID(ctx *fiber.Ctx) error {
	return nil
}

func (c *PaymentController) deleteByID(ctx *fiber.Ctx) error {
	return nil
}
