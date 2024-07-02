package service

import (
	"fmt"
	"hireplus-project/internal/entity"
	"hireplus-project/internal/repository"
	"hireplus-project/internal/utils"
	"time"

	"github.com/google/uuid"
)

type UserService interface {
	Register(firstName, lastName, phone, address, pin string) (*entity.User, error)
	Login(phone, pin string) (string, string, error)
	UpdateProfile(userID, firstName, lastName, address string) (*entity.User, error)
	GetUserBalance(userID string) (float64, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo}
}

func (s *userService) Register(firstName, lastName, phone, address, pin string) (*entity.User, error) {
	user := &entity.User{
		ID:          uuid.New().String(),
		FirstName:   firstName,
		LastName:    lastName,
		PhoneNumber: phone,
		Address:     address,
		Pin:         pin,
		CreatedAt:   time.Now(),
	}
	if err := s.userRepo.CreateUser(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) Login(phone, pin string) (string, string, error) {
	user, err := s.userRepo.GetUserByPhone(phone)
	if err != nil {
		return "", "", fmt.Errorf("phone number or PIN is incorrect")
	}

	if user.Pin != pin {
		return "", "", fmt.Errorf("phone number or PIN is incorrect")
	}

	accessToken, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", "", err
	}

	refreshToken, err := utils.GenerateToken(user.ID)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *userService) UpdateProfile(userID, firstName, lastName, address string) (*entity.User, error) {
	user, err := s.userRepo.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	user.FirstName = firstName
	user.LastName = lastName
	user.Address = address
	user.UpdatedAt = time.Now()

	if err := s.userRepo.UpdateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) GetUserBalance(userID string) (float64, error) {
	return s.userRepo.GetUserBalance(userID)
}
