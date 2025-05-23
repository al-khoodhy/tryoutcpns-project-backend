package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	ID            uint    `gorm:"primaryKey"`
	UserID        uint    `gorm:"not null"`
	PackageID     uint    `gorm:"not null"`
	Amount        float64 `gorm:"type:decimal(10,2);not null"`
	PaymentMethod string  `gorm:"type:enum('manual','midtrans','doku');not null"`
	Status        string  `gorm:"type:enum('pending','success','failed');default:'pending';not null"`
	TransactionID string  `gorm:"type:varchar(100)"`
	CreatedAt     string  `gorm:"autoCreateTime"`
	UpdatedAt     string  `gorm:"autoUpdateTime"`

	// Relasi
	User    User    `gorm:"foreignKey:UserID;references:ID"`
	Package Package `gorm:"foreignKey:PackageID;references:ID"`
}
