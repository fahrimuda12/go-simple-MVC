package models

import (
	"time"

	"github.com/google/uuid"
)

type Tabler interface {
	TableName() string
}

type Wallets struct {
	ID     				uuid.UUID	`gorm:"primaryKey""type:uuid;default:uuid_generate_v4()"`
	Name            	string  `gorm:"not null;unique"`
	Type            	string  `gorm:"not null"`
	Balance         	float64 `gorm:"not null"`
	// Key_Phrase      	string  `gorm:"not null;unique"`
	User_ID         	uint    `gorm:"not null"`
	Virtual_Account 	string  `gorm:"not null;unique"`
	Tag_Name        	string  `gorm:"not null"`
	Created_At      	time.Time  `gorm:"autoCreateTime"`
	Updated_At      	time.Time  `gorm:"autoUpdateTime"`
}

func (Wallets) TableName() string {
	return "wallets"
}