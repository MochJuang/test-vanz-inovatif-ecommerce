package repository

import "test-vanz-inovatif-ecommerce/internal/entity"

type UserRepository interface {
	Create(user entity.User) error
	FindByEmail(email string) (entity.User, error)
	FindByID(userID uint) (entity.User, error)
}
