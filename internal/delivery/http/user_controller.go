package http

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"test-vanz-inovatif-ecommerce/internal/model"
	"test-vanz-inovatif-ecommerce/internal/service"
)

type Handler struct {
	userService service.UserService
}

func NewUserHandler(us service.UserService) *Handler {
	return &Handler{userService: us}
}

func (h *Handler) Register(c *fiber.Ctx) error {
	var req model.UserRegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	token, err := h.userService.Register(req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"token": token})
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var req model.UserLoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	token, err := h.userService.Login(req)
	if err != nil {
		return c.Status(http.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"token": token})
}
