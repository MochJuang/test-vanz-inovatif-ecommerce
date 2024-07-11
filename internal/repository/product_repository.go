package repository

import "test-vanz-inovatif-ecommerce/internal/entity"

type ProductRepository interface {
	GetAllProducts() ([]entity.Product, error)
	GetProductByID(id uint) (entity.Product, error)
}
