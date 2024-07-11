package repository

import "test-vanz-inovatif-ecommerce/internal/entity"

type CartRepository interface {
	AddToCart(cart entity.Cart) error
	GetCartByUserID(userID uint) ([]entity.Cart, error)
	DeleteCartByUserID(userID uint) error
}
