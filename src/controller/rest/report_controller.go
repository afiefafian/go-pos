package rest

import (
	"github.com/afiefafian/go-pos/src/helper"
	"github.com/afiefafian/go-pos/src/middleware"
	"github.com/afiefafian/go-pos/src/service"
	"github.com/gofiber/fiber/v2"
)

type ReportController struct {
	reportService *service.ReportService
}

func NewReportController(reportService *service.ReportService) *ReportController {
	return &ReportController{reportService}
}

func (c *ReportController) Route(app *fiber.App) {
	app.Get("/revenues", middleware.Protected(), c.getRevenuesReport)
	app.Get("/solds", c.GetOrderedProducts)
}

func (c *ReportController) getRevenuesReport(ctx *fiber.Ctx) error {
	response, err := c.reportService.GetPaymentTypeRevenues()
	if err != nil {
		panic(err)
	}

	return helper.JsonOK(ctx, response, "")
}

func (c *ReportController) GetOrderedProducts(ctx *fiber.Ctx) error {
	products, err := c.reportService.GetOrderedProducts()
	if err != nil {
		panic(err)
	}

	return helper.JsonOK(ctx, fiber.Map{"orderProducts": products}, "")
}
