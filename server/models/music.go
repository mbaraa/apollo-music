package models

import (
	"github.com/mbaraa/apollo-music/utils/strings"
	"gorm.io/gorm"
)

type Music struct {
	gorm.Model  `json:"-"`
	Id          uint   `gorm:"primaryKey;autoIncrement" json:"-"`
	UserId      uint   `json:"-"`
	PublicId    string `json:"-"`
	Title       string `json:"-"`
	Audio       Audio  `gorm:"foreignKey:AudioId" json:"-"`
	AudioId     uint   `json:"-"`
	AlbumTitle  string `json:"-"`
	AlbumId     uint   `json:"-"`
	ArtistName  string `json:"-"`
	ArtistId    uint   `json:"-"`
	Year        string `json:"-"`
	YearId      uint   `json:"-"`
	Genre       string `json:"-"`
	GenreId     uint   `json:"-"`
	TrackNumber int    `json:"-"`
}

func (m Music) GetId() uint {
	return m.Id
}

func (m *Music) BeforeCreate(_ *gorm.DB) error {
	m.PublicId = strings.GeneratePublicId()
	return nil
}

func (m *Music) AfterFind(db *gorm.DB) error {
	return db.Find(&m.Audio, "id = ?", m.AudioId).Error
}

type MusicAlbum struct {
	gorm.Model `json:"-"`
	Id         uint   `gorm:"primaryKey;autoIncrement" json:"-"`
	UserId     uint   `json:"-"`
	PublicId   string `json:"-"`
	Title      string `json:"-"`
	ArtistName string `json:"-"`
	ArtistId   uint   `json:"-"`
	Year       string `json:"-"`
	YearId     uint   `json:"-"`
	Genre      string `json:"-"`
	CoverB64   string `json:"-"`
	GenreId    uint   `json:"-"`
}

func (a MusicAlbum) GetId() uint {
	return a.Id
}

func (a *MusicAlbum) BeforeCreate(_ *gorm.DB) error {
	a.PublicId = strings.GeneratePublicId()
	return nil
}

type MusicArtist struct {
	gorm.Model `json:"-"`
	Id         uint   `gorm:"primaryKey;autoIncrement" json:"-"`
	UserId     uint   `json:"-"`
	PublicId   string `json:"-"`
	Name       string `json:"-"`
}

func (a MusicArtist) GetId() uint {
	return a.Id
}

func (a *MusicArtist) BeforeCreate(_ *gorm.DB) error {
	a.PublicId = strings.GeneratePublicId()
	return nil
}

type MusicReleaseYear struct {
	gorm.Model `json:"-"`
	Id         uint   `gorm:"primaryKey;autoIncrement" json:"-"`
	UserId     uint   `json:"-"`
	PublicId   string `json:"-"`
	Name       string `json:"-"`
}

func (y MusicReleaseYear) GetId() uint {
	return y.Id
}

func (y *MusicReleaseYear) BeforeCreate(_ *gorm.DB) error {
	y.PublicId = strings.GeneratePublicId()
	return nil
}

type MusicGenre struct {
	gorm.Model `json:"-"`
	Id         uint   `gorm:"primaryKey;autoIncrement" json:"-"`
	UserId     uint   `json:"-"`
	PublicId   string `json:"-"`
	Name       string `json:"-"`
}

func (g MusicGenre) GetId() uint {
	return g.Id
}

func (g *MusicGenre) BeforeCreate(_ *gorm.DB) error {
	g.PublicId = strings.GeneratePublicId()
	return nil
}
