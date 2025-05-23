package models

type UserAnswer struct {
	ID             uint   `gorm:"primaryKey"`
	UserID         uint   `gorm:"not null"`
	QuestionID     uint   `gorm:"not null"`
	SelectedAnswer string `gorm:"type:char(1);check:selected_answer IN ('A','B','C','D');not null"`
	IsCorrect      bool   `gorm:"not null"`
	CreatedAt      string `gorm:"autoCreateTime"`
	UpdatedAt      string `gorm:"autoUpdateTime"`

	// Relasi
	User     User     `gorm:"foreignKey:UserID;references:ID"`
	Question Question `gorm:"foreignKey:QuestionID;references:ID"`
}
