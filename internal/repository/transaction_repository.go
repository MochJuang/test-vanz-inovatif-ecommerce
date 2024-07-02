package repository

import "hireplus-project/internal/entity"

type TransactionRepository interface {
	CreateTransaction(transaction *entity.Transaction) error
	GetTransactionsByUserID(userID string) ([]entity.Transaction, error)
	BeginTransaction() (DatabaseTransactionRepository, error)
}
