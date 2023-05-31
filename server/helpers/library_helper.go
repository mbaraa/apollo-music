package helpers

import (
	"log"
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
		log.Println(err)
		return response.Build(errors.InvalidToken, nil)
	}

	dbUser, err := m.userRepo.GetByConds("email = ?", claims.Payload["email"])
	if err != nil {
		log.Println(err)
		return response.Build(errors.NotFound, nil)
	}

	dbMusic, err := m.musicRepo.GetByConds("user_id = ?", dbUser[0].Id)
	if err != nil {
		log.Println(err)
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
		log.Println(err)
		return response.Build(errors.InvalidToken, nil)
	}

	dbUser, err := m.userRepo.GetByConds("email = ?", claims.Payload["email"])
	if err != nil {
		log.Println(err)
		return response.Build(errors.NotFound, nil)
	}

	dbAlbum, err := m.albumRepo.GetByConds("public_id = ?", albumPublicId)
	if err != nil {
		log.Println(err)
		return response.Build(errors.NotFound, nil)
	}

	dbMusic, err := m.musicRepo.GetByConds("user_id = ? and album_id = ?", dbUser[0].Id, dbAlbum[0].Id)
	if err != nil {
		log.Println(err)
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
		CoverB64:   dbAlbum[0].CoverB64,
		Songs:      music,
	})
}

func (m *LibraryHelper) GetAlbums(token string, fetchCovers bool) (entities.JSON, int) {
	claims, err := m.jwtUtil.Decode(token, jwt.SessionToken)
	if err != nil {
		log.Println(err)
		return response.Build(errors.InvalidToken, nil)
	}

	dbUser, err := m.userRepo.GetByConds("email = ?", claims.Payload["email"])
	if err != nil {
		log.Println(err)
		return response.Build(errors.NotFound, nil)
	}

	dbAlbums, err := m.albumRepo.GetByConds("user_id = ?", dbUser[0].Id)
	if err != nil {
		log.Println(err)
		return response.Build(errors.NotFound, nil)
	}

	albums := make([]entities.MusicAlbum, 0)
	for _, album := range dbAlbums {
		eAlbum := entities.MusicAlbum{
			PublicId:   album.PublicId,
			Title:      album.Title,
			ArtistName: album.ArtistName,
			Year:       album.Year,
			Genre:      album.Genre,
		}
		if fetchCovers {
			eAlbum.CoverB64 = album.CoverB64
		}
		albums = append(albums, eAlbum)
	}

	sort.Slice(albums, func(i, j int) bool {
		return albums[i].Title < albums[j].Title
	})

	return response.Build(errors.None, albums)
}

func (m *LibraryHelper) GetArtist(token, artistPublicId string) (entities.JSON, int) {
	claims, err := m.jwtUtil.Decode(token, jwt.SessionToken)
	if err != nil {
		log.Println(err)
		return response.Build(errors.InvalidToken, nil)
	}

	dbUser, err := m.userRepo.GetByConds("email = ?", claims.Payload["email"])
	if err != nil {
		log.Println(err)
		return response.Build(errors.NotFound, nil)
	}

	dbArtist, err := m.artistRepo.GetByConds("public_id = ?", artistPublicId)
	if err != nil {
		log.Println(err)
		return response.Build(errors.NotFound, nil)
	}

	dbAlbums, err := m.albumRepo.GetByConds("user_id = ? and artist_id = ?", dbUser[0].Id, dbArtist[0].Id)
	if err != nil {
		log.Println(err)
		return response.Build(errors.NotFound, nil)
	}

	albums := make([]entities.MusicAlbum, 0)
	for _, album := range dbAlbums {
		albums = append(albums, entities.MusicAlbum{
			PublicId:   album.PublicId,
			Title:      album.Title,
			ArtistName: album.ArtistName,
			Year:       album.Year,
			CoverB64:   album.CoverB64,
			Genre:      album.Genre,
		})
	}

	sort.Slice(albums, func(i, j int) bool {
		return albums[i].Title < albums[j].Title
	})

	return response.Build(errors.None, albums)
}

func (m *LibraryHelper) GetArtists(token string) (entities.JSON, int) {
	claims, err := m.jwtUtil.Decode(token, jwt.SessionToken)
	if err != nil {
		log.Println(err)
		return response.Build(errors.InvalidToken, nil)
	}

	dbUser, err := m.userRepo.GetByConds("email = ?", claims.Payload["email"])
	if err != nil {
		log.Println(err)
		return response.Build(errors.NotFound, nil)
	}

	dbArtists, err := m.artistRepo.GetByConds("user_id = ?", dbUser[0].Id)
	if err != nil {
		log.Println(err)
		return response.Build(errors.NotFound, nil)
	}

	artists := make([]entities.MusicArtist, 0)
	for _, artist := range dbArtists {
		artists = append(artists, entities.MusicArtist{
			PublicId: artist.PublicId,
			Name:     artist.Name,
		})
	}

	sort.Slice(artists, func(i, j int) bool {
		return artists[i].Name < artists[j].Name
	})

	return response.Build(errors.None, artists)
}

