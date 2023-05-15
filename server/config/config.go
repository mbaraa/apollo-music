package config

import (
	"log"

	"github.com/mbaraa/apollo-music/data"
	"github.com/mbaraa/apollo-music/data/db"
	"github.com/mbaraa/apollo-music/entities"
	"github.com/mbaraa/apollo-music/models"
	"github.com/mbaraa/apollo-music/utils/jwt"
)

var (
	dbConnection     = db.GetDBConnector()
	userRepo         data.CRUDRepo[models.User]
	verificationRepo data.CRUDRepo[models.Verification]
	subscriptionRepo data.CRUDRepo[models.Subscription]
	storageRepo      data.CRUDRepo[models.Storage]
	musicRepo        data.CRUDRepo[models.Music]
	albumRepo        data.CRUDRepo[models.MusicAlbum]
	artistRepo       data.CRUDRepo[models.MusicArtist]
	yearRepo         data.CRUDRepo[models.MusicReleaseYear]
	genreRepo        data.CRUDRepo[models.MusicGenre]
)

var (
	jwtUtil jwt.Manager[entities.JSON]
)

func init() {
	db.InitTables()
	userRepo = db.NewBaseDB[models.User](dbConnection)
	verificationRepo = db.NewBaseDB[models.Verification](dbConnection)
	subscriptionRepo = db.NewBaseDB[models.Subscription](dbConnection)
	storageRepo = db.NewBaseDB[models.Storage](dbConnection)
	musicRepo = db.NewBaseDB[models.Music](dbConnection)
	albumRepo = db.NewBaseDB[models.MusicAlbum](dbConnection)
	artistRepo = db.NewBaseDB[models.MusicArtist](dbConnection)
	yearRepo = db.NewBaseDB[models.MusicReleaseYear](dbConnection)
	genreRepo = db.NewBaseDB[models.MusicGenre](dbConnection)
	jwtUtil = jwt.NewJWTImpl()

	log.SetFlags(log.LstdFlags | log.Llongfile | log.LUTC)
}

func UserRepo() data.CRUDRepo[models.User] {
	return userRepo
}

func VerificationRepo() data.CRUDRepo[models.Verification] {
	return verificationRepo
}

func SubscriptionRepo() data.CRUDRepo[models.Subscription] {
	return subscriptionRepo
}

func StorageRepo() data.CRUDRepo[models.Storage] {
	return storageRepo
}

func MusicRepo() data.CRUDRepo[models.Music] {
	return musicRepo
}

func AlbumRepo() data.CRUDRepo[models.MusicAlbum] {
	return albumRepo
}

func ArtistRepo() data.CRUDRepo[models.MusicArtist] {
	return artistRepo
}

func YearRepo() data.CRUDRepo[models.MusicReleaseYear] {
	return yearRepo
}

func GenreRepo() data.CRUDRepo[models.MusicGenre] {
	return genreRepo
}

func JWTUtil() jwt.Manager[entities.JSON] {
	return jwtUtil
}
