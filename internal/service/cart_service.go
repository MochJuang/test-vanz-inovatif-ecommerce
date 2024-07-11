package service

import (
	"errors"
	"test-vanz-inovatif-ecommerce/internal/entity"
	"test-vanz-inovatif-ecommerce/internal/model"
	"test-vanz-inovatif-ecommerce/internal/repository"
)

type CartService interface {
	AddToCart(req model.AddToCartRequest) error
	GetCartByUserID(userID uint) ([]entity.Cart, error)
	ClearCartByUserID(userID uint) error
}

type cartService struct {
	cartRepo    repository.CartRepository
	userRepo    repository.UserRepository
	productRepo repository.ProductRepository
}

func NewCartService(cr repository.CartRepository, ps repository.ProductRepository, usr repository.UserRepository) CartService {
	return &cartService{cartRepo: cr, productRepo: ps, userRepo: usr}
}

func (s *cartService) AddToCart(req model.AddToCartRequest) error {
	var product entity.Product
	product, err := s.productRepo.GetProductByID(req.ProductID)
	if err != nil {
		return err
	}

	if product.Stock < req.Quantity {
		return errors.New("insufficient stock")
	}

	var user entity.User
	user, err = s.userRepo.FindByID(req.UserID)
	if err != nil {
		return err
	}

	cart := entity.Cart{
		UserID:    user.ID,
		ProductID: product.ID,
		Quantity:  req.Quantity,
	}
	return s.cartRepo.AddToCart(cart)
}

func (s *cartService) GetCartByUserID(userID uint) ([]entity.Cart, error) {
	return s.cartRepo.GetCartByUserID(userID)
}

func (s *cartService) ClearCartByUserID(userID uint) error {
	return s.cartRepo.DeleteCartByUserID(userID)
}
