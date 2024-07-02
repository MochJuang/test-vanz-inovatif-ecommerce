package mysql

import (
	"gorm.io/gorm"
	"hireplus-project/internal/entity"
	"hireplus-project/internal/repository"
)

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) repository.TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) CreateTransaction(transaction *entity.Transaction) error {
	return r.db.Create(transaction).Error
}

func (r *transactionRepository) GetTransactionsByUserID(userID string) ([]entity.Transaction, error) {
	var transactions []entity.Transaction
	if err := r.db.Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}

func (r *transactionRepository) BeginTransaction() (repository.DatabaseTransactionRepository, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return &dbTransaction{tx}, nil
}
