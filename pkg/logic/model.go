package logic

import (
	"gorm.io/gorm"
)

type TxFunc func(tx *gorm.DB) error

func Transaction(db *gorm.DB, fcs []TxFunc) (bizErr, txErr error) {
	tx := db.Begin()

	for _, fc := range fcs {
		if bizErr = fc(tx); bizErr != nil {
			return bizErr, tx.Rollback().Error
		}
	}

	return bizErr, tx.Commit().Error
}

type IDBModel[T interface{}] interface {
	Conn(db *gorm.DB) T
}

type DBModel struct {
	DB *gorm.DB
}