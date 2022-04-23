package main

import (
	"fmt"
	"log"

	"github.com/afiefafian/go-pos/src/config"
	"github.com/afiefafian/go-pos/src/controller/rest"
	"github.com/afiefafian/go-pos/src/middleware"
	"github.com/afiefafian/go-pos/src/repository"
	"github.com/afiefafian/go-pos/src/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	cfg := config.Load()

	db := config.NewMySQLDatabase(cfg)
	if db != nil {
		defer db.Close()
	}

	// Setup Repository
	cashierRepository := repository.NewCashierRepository(db)

	// Setup Service
	cashierService := service.NewCashierService(cashierRepository)
	cashierAuthService := service.NewCashierAuthService(cashierRepository)

	// Setup Controller
	cashierController := rest.NewCashierController(cashierService, cashierAuthService)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	// Setup Routing
	cashierController.Route(app)

	// Handle not founds
	app.Use(middleware.RouteNotFound)

	// Start App
	if err := app.Listen(fmt.Sprintf(":%s", cfg.Getenv("PORT", "3030"))); err != nil {
		log.Fatal(err)
	}
}
