package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"

	"github.com/afiefafian/go-pos/src/config"
	"github.com/afiefafian/go-pos/src/controller/rest"
	"github.com/afiefafian/go-pos/src/helper"
	"github.com/afiefafian/go-pos/src/middleware"
	"github.com/afiefafian/go-pos/src/model"
	"github.com/afiefafian/go-pos/src/repository"
	"github.com/afiefafian/go-pos/src/service"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/pressly/goose/v3"
)

var fiberConfig = fiber.Config{
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		newErr := fiber.ErrInternalServerError

		if e, ok := err.(*fiber.Error); ok {
			newErr.Code = e.Code
			newErr.Message = e.Error()
		} else if e, ok := err.(error); ok {
			newErr = fiber.ErrBadRequest
			newErr.Message = e.Error()
		}

		if e, ok := err.(validator.ValidationErrors); ok {
			errDetail := helper.FormatValidationError(e)

			return ctx.Status(fiber.StatusBadRequest).JSON(model.BaseResponse{
				Success: false,
				Message: helper.MsgErrValidation,
				Error:   errDetail,
			})
		}

		return ctx.Status(newErr.Code).JSON(model.BaseResponse{
			Success: false,
			Message: newErr.Error(),
			Error:   &model.EmptyObject{},
		})
	},
}

//go:embed migrations/*.sql
var migrationFs embed.FS

func runMigration(db *sql.DB) {
	goose.SetBaseFS(migrationFs)

	if err := goose.SetDialect("mysql"); err != nil {
		panic(err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}
}

func main() {
	cfg := config.Load()

	mysql := config.NewMySQLDatabase(cfg)
	if mysql.DB != nil {
		defer mysql.DB.Close()
	}

	mysql.RunMigration(migrationFs)

	// Setup Repository
	cashierRepository := repository.NewCashierRepository(mysql.DB)
	categoryRepository := repository.NewCategoryRepository(mysql.DB)
	paymentRepository := repository.NewPaymentRepository(mysql.DB)

	// Setup Service
	cashierService := service.NewCashierService(cashierRepository)
	cashierAuthService := service.NewCashierAuthService(cashierRepository)
	categoryService := service.NewCategoryService(categoryRepository)
	paymentService := service.NewPaymentService(paymentRepository)

	// Setup Controller
	cashierController := rest.NewCashierController(cashierService, cashierAuthService)
	categoryController := rest.NewCategoryController(categoryService)
	paymentController := rest.NewPaymentController(paymentService)

	// Setup Fiber
	app := fiber.New(fiberConfig)
	app.Use(recover.New())
	// Setup Routing
	cashierController.Route(app)
	categoryController.Route(app)
	paymentController.Route(app)
	// Handle not founds
	app.Use(middleware.RouteNotFound)

	// Start App
	if err := app.Listen(fmt.Sprintf(":%s", cfg.Getenv("PORT", "3030"))); err != nil {
		log.Fatal(err)
	}
}
