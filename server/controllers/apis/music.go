package apis

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/mbaraa/apollo-music/config"
	"github.com/mbaraa/apollo-music/middlewares"
)

type LsMusicApi struct{}

func NewLsMusicApi() *LsMusicApi {
	return &LsMusicApi{}
}

func (e *LsMusicApi) Bind(app *fiber.App) {
	app.Use(middlewares.AllowCors)
	app.Static("/music-static", config.MusicDirectory())
	lsMusicGroup := app.Group("/ls-music")
	lsMusicGroup.Use(middlewares.AllowJSON)
	lsMusicGroup.Use(middlewares.AllowCors)
	lsMusicGroup.Get("/", e.handleLsMusic)
}

func (e *LsMusicApi) handleLsMusic(ctx *fiber.Ctx) error {

	paths := make([]string, 0)

	err := filepath.Walk(config.MusicDirectory(), func(path string, info os.FileInfo, err error) error {
		// fmt.Println(path)
		// fmt.Println(info.Name())
		// fmt.Println(info)

		basename := filepath.Base(config.MusicDirectory())
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
