package helpers

import (
	goerrors "errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"strings"
	"time"

	"github.com/dhowden/tag"

	"github.com/mbaraa/apollo-music/config/env"
	"github.com/mbaraa/apollo-music/data"
	"github.com/mbaraa/apollo-music/entities"
	"github.com/mbaraa/apollo-music/enums"
	"github.com/mbaraa/apollo-music/errors"
	"github.com/mbaraa/apollo-music/helpers/response"
	"github.com/mbaraa/apollo-music/models"
	"github.com/mbaraa/apollo-music/utils/jwt"
)

type UploadHelper struct {
	storageRepo data.CRUDRepo[models.Storage]
	userRepo    data.CRUDRepo[models.User]
	musicRepo   data.CRUDRepo[models.Music]
	albumRepo   data.CRUDRepo[models.MusicAlbum]
	artistRepo  data.CRUDRepo[models.MusicArtist]
	yearRepo    data.CRUDRepo[models.MusicReleaseYear]
	genreRepo   data.CRUDRepo[models.MusicGenre]
	jwtUtil     jwt.Manager[entities.JSON]
}

func NewUploadHelper(
	storageRepo data.CRUDRepo[models.Storage],
	userRepo data.CRUDRepo[models.User],
	musicRepo data.CRUDRepo[models.Music],
	albumRepo data.CRUDRepo[models.MusicAlbum],
	artistRepo data.CRUDRepo[models.MusicArtist],
	yearRepo data.CRUDRepo[models.MusicReleaseYear],
	genreRepo data.CRUDRepo[models.MusicGenre],
	jwtUtil jwt.Manager[entities.JSON],
) *UploadHelper {
	return &UploadHelper{
		storageRepo: storageRepo,
		userRepo:    userRepo,
		musicRepo:   musicRepo,
		albumRepo:   albumRepo,
		artistRepo:  artistRepo,
		yearRepo:    yearRepo,
		genreRepo:   genreRepo,
		jwtUtil:     jwtUtil,
	}
}

