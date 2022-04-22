package rest

import "github.com/gofiber/fiber/v2"

type ReportController struct {
}

func NewAuthController() *ReportController {
	return &ReportController{}
}

func (c *ReportController) Route(app *fiber.App) {
	app.Post("revenues", c.getRevenuesReport)
	app.Post("solds", c.getSoldsReport)
}

func (c *ReportController) getRevenuesReport(ctx *fiber.Ctx) error {
	return nil
}

func (c *ReportController) getSoldsReport(ctx *fiber.Ctx) error {
	return nil
}
