package dao

import (
	"gorm.io/gorm"
)

type ITransactionManager interface {
	Start(txCallback TxFunc) error
}

type transactionManager struct {
}

type TxContext struct {
	Tx *gorm.DB
}

func FindTxDb(db *gorm.DB, txContext *TxContext) *gorm.DB {
	if txContext == nil {
		return db
	}
	return txContext.Tx
}

type TxFunc func(ctx *TxContext) error

func (t *transactionManager) Start(txCallback TxFunc) error {
	// 开启事务
	db := NewOrGet().GetDb()
	return db.Transaction(func(tx *gorm.DB) error {
		return txCallback(&TxContext{Tx: tx})
	})
}
func NewTransactionManager() ITransactionManager {
	return &transactionManager{}
}
