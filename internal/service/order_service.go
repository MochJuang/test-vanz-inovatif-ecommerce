package service

import (
	"test-vanz-inovatif-ecommerce/internal/entity"
	"test-vanz-inovatif-ecommerce/internal/model"
	"test-vanz-inovatif-ecommerce/internal/repository"
)

type OrderService interface {
	Checkout(userID uint, items []model.AddToCartRequest) error
}

type orderService struct {
	userRepo    repository.UserRepository
	orderRepo   repository.OrderRepository
	cartRepo    repository.CartRepository
	productRepo repository.ProductRepository
}

func NewOrderService(or repository.OrderRepository, cr repository.CartRepository, usr repository.UserRepository, productRepo repository.ProductRepository) OrderService {
	return &orderService{orderRepo: or, cartRepo: cr, userRepo: usr, productRepo: productRepo}
}

func (s *orderService) Checkout(userID uint, items []model.AddToCartRequest) error {
	var orderItems []entity.OrderItem
	var total float64

	for _, item := range items {
		price := 100.0 // Assuming each item has a fixed price for simplicity
		orderItem := entity.OrderItem{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Price:     price,
		}
		orderItems = append(orderItems, orderItem)
		total += price * float64(item.Quantity)
	}

	var user entity.User
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return err
	}

	order := entity.Order{
		UserID:     user.ID,
		Total:      total,
		OrderItems: orderItems,
	}
	trx, err := s.orderRepo.BeginTransaction()
	if err != nil {
		panic(err)
	}

	err = s.orderRepo.CreateOrder(order)
	if err != nil {
		return err
	}

	err = s.cartRepo.DeleteCartByUserID(userID)
	if err != nil {
		trx.RollbackTransaction()
		return err
	}

	return trx.CommitTransaction()
}
