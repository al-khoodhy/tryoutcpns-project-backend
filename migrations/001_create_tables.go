package migrations

import (
	"tryoutcpns-project-backend/models"

	"github.com/jinzhu/gorm"
)

func CreateTables(db *gorm.DB) {
	db.AutoMigrate(
		&models.User{},
		&models.Package{},
		&models.Question{},
		&models.Result{},
		&models.Voucher{},
		&models.Transaction{},
		&models.Affiliate{},
		&models.Leaderboard{},
		&models.UserAnswer{},
		&models.Materi{},
		&models.Notification{},
		&models.AdminLog{},
		&models.UserPackage{},
		&models.UserVoucher{},
	)
}
