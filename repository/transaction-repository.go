package repository

import (
	//"../model"
	"github.com/khalid/apiWithGO/model"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	TransactionProduct(int, int, int) error
}

type transactionRepository struct {
	connection *gorm.DB
}

func NewTransactionRepository() TransactionRepository {
	return &transactionRepository{
		connection: DB(),
	}
}

func (db *transactionRepository) TransactionProduct(userID int, productID int, quantity int) error {
	return db.connection.Create(&model.Transaction{
		ProductID: uint(productID),
		UserID:    uint(userID),
		Quantity:  quantity,
	}).Error

}
