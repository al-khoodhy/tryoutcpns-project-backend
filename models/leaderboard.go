package models

type Leaderboard struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	Score     int    `gorm:"not null"`
	PackageID *uint  `gorm:"null"`
	CreatedAt string `gorm:"autoCreateTime"`
	UpdatedAt string `gorm:"autoUpdateTime"`

	// Relasi
	User    User    `gorm:"foreignKey:UserID;references:ID"`
	Package Package `gorm:"foreignKey:PackageID;references:ID"`
}
