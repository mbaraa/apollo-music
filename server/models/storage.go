package models

import "gorm.io/gorm"

type Storage struct {
	gorm.Model `json:"-"`
	Id         uint  `gorm:"primaryKey;autoIncrement" json:"-"`
	UserId     uint  `json:"-"`
	Size       int64 `json:"-"`
	Used       int64 `json:"-"`
}

func (s Storage) GetId() uint {
	return s.Id
}
