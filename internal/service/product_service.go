package service

import (
	"test-vanz-inovatif-ecommerce/internal/entity"
	"test-vanz-inovatif-ecommerce/internal/repository"
)

type ProductService interface {
	GetAllProducts() ([]entity.Product, error)
}

type productService struct {
	productRepo repository.ProductRepository
}

func NewProductService(pr repository.ProductRepository) ProductService {
	return &productService{productRepo: pr}
}

func (s *productService) GetAllProducts() ([]entity.Product, error) {
	return s.productRepo.GetAllProducts()
}
