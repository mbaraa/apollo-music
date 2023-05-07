package models

import (
	"github.com/mbaraa/apollo-music/enums"
	"github.com/mbaraa/apollo-music/utils/strings"
	"gorm.io/gorm"
)

// User is the user's model
type User struct {
	gorm.Model `json:"-"`
	Id         uint             `gorm:"primaryKey;autoIncrement" json:"-"`
	PublicId   string           `gorm:"unique" json:"publicId"`
	FullName   string           `gorm:"not null" json:"fullName"`
	Email      string           `gorm:"unique" json:"email"`
	Password   string           `json:"-"`
	Status     enums.UserStatus `json:"status"`
}

func (u User) GetId() uint {
	return u.Id
}

func (u *User) BeforeCreate(_ *gorm.DB) error {
	u.PublicId = strings.GeneratePublicId()
	return nil
}
