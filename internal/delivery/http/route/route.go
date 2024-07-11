package route

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"test-vanz-inovatif-ecommerce/internal/config"
	"test-vanz-inovatif-ecommerce/internal/delivery/http"
	middleware "test-vanz-inovatif-ecommerce/internal/delivery/http/midlleware"
	"test-vanz-inovatif-ecommerce/internal/repository/mysql"
	"test-vanz-inovatif-ecommerce/internal/service"
)

func SetupRoutes(app *fiber.App, db *gorm.DB, cfg config.Config) {
	api := app.Group("/api")

	userRepo := mysql.NewUserRepository(db)
	productRepo := mysql.NewProductRepository(db)
	cartRepo := mysql.NewCartRepository(db)
	orderRepo := mysql.NewOrderRepository(db)

	userService := service.NewUserService(userRepo, cfg)
	productService := service.NewProductService(productRepo)
	cartService := service.NewCartService(cartRepo, productRepo, userRepo)
	orderService := service.NewOrderService(orderRepo, cartRepo, userRepo)

	userHandler := http.NewUserHandler(userService)
	productHandler := http.NewProductHandler(productService)
	cartHandler := http.NewCartHandler(cartService)
	orderHandler := http.NewOrderHandler(orderService)

	api.Post("/register", userHandler.Register)
	api.Post("/login", userHandler.Login)

	api.Use(middleware.AuthMiddleware(cfg))
	api.Get("/products", productHandler.GetProducts)

	api.Post("/cart", cartHandler.AddToCart)
	api.Get("/cart/:user_id", cartHandler.GetCartByUserID)
	api.Delete("/cart/:user_id", cartHandler.ClearCartByUserID)

	api.Post("/checkout/:user_id", orderHandler.Checkout)
}
