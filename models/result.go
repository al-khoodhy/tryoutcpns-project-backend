package models

type Result struct {
	ID             uint   `gorm:"primaryKey"`
	UserID         uint   `gorm:"not null"`
	PackageID      *uint  `gorm:"null"`
	Score          int    `gorm:"not null"`
	TimeTaken      int    `gorm:"not null"`
	TotalQuestions int    `gorm:"not null"`
	CorrectAnswers int    `gorm:"not null"`
	WrongAnswers   int    `gorm:"not null"`
	CreatedAt      string `gorm:"autoCreateTime"`
	UpdatedAt      string `gorm:"autoUpdateTime"`

	// Relasi
	User    User    `gorm:"foreignKey:UserID;references:ID"`
	Package Package `gorm:"foreignKey:PackageID;references:ID"`
}
