package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
	"test-vanz-inovatif-ecommerce/internal/config"
	"test-vanz-inovatif-ecommerce/internal/delivery/http/route"
	"test-vanz-inovatif-ecommerce/internal/repository/mysql"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	// Initialize database
	db, err := mysql.NewConnector(cfg)
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Perform database migration
	err = mysql.Migrate(db)
	if err != nil {
		log.Fatalf("Could not migrate database: %v", err)
	}

	app := fiber.New()

	app.Use(logger.New())

	// Setup routes
	route.SetupRoutes(app, db, cfg)

	// Start server
	log.Fatal(app.Listen(cfg.ServerAddress))
}
