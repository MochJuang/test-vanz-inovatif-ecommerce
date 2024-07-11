package mysql

import (
	"gorm.io/gorm"
	"test-vanz-inovatif-ecommerce/internal/entity"
	"test-vanz-inovatif-ecommerce/internal/repository"
)

type cartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) repository.CartRepository {
	return &cartRepository{db: db}
}

func (r *cartRepository) AddToCart(cart entity.Cart) error {
	return r.db.Create(&cart).Error
}

func (r *cartRepository) GetCartByUserID(userID uint) ([]entity.Cart, error) {
	var carts []entity.Cart
	err := r.db.Where("user_id = ?", userID).Find(&carts).Error
	return carts, err
}

func (r *cartRepository) DeleteCartByUserID(userID uint) error {
	return r.db.Where("user_id = ?", userID).Delete(&entity.Cart{}).Error
}
