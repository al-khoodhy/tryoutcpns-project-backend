package models

import (
	"gorm.io/gorm"
)

type AdminLog struct {
	ID          uint   `gorm:"primaryKey"`
	AdminID     uint   `gorm:"not null"`
	Action      string `gorm:"type:varchar(255);not null"`
	Description string `gorm:"type:text"`
	IPAddress   string `gorm:"type:varchar(45)"`
	CreatedAt   string `gorm:"autoCreateTime"`
	UpdatedAt   string `gorm:"autoUpdateTime"`

	// Relasi
	Admin User `gorm:"foreignKey:AdminID;references:ID"`
}
