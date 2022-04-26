package rest

import (
	"strconv"

	"github.com/afiefafian/go-pos/src/exception"
	"github.com/afiefafian/go-pos/src/helper"
	"github.com/afiefafian/go-pos/src/model"
	"github.com/afiefafian/go-pos/src/service"
	"github.com/gofiber/fiber/v2"
)

type CashierController struct {
	cashierService     *service.CashierService
	cashierAuthService *service.CashierAuthService
}

func NewCashierController(
	cashierService *service.CashierService,
	cashierAuthService *service.CashierAuthService,
) *CashierController {
	return &CashierController{
		cashierService,
		cashierAuthService,
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
	query := helper.NewPaginationQueryFromCtx(ctx)
	cashiers, pagination, err := c.cashierService.FindAll(query)
	if err != nil {
		panic(err)
	}

	return helper.JsonOK(ctx, map[string]interface{}{
		"cashiers": cashiers,
		"meta":     pagination,
	}, "")
}

func (c *CashierController) findByID(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		panic(exception.EntityNotFound("Cashier"))
	}

	var response *model.GetCashierResponse
	response, err = c.cashierService.GetByID(id)
	if err != nil {
		panic(err)
	}

	return helper.JsonOK(ctx, response, "")
}

func (c *CashierController) create(ctx *fiber.Ctx) error {
	var request model.CreateCashierRequest
	if err := ctx.BodyParser(&request); err != nil {
		panic(err)
	}

	response, err := c.cashierService.Create(request)
	if err != nil {
		panic(err)
	}

	return helper.JsonOK(ctx, response, "")
}

func (c *CashierController) updateByID(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		panic(exception.EntityNotFound("Cashier"))
	}

	var request model.UpdateCashierRequest
	if err := ctx.BodyParser(&request); err != nil {
		panic(err)
	}

	request.ID = id

	err = c.cashierService.UpdateByID(request)
	if err != nil {
		panic(err)
	}

	return helper.JsonOK(ctx, nil, "")
}

func (c *CashierController) deleteByID(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		panic(exception.EntityNotFound("Cashier"))
	}

	err = c.cashierService.DeleteByID(id)
	if err != nil {
		panic(err)
	}

	return helper.JsonOK(ctx, nil, "")
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