// TODO
// handle duplicate files
func (u *UploadHelper) UploadFile(token string, audioType enums.AudioType, fileHeader *multipart.FileHeader) (entities.JSON, int) {
	claims, err := u.jwtUtil.Decode(token, jwt.SessionToken)
	if err != nil {
		log.Println(err)
		return response.Build(errors.InvalidToken, nil)
	}

	if !strings.Contains(fileHeader.Header.Get("Content-Type"), "audio") {
		return response.Build(errors.BadRequest, nil)
	}

	dbUser, err := u.userRepo.GetByConds("email = ?", claims.Payload["email"])
	if err != nil {
		log.Println(err)
		return response.Build(errors.NotFound, nil)
	}

	dbStorage, err := u.storageRepo.GetByConds("user_id = ?", dbUser[0].Id)
	if err != nil {
		log.Println(err)
		return response.Build(errors.NotFound, nil)
	}

	if dbStorage[0].Size-dbStorage[0].Used < fileHeader.Size/(1024*1024) {
		return response.Build(errors.InsufficientStorage, nil)
	}

	saveDirectory := fmt.Sprintf("%s/%s/%s", env.MusicDirectory(), dbUser[0].PublicId, audioType.String())
	err = os.Mkdir(saveDirectory, 0755)
	if err != nil && !goerrors.Is(err, os.ErrExist) {
		return response.Build(errors.InternalServerError, nil)
	}

	file, err := fileHeader.Open()
	if err != nil {
		log.Println(err)
		return response.Build(errors.BadRequest, nil)
	}
	defer file.Close()

	filePath := fmt.Sprintf("%s/%s", saveDirectory, fileHeader.Filename)
	fileToSave, err := os.Create(filePath)

	_, err = io.Copy(fileToSave, file)
	if err != nil {
		log.Println(err)
		return response.Build(errors.InternalServerError, nil)
	}

	err = u.storageRepo.Update(&models.Storage{
		Used: dbStorage[0].Used + fileHeader.Size/(1024*1024),
	}, "id = ?", dbStorage[0].Id)
	if err != nil {
		log.Println(err)
		return response.Build(errors.InternalServerError, nil)
	}

	if audioType == enums.MusicType {
		_, _ = file.Seek(0, 0)
		musicMetaData, err := tag.ReadFrom(file)
		if err != nil {
			log.Println(err)
			musicMetaData = &defaultMusicValues{fileHeader.Filename[:strings.Index(fileHeader.Filename, ".")]}
			return response.Build(errors.InternalServerError, nil)
		}

		title := musicMetaData.Title()
		album := musicMetaData.Album()
		year := musicMetaData.Year()
		genre := musicMetaData.Genre()
		artist := musicMetaData.Artist()
		if len(artist) == 0 {
			artist = musicMetaData.AlbumArtist()
		}
		trackNumber, _ := musicMetaData.Track()

		_dbArtist := models.MusicArtist{
			Name:   artist,
			UserId: dbUser[0].Id,
		}
		_dbYear := models.MusicReleaseYear{
			Name:   fmt.Sprint(year),
			UserId: dbUser[0].Id,
		}
		_dbGenre := models.MusicGenre{
			Name:   genre,
			UserId: dbUser[0].Id,
		}

		dbArtist, err := u.artistRepo.GetByConds("name = ?", artist)
		if err != nil {
			err = u.artistRepo.Add(&_dbArtist)
			if err != nil {
				log.Println(err)
				return response.Build(errors.InternalServerError, nil)
			}
		} else {
			_dbArtist = dbArtist[0]
		}

		dbYear, err := u.yearRepo.GetByConds("name = ?", year)
		if err != nil {
			err = u.yearRepo.Add(&_dbYear)
			if err != nil {
				log.Println(err)
				return response.Build(errors.InternalServerError, nil)
			}
		} else {
			_dbYear = dbYear[0]
		}

		dbGenre, err := u.genreRepo.GetByConds("name = ?", genre)
		if err != nil {
			err = u.genreRepo.Add(&_dbGenre)
			if err != nil {
				log.Println(err)
				return response.Build(errors.InternalServerError, nil)
			}
		} else {
			_dbGenre = dbGenre[0]
		}

		_dbAlbum := models.MusicAlbum{
			UserId:     dbUser[0].Id,
			Title:      album,
			ArtistName: artist,
			ArtistId:   _dbArtist.Id,
			Year:       fmt.Sprint(year),
			YearId:     _dbYear.Id,
			Genre:      genre,
			GenreId:    _dbGenre.Id,
		}
		dbAlbum, err := u.albumRepo.GetByConds("title = ?", album)
		if err != nil {
			err = u.albumRepo.Add(&_dbAlbum)
			if err != nil {
				log.Println(err)
				return response.Build(errors.InternalServerError, nil)
			}
		} else {
			_dbAlbum = dbAlbum[0]
		}

		_dbMusic := models.Music{
			UserId:      dbUser[0].Id,
			Title:       title,
			AlbumTitle:  album,
			AlbumId:     _dbAlbum.Id,
			ArtistName:  artist,
			ArtistId:    _dbArtist.Id,
			Year:        fmt.Sprint(year),
			YearId:      _dbYear.Id,
			Genre:       genre,
			GenreId:     _dbGenre.Id,
			TrackNumber: trackNumber,
			Audio: models.Audio{
				FileName:    fileHeader.Filename,
				FileSize:    fileHeader.Size / (1024 * 1024),
				AccessTimes: 0,
				LastAccess:  time.Now(),
				PublicPath:  fmt.Sprintf("%s/%s/%s", dbUser[0].PublicId, audioType.String(), fileHeader.Filename),
				Type:        audioType,
			},
		}
		dbMusic, err := u.musicRepo.GetByConds("title = ?", title)
		if err != nil {
			err = u.musicRepo.Add(&_dbMusic)
			if err != nil {
				log.Println(err)
				return response.Build(errors.InternalServerError, nil)
			}
		} else {
			_dbMusic = dbMusic[0]
		}
	}

	return response.Build(errors.None, nil)
}

// defaultMusicValues implements tag.Metadata for default music metadata values
type defaultMusicValues struct {
	title string
}

func (d *defaultMusicValues) Format() tag.Format { return tag.Format("") }

func (d *defaultMusicValues) FileType() tag.FileType { return tag.FileType("") }

func (d *defaultMusicValues) Title() string { return d.title }

func (d *defaultMusicValues) Album() string { return "Unknown Album" }

func (d *defaultMusicValues) Artist() string { return "Unknown Artist" }

func (d *defaultMusicValues) AlbumArtist() string { return "Unknown Artist" }

func (d *defaultMusicValues) Composer() string { return "Unknown Composer" }

func (d *defaultMusicValues) Year() int { return 0 }

func (d *defaultMusicValues) Genre() string { return "Unknown Genre" }

func (d *defaultMusicValues) Track() (int, int) { return 0, 0 }

func (d *defaultMusicValues) Disc() (int, int) { return 0, 0 }

func (d *defaultMusicValues) Picture() *tag.Picture { return nil }

func (d *defaultMusicValues) Lyrics() string { return "" }

func (d *defaultMusicValues) Comment() string { return "" }

func (d *defaultMusicValues) Raw() map[string]interface{} { return nil }
