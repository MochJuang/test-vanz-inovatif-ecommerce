package mysql

import "gorm.io/gorm"

type dbTransaction struct {
	tx *gorm.DB
}

func (t *dbTransaction) CommitTransaction() error {
	return t.tx.Commit().Error
}

func (t *dbTransaction) RollbackTransaction() error {
	return t.tx.Rollback().Error
}