func (m *LibraryHelper) GetYear(token, yearPublicId string) (entities.JSON, int) {
	claims, err := m.jwtUtil.Decode(token, jwt.SessionToken)
	if err != nil {
		log.Println(err)
		return response.Build(errors.InvalidToken, nil)
	}

	dbUser, err := m.userRepo.GetByConds("email = ?", claims.Payload["email"])
	if err != nil {
		return response.Build(errors.NotFound, nil)
	}

	dbYear, err := m.yearRepo.GetByConds("public_id = ?", yearPublicId)
	if err != nil {
		log.Println(err)
		return response.Build(errors.NotFound, nil)
	}

	dbMusic, err := m.musicRepo.GetByConds("user_id = ? and year_id = ?", dbUser[0].Id, dbYear[0].Id)
	if err != nil {
		log.Println(err)
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

	return response.Build(errors.None, entities.MusicReleaseYear{
		PublicId: dbYear[0].PublicId,
		Name:     dbYear[0].Name,
		Songs:    music,
	})
}

func (m *LibraryHelper) GetYears(token string) (entities.JSON, int) {
	claims, err := m.jwtUtil.Decode(token, jwt.SessionToken)
	if err != nil {
		log.Println(err)
		return response.Build(errors.InvalidToken, nil)
	}

	dbUser, err := m.userRepo.GetByConds("email = ?", claims.Payload["email"])
	if err != nil {
		log.Println(err)
		return response.Build(errors.NotFound, nil)
	}

	dbYears, err := m.yearRepo.GetByConds("user_id = ?", dbUser[0].Id)
	if err != nil {
		log.Println(err)
		return response.Build(errors.NotFound, nil)
	}

	years := make([]entities.MusicReleaseYear, 0)
	for _, year := range dbYears {
		years = append(years, entities.MusicReleaseYear{
			PublicId: year.PublicId,
			Name:     year.Name,
		})
	}

	sort.Slice(years, func(i, j int) bool {
		return years[i].Name < years[j].Name
	})

	return response.Build(errors.None, years)
}

func (m *LibraryHelper) GetGenre(token, genrePublicId string) (entities.JSON, int) {
	claims, err := m.jwtUtil.Decode(token, jwt.SessionToken)
	if err != nil {
		log.Println(err)
		return response.Build(errors.InvalidToken, nil)
	}

	dbUser, err := m.userRepo.GetByConds("email = ?", claims.Payload["email"])
	if err != nil {
		log.Println(err)
		return response.Build(errors.NotFound, nil)
	}

	dbGenre, err := m.genreRepo.GetByConds("public_id = ?", genrePublicId)
	if err != nil {
		log.Println(err)
		return response.Build(errors.NotFound, nil)
	}

	dbMusic, err := m.musicRepo.GetByConds("user_id = ? and genre_id = ?", dbUser[0].Id, dbGenre[0].Id)
	if err != nil {
		log.Println(err)
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

	return response.Build(errors.None, entities.MusicGenre{
		PublicId: dbGenre[0].PublicId,
		Name:     dbGenre[0].Name,
		Songs:    music,
	})
}

func (m *LibraryHelper) GetGenres(token string) (entities.JSON, int) {
	claims, err := m.jwtUtil.Decode(token, jwt.SessionToken)
	if err != nil {
		log.Println(err)
		return response.Build(errors.InvalidToken, nil)
	}

	dbUser, err := m.userRepo.GetByConds("email = ?", claims.Payload["email"])
	if err != nil {
		log.Println(err)
		return response.Build(errors.NotFound, nil)
	}

	dbGenres, err := m.genreRepo.GetByConds("user_id = ?", dbUser[0].Id)
	if err != nil {
		log.Println(err)
		return response.Build(errors.NotFound, nil)
	}

	genres := make([]entities.MusicGenre, 0)
	for _, genre := range dbGenres {
		genres = append(genres, entities.MusicGenre{
			PublicId: genre.PublicId,
			Name:     genre.Name,
		})
	}

	sort.Slice(genres, func(i, j int) bool {
		return genres[i].Name < genres[j].Name
	})

	return response.Build(errors.None, genres)
}
