package rest

import "github.com/gofiber/fiber/v2"

type ProductController struct {
}

func NewProductController() *ProductController {
	return &ProductController{}
}

func (c *ProductController) Route(app *fiber.App) {
	app.Get("products", c.findAll)
	app.Get("products/:id", c.findByID)
	app.Post("products", c.create)
	app.Put("products/:id", c.updateByID)
	app.Delete("products/:id", c.deleteByID)
}

func (c *ProductController) findAll(ctx *fiber.Ctx) error {
	return nil
}

func (c *ProductController) findByID(ctx *fiber.Ctx) error {
	return nil
}

func (c *ProductController) create(ctx *fiber.Ctx) error {
	return nil
}

func (c *ProductController) updateByID(ctx *fiber.Ctx) error {
	return nil
}

func (c *ProductController) deleteByID(ctx *fiber.Ctx) error {
	return nil
}
