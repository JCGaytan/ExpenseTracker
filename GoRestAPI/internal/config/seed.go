package config

import (
	"expense-tracker/internal/models"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// SeedUsers inserts initial users for testing login
func SeedUsers() {
	users := []struct {
		Email    string
		Password string
	}{
		{"test1@example.com", "password123"},
		{"test2@example.com", "password456"},
	}
	for _, u := range users {
		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("failed to hash password for %s: %v", u.Email, err)
			continue
		}
		user := models.User{Email: u.Email, PasswordHash: string(hash)}
		if err := DB.Where(models.User{Email: u.Email}).FirstOrCreate(&user).Error; err != nil {
			log.Printf("failed to seed user %s: %v", u.Email, err)
		}
	}
}
