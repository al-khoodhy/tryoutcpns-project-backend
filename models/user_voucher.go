package models

import (
	"gorm.io/gorm"
)

type UserVoucher struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	VoucherID uint   `gorm:"not null"`
	UsedAt    string `gorm:"type:datetime;default:CURRENT_TIMESTAMP"`
	CreatedAt string `gorm:"autoCreateTime"`
	UpdatedAt string `gorm:"autoUpdateTime"`

	// Relasi
	User    User    `gorm:"foreignKey:UserID;references:ID"`
	Voucher Voucher `gorm:"foreignKey:VoucherID;references:ID"`
}
