package mysql

import (
	"gorm.io/gorm"
	"test-vanz-inovatif-ecommerce/internal/entity"
	"test-vanz-inovatif-ecommerce/internal/repository"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) repository.ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) GetAllProducts() ([]entity.Product, error) {
	var products []entity.Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *productRepository) GetProductByID(id uint) (entity.Product, error) {
	var product entity.Product
	result := r.db.First(&product, id)
	if result.Error != nil {
		return entity.Product{}, result.Error
	}
	return product, nil
}
