package http

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"test-vanz-inovatif-ecommerce/internal/model"
	"test-vanz-inovatif-ecommerce/internal/service"
)

type OrderHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(os service.OrderService) *OrderHandler {
	return &OrderHandler{orderService: os}
}

func (h *OrderHandler) Checkout(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("user_id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid user_id"})
	}

	var items []model.AddToCartRequest
	if err := c.BodyParser(&items); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err = h.orderService.Checkout(uint(userID), items)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(http.StatusOK)
}
