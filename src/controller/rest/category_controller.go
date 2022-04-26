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

type CategoryController struct {
	categoryService *service.CategoryService
}

func NewCategoryController(categoryService *service.CategoryService) *CategoryController {
	return &CategoryController{categoryService}
}

func (c *CategoryController) Route(app *fiber.App) {
	route := app.Group("/categories")

	route.Get("/", middleware.Protected(), c.findAll)
	route.Get("/:id", middleware.Protected(), c.findByID)
	route.Post("/", c.create)
	route.Put("/:id", c.updateByID)
	route.Delete("/:id", c.deleteByID)
}

func (c *CategoryController) findAll(ctx *fiber.Ctx) error {
	query := helper.NewPaginationQueryFromCtx(ctx)
	cashiers, pagination, err := c.categoryService.FindAll(query)
	if err != nil {
		panic(err)
	}

	return helper.JsonOK(ctx, map[string]interface{}{
		"categories": cashiers,
		"meta":       pagination,
	}, "")
}

func (c *CategoryController) findByID(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		panic(exception.EntityNotFound("Category"))
	}

	var response *model.GetCategoryResponse
	response, err = c.categoryService.GetByID(id)
	if err != nil {
		panic(err)
	}

	return helper.JsonOK(ctx, response, "")
}

func (c *CategoryController) create(ctx *fiber.Ctx) error {
	var request model.CreateCategoryRequest
	if err := ctx.BodyParser(&request); err != nil {
		panic(err)
	}

	response, err := c.categoryService.Create(request)
	if err != nil {
		panic(err)
	}

	return helper.JsonOK(ctx, response, "")
}

func (c *CategoryController) updateByID(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		panic(exception.EntityNotFound("Category"))
	}

	var request model.UpdateCategoryRequest
	if err := ctx.BodyParser(&request); err != nil {
		panic(err)
	}

	request.ID = id

	err = c.categoryService.UpdateByID(request)
	if err != nil {
		panic(err)
	}

	return helper.JsonOK(ctx, nil, "")
}

func (c *CategoryController) deleteByID(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		panic(exception.EntityNotFound("Category"))
	}

	err = c.categoryService.DeleteByID(id)
	if err != nil {
		panic(err)
	}

	return helper.JsonOK(ctx, nil, "")
}
