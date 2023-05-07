package apis

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mbaraa/apollo-music/config/env"
	"github.com/mbaraa/apollo-music/middlewares"
)

type LsMusicApi struct{}

func NewLsMusicApi() *LsMusicApi {
	return &LsMusicApi{}
}

func (e *LsMusicApi) Bind(app *fiber.App) {
	app.Use(middlewares.AllowCors)
	app.Static("/static", env.MusicDirectory(), fiber.Static{
		Next: func(ctx *fiber.Ctx) bool {
			name := ctx.Query("name")
			fmt.Println("name: ", name)
			if name == "hello" {
				return false
			}
			return true
		}})

	app.Static("/music-static", env.MusicDirectory())
	lsMusicGroup := app.Group("/ls-music")
	lsMusicGroup.Use(middlewares.AllowJSON)
	lsMusicGroup.Use(middlewares.AllowCors)
	lsMusicGroup.Get("/", e.handleLsMusic)
}

func (e *LsMusicApi) handleLsMusic(ctx *fiber.Ctx) error {

	paths := make([]string, 0)

	err := filepath.Walk(env.MusicDirectory(), func(path string, info os.FileInfo, err error) error {
		// fmt.Println(path)
		// fmt.Println(info.Name())
		// fmt.Println(info)

		basename := filepath.Base(env.MusicDirectory())
		removeIndex := strings.Index(path, basename)
		paths = append(paths, path[removeIndex+len(basename):])
		return nil
	})
	if err != nil {
		return err
		log.Fatal(e)
	}
	return ctx.JSON(paths)
}
