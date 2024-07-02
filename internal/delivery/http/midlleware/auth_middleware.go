package middleware

import (
	"github.com/gofiber/fiber/v2"
	"hireplus-project/internal/utils"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthenticated"})
		}

		claims, err := utils.ParseToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": err.Error()})
		}

		c.Locals("user_id", claims.UserID)
		return c.Next()
	}
}
