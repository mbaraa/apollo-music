package router

import (
	"log"

	"github.com/mbaraa/apollo-music/config/env"
	"github.com/mbaraa/apollo-music/controllers"

	"github.com/gofiber/fiber/v2"
)

var server *fiber.App = nil

func Start() {
	err := server.Listen(":" + env.PortNumber())
	if err != nil {
		panic(err)
	}
}

func init() {
	server = fiber.New(fiber.Config{
		AppName:   "Apollo Music",
		BodyLimit: env.MaxSingleFileSize(),
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			log.Println("---------------- http error ----------------")
			log.Println(ctx.BaseURL())
			log.Println(err)
			return fiber.DefaultErrorHandler(ctx, err)
		},
		ReadBufferSize:    4096 * 1024,
		StreamRequestBody: true,
	})

	for _, controller := range controllers.GetControllers() {
		controller.Bind(server)
	}
}
