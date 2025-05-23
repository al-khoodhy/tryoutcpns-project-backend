package models

import (
	"gorm.io/gorm"
)

type Voucher struct {
	ID           uint    `gorm:"primaryKey"`
	Code         string  `gorm:"type:varchar(50);unique;not null"`
	DiscountType string  `gorm:"type:enum('percentage','fixed');not null"`
	Value        float64 `gorm:"type:decimal(10,2);not null"`
	ValidFrom    string  `gorm:"type:datetime;not null"`
	ValidUntil   string  `gorm:"type:datetime;not null"`
	UsageLimit   int     `gorm:"type:int;default:1"`
	UsedCount    int     `gorm:"type:int;default:0"`
	IsGlobal     bool    `gorm:"default:false"`
	CreatedAt    string  `gorm:"autoCreateTime"`
	UpdatedAt    string  `gorm:"autoUpdateTime"`
}
