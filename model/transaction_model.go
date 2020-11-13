package model

import "gorm.io/gorm"

type Transaction struct{
	gorm.Model
	User 	User	`gorm:"foreignkey":UserID`
	Product Product `gorm:"foreignkey":ProductID`
	UserID 	uint	
	ProductID uint 
	Quantity int  	`json:"quantity"` 
}

func (Transaction) TableName() string{
	return "transactions"
}