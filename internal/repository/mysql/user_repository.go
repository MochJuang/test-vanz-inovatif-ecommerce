package mysql

import (
	"gorm.io/gorm"
	"hireplus-project/internal/entity"
	"hireplus-project/internal/repository"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) CreateUser(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetUserByPhone(phone string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("phone_number = ?", phone).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetUserByID(userID string) (*entity.User, error) {
	var user entity.User
	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpdateUser(user *entity.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) GetUserBalance(userID string) (float64, error) {
	var balance entity.UserBalance
	if err := r.db.Where("user_id = ?", userID).First(&balance).Error; err != nil {
		return 0, err
	}
	return balance.Balance, nil
}

func (r *userRepository) UpdateUserBalance(userID string, balance float64) error {
	userBalance := &entity.UserBalance{
		UserID:  userID,
		Balance: balance,
	}
	return r.db.Save(userBalance).Error
}
