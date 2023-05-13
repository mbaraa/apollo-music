package controllers

import (
	"github.com/mbaraa/apollo-music/config"
	"github.com/mbaraa/apollo-music/controllers/apis"
	"github.com/mbaraa/apollo-music/helpers"
)

var bindables []Bindable = []Bindable{
	apis.NewAuthApi(
		helpers.NewEmailHelper(config.UserRepo(), config.VerificationRepo(), config.JWTUtil()),
		helpers.NewOTPHelper(config.VerificationRepo(), config.UserRepo(), config.JWTUtil()),
		helpers.NewPasswordResetHelper(config.UserRepo(), config.JWTUtil()),
		helpers.NewSessionHelper(config.UserRepo(), config.JWTUtil()),
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
		helpers.NewUploadHelper(config.StorageRepo(), config.UserRepo(), config.JWTUtil()),
	),
}

func GetControllers() []Bindable {
	return bindables
}
