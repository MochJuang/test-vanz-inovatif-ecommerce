package service

import (
	"test-vanz-inovatif-ecommerce/internal/entity"
	"test-vanz-inovatif-ecommerce/internal/repository"
)

type OrderService interface {
	Checkout(userID uint) error
}

type orderService struct {
	userRepo  repository.UserRepository
	orderRepo repository.OrderRepository
	cartRepo  repository.CartRepository
}

func NewOrderService(or repository.OrderRepository, cr repository.CartRepository, usr repository.UserRepository) OrderService {
	return &orderService{orderRepo: or, cartRepo: cr, userRepo: usr}
}

func (s *orderService) Checkout(userID uint) error {
	var orderItems []entity.OrderItem
	var total float64

	var user entity.User
	user, err := s.userRepo.FindByID(userID)
	if err != nil {
		return err
	}

	var carts []entity.Cart
	carts, err = s.cartRepo.GetCartByUserID(userID)
	if err != nil {
		return err
	}
	for _, cart := range carts {
		price := cart.Product.Price
		orderItem := entity.OrderItem{
			ProductID: cart.ProductID,
			Quantity:  cart.Quantity,
			Price:     price,
		}
		orderItems = append(orderItems, orderItem)
		total += price * float64(cart.Quantity)

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
