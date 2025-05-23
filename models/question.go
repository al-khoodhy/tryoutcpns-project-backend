package models

import (
	"gorm.io/gorm"
)

type Question struct {
	gorm.Model // otomatis dapet ID, CreatedAt, UpdatedAt, DeletedAt

	Title         string `gorm:"type:varchar(200);not null"`
	Content       string `gorm:"type:text;not null"`
	OptionA       string `gorm:"type:varchar(255);not null"`
	OptionB       string `gorm:"type:varchar(255);not null"`
	OptionC       string `gorm:"type:varchar(255);not null"`
	OptionD       string `gorm:"type:varchar(255);not null"`
	CorrectAnswer string `gorm:"type:char(1);check:correct_answer IN ('A','B','C','D');not null"`
	Subject       string `gorm:"type:varchar(10);not null;check:subject IN ('TWK','TIU','TKP')"`
	Difficulty    string `gorm:"type:varchar(10);not null;check:difficulty IN ('easy','medium','hard')"`
	Explanation   string `gorm:"type:text"`
	IsPublished   bool   `gorm:"default:false"`
}
