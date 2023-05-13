package entities

type Music struct {
	PublicId   string `json:"publicId"`
	Title      string `json:"title"`
	Audio      Audio  `json:"audio"`
	AlbumTitle string `json:"albumTitle"`
	ArtistName string `json:"artistName"`
	Year       string `json:"year"`
	Genre      string `json:"genre"`
}

type MusicAlbum struct {
	PublicId   string `json:"publicId"`
	Title      string `json:"title"`
	ArtistName string `json:"artistName"`
	Year       string `json:"year"`
	Genre      string `json:"genre"`
}

type MusicArtist struct {
	PublicId string `json:"publicId"`
	Name     string `json:"name"`
}

type MusicReleaseYear struct {
	PublicId string `json:"publicId"`
	Name     string `json:"name"`
}

type MusicGenre struct {
	PublicId string `json:"publicId"`
	Name     string `json:"name"`
}
