package models

import (
	"time"
)

type Transaction struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Amount      float64   `gorm:"not null" json:"amount"`
	Date        time.Time `gorm:"not null" json:"date"`
	CategoryID  uint      `gorm:"not null" json:"category_id"`
	UserID      uint      `gorm:"not null" json:"user_id"`
	Description string    `json:"description"`
}
