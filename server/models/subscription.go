package models

import (
	"time"

	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model           `json:"-"`
	Id                   uint      `gorm:"primaryKey;autoIncrement" json:"-"`
	UserId               uint      `json:"-"`
	Size                 int64     `json:"-"`
	Price                int64     `json:"-"`
	ValidBefore          time.Time `json:"-"`
	StripeSubscriptionId string    `json:"-"`
	StripeCustomerId     string    `json:"-"`
}

func (s Subscription) GetId() uint {
	return s.Id
}
