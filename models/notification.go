package models

type Notification struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	Message   string `gorm:"type:text;not null"`
	Read      bool   `gorm:"default:false"`
	CreatedAt string `gorm:"autoCreateTime"`
	UpdatedAt string `gorm:"autoUpdateTime"`

	// Relasi
	User User `gorm:"foreignKey:UserID;references:ID"`
}
