package models

import (
	"gorm.io/gorm"
)

type UserPackage struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	PackageID uint   `gorm:"not null"`
	StartDate string `gorm:"type:datetime;not null"`
	EndDate   string `gorm:"type:datetime;not null"`
	IsActive  bool   `gorm:"default:true"`
	CreatedAt string `gorm:"autoCreateTime"`
	UpdatedAt string `gorm:"autoUpdateTime"`

	// Relasi
	User    User    `gorm:"foreignKey:UserID;references:ID"`
	Package Package `gorm:"foreignKey:PackageID;references:ID"`
}
