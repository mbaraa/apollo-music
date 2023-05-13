package models

import (
	"time"

	"github.com/mbaraa/apollo-music/enums"
	"gorm.io/gorm"
)

type Audio struct {
	gorm.Model  `json:"-"`
	Id          uint            `gorm:"primaryKey;autoIncrement" json:"-"`
	UserId      uint            `json:"-"`
	FileName    string          `json:"-"`
	FileSize    int64           `json:"-"`
	AccessTimes uint            `json:"-"`
	LastAccess  time.Time       `json:"-"`
	PublicPath  string          `json:"-"`
	Type        enums.AudioType `json:"-"`
}

func (a Audio) GetId() uint {
	return a.Id
}
