package models

import (
	"gorm.io/gorm"
)

type Materi struct {
	ID          uint   `gorm:"primaryKey"`
	Title       string `gorm:"type:varchar(255);not null"`
	Content     string `gorm:"type:text;not null"`
	Type        string `gorm:"type:enum('article','video','ebook');not null"`
	FileURL     string `gorm:"type:varchar(255)"`
	Subject     string `gorm:"type:enum('TWK','TIU','TKP');not null"`
	IsPublished bool   `gorm:"default:false"`
	CreatedAt   string `gorm:"autoCreateTime"`
	UpdatedAt   string `gorm:"autoUpdateTime"`
}
