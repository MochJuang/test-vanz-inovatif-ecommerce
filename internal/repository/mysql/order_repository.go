package mysql

import (
	"gorm.io/gorm"
	"test-vanz-inovatif-ecommerce/internal/entity"
	"test-vanz-inovatif-ecommerce/internal/repository"
)

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) repository.OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) CreateOrder(order entity.Order) error {

	if err := r.db.Create(&order).Error; err != nil {
		return err
	}

	return order
}

func (r *orderRepository) BeginTransaction() (repository.DatabaseTransactionRepository, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &dbTransaction{tx}, nil
}
