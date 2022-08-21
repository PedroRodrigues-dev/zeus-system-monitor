package configs

import (
	"zeus/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, _ := gorm.Open(sqlite.Open("cfg.sqlite3"), &gorm.Config{})
	db.AutoMigrate(&models.Config{})
	DB = db
}
