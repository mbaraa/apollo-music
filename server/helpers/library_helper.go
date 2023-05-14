package helpers

import (
	"sort"

	"github.com/mbaraa/apollo-music/data"
	"github.com/mbaraa/apollo-music/entities"
	"github.com/mbaraa/apollo-music/errors"
	"github.com/mbaraa/apollo-music/helpers/response"
	"github.com/mbaraa/apollo-music/models"
	"github.com/mbaraa/apollo-music/utils/jwt"
)

type LibraryHelper struct {
	musicRepo  data.CRUDRepo[models.Music]
	albumRepo  data.CRUDRepo[models.MusicAlbum]
	artistRepo data.CRUDRepo[models.MusicArtist]
	yearRepo   data.CRUDRepo[models.MusicReleaseYear]
	genreRepo  data.CRUDRepo[models.MusicGenre]
	userRepo   data.CRUDRepo[models.User]
	jwtUtil    jwt.Decoder[entities.JSON]
}

func NewLibraryHelper(
	musicRepo data.CRUDRepo[models.Music],
	albumRepo data.CRUDRepo[models.MusicAlbum],
	artistRepo data.CRUDRepo[models.MusicArtist],
	yearRepo data.CRUDRepo[models.MusicReleaseYear],
	genreRepo data.CRUDRepo[models.MusicGenre],
	userRepo data.CRUDRepo[models.User],
	jwtUtil jwt.Decoder[entities.JSON],
) *LibraryHelper {
	return &LibraryHelper{
		musicRepo:  musicRepo,
		albumRepo:  albumRepo,
		artistRepo: artistRepo,
		yearRepo:   yearRepo,
		genreRepo:  genreRepo,
		userRepo:   userRepo,
		jwtUtil:    jwtUtil,
	}
}

func (m *LibraryHelper) GetMusic(token string) (entities.JSON, int) {
	claims, err := m.jwtUtil.Decode(token, jwt.SessionToken)
	if err != nil {
		return response.Build(errors.InvalidToken, nil)
	}

	dbUser, err := m.userRepo.GetByConds("email = ?", claims.Payload["email"])
	if err != nil {
		return response.Build(errors.NotFound, nil)
	}

	dbMusic, err := m.musicRepo.GetByConds("user_id = ?", dbUser[0].Id)
	if err != nil {
		return response.Build(errors.NotFound, nil)
	}

	music := make([]entities.Music, 0)
	for _, m := range dbMusic {
		music = append(music, entities.Music{
			PublicId:    m.PublicId,
			Title:       m.Title,
			AlbumTitle:  m.AlbumTitle,
			ArtistName:  m.ArtistName,
			Year:        m.Year,
			Genre:       m.Genre,
			TrackNumber: m.TrackNumber,
			Audio: entities.Audio{
				FileName:    m.Audio.FileName,
				FileSize:    float64(m.Audio.FileSize),
				LastAccess:  m.Audio.LastAccess,
				AccessTimes: m.Audio.AccessTimes,
				PublicPath:  m.Audio.PublicPath,
				Type:        m.Audio.Type,
			},
		})
	}

	sort.Slice(music, func(i, j int) bool {
		return music[i].Title < music[j].Title
	})

	return response.Build(errors.None, music)
}

func (m *LibraryHelper) GetAlbum(token, albumPublicId string) (entities.JSON, int) {
	claims, err := m.jwtUtil.Decode(token, jwt.SessionToken)
	if err != nil {
		return response.Build(errors.InvalidToken, nil)
	}

	dbUser, err := m.userRepo.GetByConds("email = ?", claims.Payload["email"])
	if err != nil {
		return response.Build(errors.NotFound, nil)
	}

	dbAlbum, err := m.albumRepo.GetByConds("public_id = ?", albumPublicId)
	if err != nil {
		return response.Build(errors.NotFound, nil)
	}

	dbMusic, err := m.musicRepo.GetByConds("user_id = ? and album_id = ?", dbUser[0].Id, dbAlbum[0].Id)
	if err != nil {
		return response.Build(errors.NotFound, nil)
	}

	music := make([]entities.Music, 0)
	for _, m := range dbMusic {
		music = append(music, entities.Music{
			PublicId:    m.PublicId,
			Title:       m.Title,
			AlbumTitle:  m.AlbumTitle,
			ArtistName:  m.ArtistName,
			Year:        m.Year,
			Genre:       m.Genre,
			TrackNumber: m.TrackNumber,
			Audio: entities.Audio{
				FileName:    m.Audio.FileName,
				FileSize:    float64(m.Audio.FileSize),
				LastAccess:  m.Audio.LastAccess,
				AccessTimes: m.Audio.AccessTimes,
				PublicPath:  m.Audio.PublicPath,
				Type:        m.Audio.Type,
			},
		})
	}

	sort.Slice(music, func(i, j int) bool {
		return music[i].TrackNumber < music[j].TrackNumber
	})

	return response.Build(errors.None, entities.MusicAlbum{
		PublicId:   dbAlbum[0].PublicId,
		Title:      dbAlbum[0].Title,
		ArtistName: dbAlbum[0].ArtistName,
		Year:       dbAlbum[0].Year,
		Genre:      dbAlbum[0].Genre,
		Songs:      music,
	})
}

func (m *LibraryHelper) GetAlbums(token string) (entities.JSON, int) {
	claims, err := m.jwtUtil.Decode(token, jwt.SessionToken)
	if err != nil {
		return response.Build(errors.InvalidToken, nil)
	}

	dbUser, err := m.userRepo.GetByConds("email = ?", claims.Payload["email"])
	if err != nil {
		return response.Build(errors.NotFound, nil)
	}

	dbAlbums, err := m.albumRepo.GetByConds("user_id = ?", dbUser[0].Id)
	if err != nil {
		return response.Build(errors.NotFound, nil)
	}

	albums := make([]entities.MusicAlbum, 0)
	for _, album := range dbAlbums {
		albums = append(albums, entities.MusicAlbum{
			PublicId:   album.PublicId,
			Title:      album.Title,
			ArtistName: album.ArtistName,
			Year:       album.Year,
			Genre:      album.Genre,
		})
	}

	sort.Slice(albums, func(i, j int) bool {
		return albums[i].Title < albums[j].Title
	})

	return response.Build(errors.None, albums)
}
