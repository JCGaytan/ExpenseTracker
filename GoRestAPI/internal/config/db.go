package config

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"expense-tracker/internal/models"
)

var DB *gorm.DB

func InitDB() {
	db, err := gorm.Open(sqlite.Open("expense_tracker.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}
	// Auto-migrate models
	db.AutoMigrate(&models.User{}, &models.Category{}, &models.Transaction{})
	DB = db
}
