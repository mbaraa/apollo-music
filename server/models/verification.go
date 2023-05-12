package models

import (
	"time"

	"gorm.io/gorm"
)

type Verification struct {
	gorm.Model  `json:"-"`
	Id          uint      `gorm:"primaryKey;autoIncrement" json:"-"`
	UserId      uint      `json:"-"`
	OTP         string    `json:"-"`
	ValidBefore time.Time `json:"-"`
}

func (v Verification) GetId() uint {
	return v.Id
}

func (v *Verification) AfterDelete(_ *gorm.DB) error {
	v.OTP = ""
	return nil
}
