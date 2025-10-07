package models

import (
	"time"

	"gorm.io/gorm"
)

// Member represents a church member
type Member struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	FirstName   string     `json:"first_name" gorm:"not null"`
	LastName    string     `json:"last_name" gorm:"not null"`
	Email       string     `json:"email" gorm:"uniqueIndex"`
	Phone       string     `json:"phone"`
	Address     string     `json:"address"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	JoinDate    time.Time  `json:"join_date" gorm:"default:CURRENT_TIMESTAMP"`
	IsActive    bool       `json:"is_active" gorm:"default:true"`
	Notes       string     `json:"notes"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// Event represents a church event
type Event struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description"`
	Date        time.Time `json:"date" gorm:"not null"`
	Time        string    `json:"time"`
	Location    string    `json:"location"`
	IsPublic    bool      `json:"is_public" gorm:"default:true"`
	CreatedBy   uint      `json:"created_by"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// ChurchSettings represents church configuration
type ChurchSettings struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	ChurchName   string    `json:"church_name" gorm:"not null"`
	ChurchAddress string   `json:"church_address"`
	ChurchPhone  string    `json:"church_phone"`
	ChurchEmail  string    `json:"church_email"`
	ServiceTimes string    `json:"service_times"` // JSON string for multiple service times
	UpdatedAt    time.Time `json:"updated_at"`
}

// TableName returns the table name for Member
func (Member) TableName() string {
	return "members"
}

// TableName returns the table name for Event
func (Event) TableName() string {
	return "events"
}

// TableName returns the table name for ChurchSettings
func (ChurchSettings) TableName() string {
	return "church_settings"
}
