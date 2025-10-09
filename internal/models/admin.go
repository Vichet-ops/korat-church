package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Admin represents an admin user
type Admin struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username" gorm:"uniqueIndex;not null" binding:"required,min=3,max=50"`
	Email     string    `json:"email" gorm:"uniqueIndex;not null" binding:"required,email"`
	Password  string    `json:"-" gorm:"not null"` // Never send password in JSON
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	IsActive  bool      `json:"is_active" gorm:"default:true"`
	LastLogin *time.Time `json:"last_login"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName returns the table name for Admin
func (Admin) TableName() string {
	return "admins"
}

// HashPassword hashes the password using bcrypt
func (a *Admin) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	a.Password = string(hashedPassword)
	return nil
}

// CheckPassword checks if the provided password is correct
func (a *Admin) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(password))
	return err == nil
}

