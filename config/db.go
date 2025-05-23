package config

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	"root", // Ganti dengan username MySQL Anda
	// 	"",     // Ganti dengan password MySQL Anda
	// 	"localhost",
	// 	"3306",
	// 	"cpns_tryout",
	// )

	// var err error
	// DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal("Failed to connect to database:", err)
	// }

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate
	DB.AutoMigrate(
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

	// fmt.Println("✅ Database migrated successfully.")

	fmt.Println("✅ Successfully connected to the database")
}
