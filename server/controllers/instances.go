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
	),
}

func GetControllers() []Bindable {
	return bindables
}
