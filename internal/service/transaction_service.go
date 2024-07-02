package service

import (
	"fmt"
	"hireplus-project/internal/entity"
	"hireplus-project/internal/repository"
	"time"

	"github.com/google/uuid"
)

type TransactionService interface {
	TopUp(userID string, amount float64) (*entity.Transaction, error)
	Payment(userID, remarks string, amount float64) (*entity.Transaction, error)
	Transfer(userID, targetUserID, remarks string, amount float64) (*entity.Transaction, error)
	TransactionsReport(userID string) ([]entity.Transaction, error)
}

type transactionService struct {
	transactionRepo repository.TransactionRepository
	userRepo        repository.UserRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository, userRepo repository.UserRepository) TransactionService {
	return &transactionService{transactionRepo, userRepo}
}

func (s *transactionService) TopUp(userID string, amount float64) (*entity.Transaction, error) {
	txRepo, err := s.transactionRepo.BeginTransaction()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txRepo.RollbackTransaction()
		}
	}()

	transaction := &entity.Transaction{
		ID:              uuid.New().String(),
		UserID:          userID,
		Amount:          amount,
		TransactionType: entity.CREDIT,
		Remarks:         "Top Up",
		CreatedAt:       time.Now(),
	}
	if err := s.transactionRepo.CreateTransaction(transaction); err != nil {
		txRepo.RollbackTransaction()
		return nil, err
	}

	balance, err := s.userRepo.GetUserBalance(userID)
	if err != nil {
		txRepo.RollbackTransaction()
		return nil, err
	}
	balance += amount
	if err := s.userRepo.UpdateUserBalance(userID, balance); err != nil {
		txRepo.RollbackTransaction()
		return nil, err
	}

	if err := txRepo.CommitTransaction(); err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *transactionService) Payment(userID, remarks string, amount float64) (*entity.Transaction, error) {
	txRepo, err := s.transactionRepo.BeginTransaction()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txRepo.RollbackTransaction()
		}
	}()

	balance, err := s.userRepo.GetUserBalance(userID)
	if err != nil {
		txRepo.RollbackTransaction()
		return nil, err
	}
	if balance < amount {
		txRepo.RollbackTransaction()
		return nil, fmt.Errorf("balance is not enough")
	}

	transaction := &entity.Transaction{
		ID:              uuid.New().String(),
		UserID:          userID,
		Amount:          amount,
		TransactionType: entity.DEBIT,
		Remarks:         remarks,
		CreatedAt:       time.Now(),
	}
	if err := s.transactionRepo.CreateTransaction(transaction); err != nil {
		txRepo.RollbackTransaction()
		return nil, err
	}

	balance -= amount
	if err := s.userRepo.UpdateUserBalance(userID, balance); err != nil {
		txRepo.RollbackTransaction()
		return nil, err
	}

	if err := txRepo.CommitTransaction(); err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *transactionService) Transfer(userID, targetUserID, remarks string, amount float64) (*entity.Transaction, error) {
	txRepo, err := s.transactionRepo.BeginTransaction()
	if err != nil {
		return nil, err
	}
	defer func() {
		if r := recover(); r != nil {
			txRepo.RollbackTransaction()
		}
	}()

	balance, err := s.userRepo.GetUserBalance(userID)
	if err != nil {
		txRepo.RollbackTransaction()
		return nil, err
	}
	if balance < amount {
		txRepo.RollbackTransaction()
		return nil, fmt.Errorf("balance is not enough")
	}

	debitTransaction := &entity.Transaction{
		ID:              uuid.New().String(),
		UserID:          userID,
		Amount:          amount,
		TransactionType: entity.DEBIT,
		Remarks:         "Transfer to " + targetUserID + ": " + remarks,
		CreatedAt:       time.Now(),
	}
	if err := s.transactionRepo.CreateTransaction(debitTransaction); err != nil {
		txRepo.RollbackTransaction()
		return nil, err
	}

	creditTransaction := &entity.Transaction{
		ID:              uuid.New().String(),
		UserID:          targetUserID,
		Amount:          amount,
		TransactionType: entity.CREDIT,
		Remarks:         "Transfer from " + userID + ": " + remarks,
		CreatedAt:       time.Now(),
	}
	if err := s.transactionRepo.CreateTransaction(creditTransaction); err != nil {
		txRepo.RollbackTransaction()
		return nil, err
	}

	balance -= amount
	if err := s.userRepo.UpdateUserBalance(userID, balance); err != nil {
		txRepo.RollbackTransaction()
		return nil, err
	}
	targetBalance, err := s.userRepo.GetUserBalance(targetUserID)
	if err != nil {
		txRepo.RollbackTransaction()
		return nil, err
	}
	targetBalance += amount
	if err := s.userRepo.UpdateUserBalance(targetUserID, targetBalance); err != nil {
		txRepo.RollbackTransaction()
		return nil, err
	}

	if err := txRepo.CommitTransaction(); err != nil {
		return nil, err
	}

	return debitTransaction, nil
}

func (s *transactionService) TransactionsReport(userID string) ([]entity.Transaction, error) {
	return s.transactionRepo.GetTransactionsByUserID(userID)
}
