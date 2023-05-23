package entities

type Music struct {
	PublicId    string `json:"publicId"`
	Title       string `json:"title"`
	Audio       Audio  `json:"audio"`
	AlbumTitle  string `json:"albumTitle"`
	ArtistName  string `json:"artistName"`
	Year        string `json:"year"`
	Genre       string `json:"genre"`
	TrackNumber int    `json:"trackNumber"`
}

type MusicAlbum struct {
	PublicId   string  `json:"publicId"`
	Title      string  `json:"title"`
	ArtistName string  `json:"artistName"`
	Year       string  `json:"year"`
	Genre      string  `json:"genre"`
	CoverB64   string  `json:"coverB64"`
	Songs      []Music `json:"songs"`
}

type MusicArtist struct {
	PublicId string  `json:"publicId"`
	Name     string  `json:"name"`
	Songs    []Music `json:"songs"`
}

type MusicReleaseYear struct {
	PublicId string  `json:"publicId"`
	Name     string  `json:"name"`
	Songs    []Music `json:"songs"`
}

type MusicGenre struct {
	PublicId string  `json:"publicId"`
	Name     string  `json:"name"`
	Songs    []Music `json:"songs"`
}
