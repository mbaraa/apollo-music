package models

import "gorm.io/gorm"

// User is the user's model
type User struct {
	gorm.Model
	Id       uint   `gorm:"primaryKey;autoIncrement"`
	PublicId string `gorm:"unique"`
	FullName string `gorm:"not null"`
	Email    string `gorm:"unique"`
	Password string
	IsOAuth  bool
	Status   UserStatus
}

// UserStatus represents the user's current subscription status
// where an active user is a user with a free subscription or an ongoing paid subscription
// and an inactive user is a user with an overdue subscription by 10 days
type UserStatus string

const (
	ActiveStatus   UserStatus = "ACTIVE"
	InactiveStatus UserStatus = "INACTIVE"
)
