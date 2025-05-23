package models

type Package struct {
	ID           uint    `gorm:"primaryKey"`
	Name         string  `gorm:"type:varchar(100);not null"`
	Description  string  `gorm:"type:text"`
	Type         string  `gorm:"type:enum('free','premium','combo');not null"`
	Price        float64 `gorm:"type:decimal(10,2);default:0"`
	DurationDays int     `gorm:"type:int;default:0"`
	MaxTryouts   int     `gorm:"type:int;default:0"`
	IsAvailable  bool    `gorm:"default:true"`
	CreatedAt    string  `gorm:"autoCreateTime"`
	UpdatedAt    string  `gorm:"autoUpdateTime"`
}
