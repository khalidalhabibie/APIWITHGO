package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"  binding:"required" `
	Email    string `json:"email" binding:"required,email" gorm:"unique"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}

func (User) TableName() string {
	return "userss"
}
