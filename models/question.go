package models

import (
	"gorm.io/gorm"
)

type Question struct {
	ID            uint   `gorm:"primaryKey"`
	Title         string `gorm:"type:varchar(200);not null"`
	Content       string `gorm:"type:text;not null"`
	OptionA       string `gorm:"type:varchar(255);not null"`
	OptionB       string `gorm:"type:varchar(255);not null"`
	OptionC       string `gorm:"type:varchar(255);not null"`
	OptionD       string `gorm:"type:varchar(255);not null"`
	CorrectAnswer string `gorm:"type:char(1);check:correct_answer IN ('A','B','C','D');not null"`
	Subject       string `gorm:"type:enum('TWK','TIU','TKP');not null"`
	Difficulty    string `gorm:"type:enum('easy','medium','hard');not null"`
	Explanation   string `gorm:"type:text"`
	IsPublished   bool   `gorm:"default:false"`
	CreatedAt     string `gorm:"autoCreateTime"`
	UpdatedAt     string `gorm:"autoUpdateTime"`
}
