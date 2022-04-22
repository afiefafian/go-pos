package rest

import "github.com/gofiber/fiber/v2"

type CashierController struct {
}

func NewCashierController() *CashierController {
	return &CashierController{}
}

func (c *CashierController) Route(app *fiber.App) {
	app.Get("cashiers", c.findAll)
	app.Get("cashiers/:id", c.findByID)
	app.Post("cashiers", c.create)
	app.Put("cashiers/:id", c.updateByID)
	app.Delete("cashiers/:id", c.deleteByID)
	app.Get("cashiers/:id/passcode", c.getPasscode)
	app.Post("cashiers/:id/login", c.login)
	app.Get("cashiers/:id/logout", c.logout)
}

func (c *CashierController) findAll(ctx *fiber.Ctx) error {
	return nil
}

func (c *CashierController) findByID(ctx *fiber.Ctx) error {
	return nil
}

func (c *CashierController) create(ctx *fiber.Ctx) error {
	return nil
}

func (c *CashierController) updateByID(ctx *fiber.Ctx) error {
	return nil
}

func (c *CashierController) deleteByID(ctx *fiber.Ctx) error {
	return nil
}

func (c *CashierController) getPasscode(ctx *fiber.Ctx) error {
	return nil
}

func (c *CashierController) login(ctx *fiber.Ctx) error {
	return nil
}

func (c *CashierController) logout(ctx *fiber.Ctx) error {
	return nil
}
