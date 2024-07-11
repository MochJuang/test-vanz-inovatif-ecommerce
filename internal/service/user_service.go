package service

import (
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"test-vanz-inovatif-ecommerce/internal/config"
	"test-vanz-inovatif-ecommerce/internal/entity"
	"test-vanz-inovatif-ecommerce/internal/model"
	"test-vanz-inovatif-ecommerce/internal/repository"
	"test-vanz-inovatif-ecommerce/internal/utils"
)

type UserService interface {
	Register(req model.UserRegisterRequest) (string, error)
	Login(req model.UserLoginRequest) (string, error)
}

type userService struct {
	userRepo repository.UserRepository
	config   config.Config
}

func NewUserService(ur repository.UserRepository, cfg config.Config) UserService {
	return &userService{userRepo: ur, config: cfg}
}

func (s *userService) Register(req model.UserRegisterRequest) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user := entity.User{
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return "", err
	}

	token, err := utils.GenerateToken(strconv.Itoa(int(user.ID)), s.config.JWTSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) Login(req model.UserLoginRequest) (string, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return "", err
	}

	token, err := utils.GenerateToken(strconv.Itoa(int(user.ID)), s.config.JWTSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}
