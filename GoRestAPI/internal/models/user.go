
package models

type User struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	Email        string `gorm:"unique;not null" json:"email"`
	PasswordHash string `gorm:"not null" json:"-"`
	Categories   []Category
	Transactions []Transaction
}
