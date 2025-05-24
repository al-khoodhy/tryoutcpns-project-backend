package models

import "time"

type User struct {
	ID         uint      `gorm:"primaryKey"`
	Name       string    `gorm:"type:varchar(100);not null"`
	Email      string    `gorm:"type:varchar(255);unique;not null"`
	Phone      string    `gorm:"type:varchar(15);unique;not null"`
	Password   string    `gorm:"type:varchar(255);not null"`
	Role       string    `gorm:"type:enum('admin','affiliate','user');default:'user';not null"`
	IsActive   bool      `gorm:"default:true"`
	ReferralID string    `gorm:"type:varchar(36);not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
