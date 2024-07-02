package repository

import "hireplus-project/internal/entity"

type UserRepository interface {
	CreateUser(user *entity.User) error
	GetUserByPhone(phone string) (*entity.User, error)
	GetUserByID(userID string) (*entity.User, error)
	UpdateUser(user *entity.User) error
	GetUserBalance(userID string) (float64, error)
	UpdateUserBalance(userID string, balance float64) error
}
