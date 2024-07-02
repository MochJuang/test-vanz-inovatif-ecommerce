package controller

import (
	"hireplus-project/internal/model"
	"hireplus-project/internal/service"

	"github.com/gofiber/fiber/v2"
)

type TransactionController struct {
	transactionService service.TransactionService
}

func NewTransactionController(transactionService service.TransactionService) *TransactionController {
	return &TransactionController{transactionService}
}

func (h *TransactionController) TopUp(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	var req model.TopUpRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	transaction, err := h.transactionService.TopUp(userID, req.Amount)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "SUCCESS", "result": transaction})
}

func (h *TransactionController) Payment(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	var req model.PaymentRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	transaction, err := h.transactionService.Payment(userID, req.Remarks, req.Amount)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "SUCCESS", "result": transaction})
}

func (h *TransactionController) Transfer(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	var req model.TransferRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
	}

	transaction, err := h.transactionService.Transfer(userID, req.TargetUserID, req.Remarks, req.Amount)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "SUCCESS", "result": transaction})
}

func (h *TransactionController) TransactionsReport(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	transactions, err := h.transactionService.TransactionsReport(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "SUCCESS", "result": transactions})
}
