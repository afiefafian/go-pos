package main

import (
	"database/sql"
	"embed"
	"fmt"
	"log"

	"github.com/afiefafian/go-pos/src/config"
	"github.com/afiefafian/go-pos/src/controller/rest"
	"github.com/afiefafian/go-pos/src/middleware"
	"github.com/afiefafian/go-pos/src/repository"
	"github.com/afiefafian/go-pos/src/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/pressly/goose/v3"
)

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
