package http

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"test-vanz-inovatif-ecommerce/internal/model"
	"test-vanz-inovatif-ecommerce/internal/service"
)

type CartHandler struct {
	cartService service.CartService
}

func NewCartHandler(cs service.CartService) *CartHandler {
	return &CartHandler{cartService: cs}
}

func (h *CartHandler) AddToCart(c *fiber.Ctx) error {
	var req model.AddToCartRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	err := h.cartService.AddToCart(req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(http.StatusCreated)
}

func (h *CartHandler) GetCartByUserID(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("user_id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid user_id"})
	}

	carts, err := h.cartService.GetCartByUserID(uint(userID))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(carts)
}

func (h *CartHandler) ClearCartByUserID(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("user_id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid user_id"})
	}

	err = h.cartService.ClearCartByUserID(uint(userID))
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(http.StatusOK)
}
