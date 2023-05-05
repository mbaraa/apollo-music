package controllers

import (
	"github.com/mbaraa/apollo-music/controllers/apis"
	"github.com/mbaraa/apollo-music/controllers/sockets"
)

var bindables []Bindable = []Bindable{
	apis.NewLsMusicApi(),
	apis.NewExampleApi(),
	sockets.NewEchoSocket(),
}

func GetControllers() []Bindable {
	return bindables
}
