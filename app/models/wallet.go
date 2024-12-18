package models

import (
	"gorm.io/gorm"
)

type Tabler interface {
	TableName() string
}

type Wallets struct {
	gorm.Model
	ID     				string	`gorm:"primaryKey"`
	Name            	string  `gorm:"not null;unique"`
	Type            	string  `gorm:"not null"`
	Balance         	float64 `gorm:"not null"`
	Key_Phrase      	string  `gorm:"not null;unique"`
	User_ID         	uint    `gorm:"not null"`
	Virtual_Account 	string  `gorm:"not null;unique"`
	Tag_Name        	string  `gorm:"not null"`
}

func (Wallets) TableName() string {
	return "wallets"
}