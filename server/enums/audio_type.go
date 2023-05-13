package enums

import "strings"

type AudioType string

const (
	MusicType AudioType = "MUSIC"
)

func (a AudioType) String() string {
	return strings.ToLower(string(a))
}

func GetAudioType(typeText string) AudioType {
	switch strings.ToUpper(typeText) {
	case string(MusicType):
		fallthrough
	default:
		return MusicType
	}
}
