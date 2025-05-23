package models

import (
	"time"
)

type Voucher struct {
	ID           uint      `gorm:"primaryKey"`
	Code         string    `gorm:"type:varchar(50);unique;not null"`
	DiscountType string    `gorm:"type:varchar(20);not null;check:discount_type IN ('percentage','fixed')"`
	Value        float64   `gorm:"not null"`
	ValidFrom    time.Time `gorm:"not null"`
	ValidUntil   time.Time `gorm:"not null"`
	UsageLimit   int       `gorm:"default:1"`
	UsedCount    int       `gorm:"default:0"`
	IsGlobal     bool      `gorm:"default:false"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
