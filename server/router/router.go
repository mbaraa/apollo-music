package router

import (
	"github.com/mbaraa/apollo-music/config"
	"github.com/mbaraa/apollo-music/controllers"

	"github.com/gofiber/fiber/v2"
)

var server *fiber.App = nil

func Start() {
	err := server.Listen(":" + config.PortNumber())
	if err != nil {
		panic(err)
	}
}

func init() {
	server = fiber.New(fiber.Config{
		AppName: "Apollo Music",
	})

	for _, controller := range controllers.GetControllers() {
		controller.Bind(server)
	}
}
