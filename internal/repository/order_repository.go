package repository

import "test-vanz-inovatif-ecommerce/internal/entity"

type OrderRepository interface {
	CreateOrder(order entity.Order) error
	BeginTransaction() (DatabaseTransactionRepository, error)
}
