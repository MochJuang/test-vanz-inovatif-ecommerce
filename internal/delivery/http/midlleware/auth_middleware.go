package middleware

import (
	"github.com/gofiber/fiber/v2"
	"strings"
	"test-vanz-inovatif-ecommerce/internal/config"
	"test-vanz-inovatif-ecommerce/internal/utils"
)

func AuthMiddleware(cfg config.Config) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthenticated"})
		}
		token = strings.Replace(token, "Bearer ", "", 1)

		claims, err := utils.ParseToken(token, cfg.JWTSecret)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": err.Error()})
		}

		c.Locals("user_id", claims.UserID)
		return c.Next()
	}
}
