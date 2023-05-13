package controllers

import (
	"github.com/mbaraa/apollo-music/config"
	"github.com/mbaraa/apollo-music/controllers/apis"
	"github.com/mbaraa/apollo-music/helpers"
	"github.com/mbaraa/apollo-music/helpers/auth"
)

var bindables []Bindable = []Bindable{
	apis.NewAuthApi(
		auth.NewEmailHelper(config.UserRepo(), config.VerificationRepo(), config.JWTUtil()),
		auth.NewOTPHelper(config.VerificationRepo(), config.UserRepo(), config.JWTUtil()),
		auth.NewPasswordResetHelper(config.UserRepo(), config.JWTUtil()),
		auth.NewSessionHelper(config.UserRepo(), config.JWTUtil()),
	),
	apis.NewSubscriptionApi(
		helpers.NewSubscriptionHelper(config.SubscriptionRepo(), config.UserRepo(),
			helpers.NewStorageHelper(config.StorageRepo(), config.UserRepo(), config.JWTUtil()),
			config.JWTUtil()),
	),
	apis.NewStorageApi(
		helpers.NewStorageHelper(config.StorageRepo(), config.UserRepo(), config.JWTUtil()),
	),
	apis.NewUploadApi(
		helpers.NewUploadHelper(config.StorageRepo(), config.UserRepo(), config.MusicRepo(),
			config.AlbumRepo(), config.ArtistRepo(), config.YearRepo(), config.GenreRepo(),
			config.JWTUtil()),
	),
	apis.NewLibraryApi(
		helpers.NewLibraryHelper(config.MusicRepo(), config.UserRepo(), config.JWTUtil()),
	),
}

func GetControllers() []Bindable {
	return bindables
}
