package entities

import (
	"time"

	"github.com/mbaraa/apollo-music/enums"
)

type Audio struct {
	FileName    string          `json:"fileName"`
	FileSize    float64         `json:"fileSize"`
	LastAccess  time.Time       `json:"lastAccess"`
	AccessTimes uint            `json:"accessTimes"`
	PublicPath  string          `json:"publicPath"`
	Type        enums.AudioType `json:"audioType"`
}
