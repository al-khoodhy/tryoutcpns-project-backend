package models

type Affiliate struct {
	ID              uint    `gorm:"primaryKey"`
	UserID          uint    `gorm:"not null"`
	ReferralCode    string  `gorm:"type:varchar(36);unique;not null"`
	CommissionRate  float64 `gorm:"type:decimal(5,2);default:10.0"`
	TotalReferrals  int     `gorm:"default:0"`
	TotalCommission float64 `gorm:"default:0.0"`
	CreatedAt       string  `gorm:"autoCreateTime"`
	UpdatedAt       string  `gorm:"autoUpdateTime"`

	// Relasi
	User User `gorm:"foreignKey:UserID;references:ID"`
}
