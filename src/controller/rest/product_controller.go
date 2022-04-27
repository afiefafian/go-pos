package rest

import (
	"strconv"
	"strings"

	"github.com/afiefafian/go-pos/src/exception"
	"github.com/afiefafian/go-pos/src/helper"
	"github.com/afiefafian/go-pos/src/middleware"
	"github.com/afiefafian/go-pos/src/model"
	"github.com/afiefafian/go-pos/src/service"
	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	productService *service.ProductService
}

func NewProductController(productService *service.ProductService) *ProductController {
	return &ProductController{productService}
}

func (c *ProductController) Route(app *fiber.App) {
	route := app.Group("/products")

	route.Get("/", middleware.Protected(), c.findAll)
	route.Get("/:id", middleware.Protected(), c.findByID)
	route.Post("/", c.create)
	route.Put("/:id", c.updateByID)
	route.Delete("/:id", c.deleteByID)
}

func (c *ProductController) findAll(ctx *fiber.Ctx) error {
	// Query Params
	strCatId := ctx.Query("categoryId")
	querySearch := ctx.Query("q")

	var categoryID int64
	if strCatId != "" {
		categoryID, _ = strconv.ParseInt(strCatId, 10, 64)
	}

	querySearch = strings.TrimSpace(querySearch)

	paginationQuery := helper.NewPaginationQueryFromCtx(ctx)
	query := model.GetProductQuery{
		Pagination: *paginationQuery,
	}

	if categoryID != 0 {
		query.CategoryId = &categoryID
	}

	if querySearch != "" {
		query.Q = &querySearch
	}

	// Get data from service
	products, pagination, err := c.productService.FindAll(query)
	if err != nil {
		panic(err)
	}

	// Return result
	return helper.JsonOK(ctx, map[string]interface{}{
		"products": products,
		"meta":     pagination,
	}, "")
}

func (c *ProductController) findByID(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		panic(exception.EntityNotFound("Product"))
	}

	var response *model.GetProductResponse
	response, err = c.productService.GetByID(id)
	if err != nil {
		panic(err)
	}

	return helper.JsonOK(ctx, response, "")
}

func (c *ProductController) create(ctx *fiber.Ctx) error {
	var request model.CreateProductRequest
	if err := ctx.BodyParser(&request); err != nil {
		panic(err)
	}

	response, err := c.productService.Create(request)
	if err != nil {
		panic(err)
	}

	return helper.JsonOK(ctx, response, "")
}

func (c *ProductController) updateByID(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		panic(exception.EntityNotFound("Product"))
	}

	var request model.UpdateProductRequest
	if err := ctx.BodyParser(&request); err != nil {
		panic(err)
	}

	request.ID = id

	err = c.productService.UpdateByID(request)
	if err != nil {
		panic(err)
	}

	return helper.JsonOK(ctx, nil, "")
}

func (c *ProductController) deleteByID(ctx *fiber.Ctx) error {
	id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
	if err != nil {
		panic(exception.EntityNotFound("Product"))
	}

	err = c.productService.DeleteByID(id)
	if err != nil {
		panic(err)
	}

	return helper.JsonOK(ctx, nil, "")
}
