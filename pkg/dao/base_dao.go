package dao

import "gorm.io/gorm"

type IBaseDao[E interface{}] interface {
	FindDb(tx *TxContext, db *gorm.DB) *gorm.DB
	Insert(e *E, txContext *TxContext) (bool, error, uint64)

	Update(e *E, txContext *TxContext) (bool, error)

	Delete(e *E, txContext *TxContext) (bool, error)

	SelectOne(id uint64) *E

	SelectList(conditions ...interface{}) []E
}

type BaseDao[E interface{}] struct {
}

func (b *BaseDao[E]) FindDb(tx *TxContext, db *gorm.DB) *gorm.DB {
	if tx == nil || tx.Tx == nil {
		return db
	}
	return tx.Tx
}
