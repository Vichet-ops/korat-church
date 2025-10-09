package models

import (
	"time"
)

// ContactMessage represents a contact form submission
type ContactMessage struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null" binding:"required,min=2,max=100"`
	Email     string    `json:"email" gorm:"not null" binding:"required,email"`
	Subject   string    `json:"subject" gorm:"not null" binding:"required,min=3,max=200"`
	Message   string    `json:"message" gorm:"type:text;not null" binding:"required,min=10"`
	Status    string    `json:"status" gorm:"default:'new'"` // new, read, replied
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName returns the table name for ContactMessage
func (ContactMessage) TableName() string {
	return "contact_messages"
}

