package rest

import "github.com/gofiber/fiber/v2"

type CategoryController struct {
}

func NewCategoryController() *CategoryController {
	return &CategoryController{}
}

func (c *CategoryController) Route(app *fiber.App) {
	app.Get("categories", c.findAll)
	app.Get("categories/:id", c.findByID)
	app.Post("categories", c.create)
	app.Put("categories/:id", c.updateByID)
	app.Delete("categories/:id", c.deleteByID)
}

func (c *CategoryController) findAll(ctx *fiber.Ctx) error {
	return nil
}

func (c *CategoryController) findByID(ctx *fiber.Ctx) error {
	return nil
}

func (c *CategoryController) create(ctx *fiber.Ctx) error {
	return nil
}

func (c *CategoryController) updateByID(ctx *fiber.Ctx) error {
	return nil
}

func (c *CategoryController) deleteByID(ctx *fiber.Ctx) error {
	return nil
}
