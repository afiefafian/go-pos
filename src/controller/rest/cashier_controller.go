package rest

import (
	"github.com/afiefafian/go-pos/src/service"
	"github.com/gofiber/fiber/v2"
)

type CashierController struct {
	cashierController     *service.CashierService
	cashierAuthController *service.CashierAuthService
}

func NewCashierController(
	cashierController *service.CashierService,
	cashierAuthController *service.CashierAuthService,
) *CashierController {
	return &CashierController{
		cashierController,
		cashierAuthController,
	}
}

func (c *CashierController) Route(app *fiber.App) {
	route := app.Group("cashiers")

	route.Get("/", c.findAll)
	route.Get("/:id", c.findByID)
	route.Post("/", c.create)
	route.Put("/:id", c.updateByID)
	route.Delete("/:id", c.deleteByID)
	route.Get("/:id/passcode", c.getPasscode)
	route.Post("/:id/login", c.login)
	route.Get("/:id/logout", c.logout)
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
