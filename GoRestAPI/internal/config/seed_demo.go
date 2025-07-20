package config

import (
	"expense-tracker/internal/models"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

// SeedDemoData inserts demo data for all tables
func SeedDemoData() {
	// Seed user
	hash, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	user := models.User{Email: "demo@example.com", PasswordHash: string(hash)}
	if err := DB.Where(models.User{Email: user.Email}).FirstOrCreate(&user).Error; err != nil {
		log.Printf("failed to seed user: %v", err)
	}

	// Seed category
	category := models.Category{Name: "Groceries", UserID: user.ID}
	if err := DB.Where(models.Category{Name: category.Name, UserID: user.ID}).FirstOrCreate(&category).Error; err != nil {
		log.Printf("failed to seed category: %v", err)
	}

	// Seed transaction
	date, err := time.Parse("2006-01-02", "2025-07-19")
	if err != nil {
		log.Printf("failed to parse date: %v", err)
		return
	}
	transaction := models.Transaction{
		Amount:      -50.0,
		Date:        date,
		CategoryID:  category.ID,
		UserID:      user.ID,
		Description: "Demo grocery shopping",
	}
	if err := DB.Where(models.Transaction{Description: transaction.Description, UserID: user.ID}).FirstOrCreate(&transaction).Error; err != nil {
		log.Printf("failed to seed transaction: %v", err)
	}
}
