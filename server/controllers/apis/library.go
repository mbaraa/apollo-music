package apis

import (
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/mbaraa/apollo-music/entities"
	"github.com/mbaraa/apollo-music/errors"
	"github.com/mbaraa/apollo-music/helpers"
	"github.com/mbaraa/apollo-music/helpers/response"
	"github.com/mbaraa/apollo-music/middlewares"
)

type LibraryApi struct {
	musicHelper *helpers.LibraryHelper
}

func NewLibraryApi(
	musicHelper *helpers.LibraryHelper,
) *LibraryApi {
	return &LibraryApi{
		musicHelper: musicHelper,
	}
}

func (l *LibraryApi) Bind(app *fiber.App) {
	library := app.Group("/library")
	library.Use(middlewares.AllowCors)
	library.Use(middlewares.AllowJSON)

	library.Get("/music", l.handleGetMusic)
	library.Get("/album/:publicId", l.handleGetAlbum)
	// quary param fetch-covers:bool
	library.Get("/albums", l.handleGetAlbums)
	library.Get("/artist/:publicId", l.handleGetArtist)
	library.Get("/artists", l.handleGetArtists)
	library.Get("/year/:publicId", l.handleGetYear)
	library.Get("/years", l.handleGetYears)
	library.Get("/genre/:publicId", l.handleGetGenre)
	library.Get("/genres", l.handleGetGenres)
}

func (l *LibraryApi) handleGetMusic(ctx *fiber.Ctx) error {
	var (
		resp   entities.JSON
		status int
		token  = ctx.Get("Authorization")
	)
	if len(token) == 0 {
		resp, status = response.Build(errors.BadRequest, nil)
		return ctx.Status(status).JSON(resp)
	}

	resp, status = l.musicHelper.GetMusic(token)
	return ctx.Status(status).JSON(resp)
}

func (l *LibraryApi) handleGetAlbum(ctx *fiber.Ctx) error {
	var (
		resp             entities.JSON
		status           int
		token            = ctx.Get("Authorization")
		albumPublicId, _ = url.PathUnescape(ctx.Params("publicId"))
	)
	if len(token) == 0 {
		resp, status = response.Build(errors.BadRequest, nil)
		return ctx.Status(status).JSON(resp)
	}

	resp, status = l.musicHelper.GetAlbum(token, albumPublicId)
	return ctx.Status(status).JSON(resp)
}

func (l *LibraryApi) handleGetAlbums(ctx *fiber.Ctx) error {
	var (
		resp        entities.JSON
		status      int
		token       = ctx.Get("Authorization")
		fetchCovers = ctx.Query("fetch-covers") == "true"
	)
	if len(token) == 0 {
		resp, status = response.Build(errors.BadRequest, nil)
		return ctx.Status(status).JSON(resp)
	}

	resp, status = l.musicHelper.GetAlbums(token, fetchCovers)
	return ctx.Status(status).JSON(resp)
}

func (l *LibraryApi) handleGetArtist(ctx *fiber.Ctx) error {
	var (
		resp              entities.JSON
		status            int
		token             = ctx.Get("Authorization")
		artistPublicId, _ = url.PathUnescape(ctx.Params("publicId"))
	)
	if len(token) == 0 {
		resp, status = response.Build(errors.BadRequest, nil)
		return ctx.Status(status).JSON(resp)
	}

	resp, status = l.musicHelper.GetArtist(token, artistPublicId)
	return ctx.Status(status).JSON(resp)
}

func (l *LibraryApi) handleGetArtists(ctx *fiber.Ctx) error {
	var (
		resp   entities.JSON
		status int
		token  = ctx.Get("Authorization")
	)
	if len(token) == 0 {
		resp, status = response.Build(errors.BadRequest, nil)
		return ctx.Status(status).JSON(resp)
	}

	resp, status = l.musicHelper.GetArtists(token)
	return ctx.Status(status).JSON(resp)
}

func (l *LibraryApi) handleGetYear(ctx *fiber.Ctx) error {
	var (
		resp            entities.JSON
		status          int
		token           = ctx.Get("Authorization")
		yearPublicId, _ = url.PathUnescape(ctx.Params("publicId"))
	)
	if len(token) == 0 {
		resp, status = response.Build(errors.BadRequest, nil)
		return ctx.Status(status).JSON(resp)
	}

	resp, status = l.musicHelper.GetYear(token, yearPublicId)
	return ctx.Status(status).JSON(resp)
}

func (l *LibraryApi) handleGetYears(ctx *fiber.Ctx) error {
	var (
		resp   entities.JSON
		status int
		token  = ctx.Get("Authorization")
	)
	if len(token) == 0 {
		resp, status = response.Build(errors.BadRequest, nil)
		return ctx.Status(status).JSON(resp)
	}

	resp, status = l.musicHelper.GetYears(token)
	return ctx.Status(status).JSON(resp)
}

func (l *LibraryApi) handleGetGenre(ctx *fiber.Ctx) error {
	var (
		resp             entities.JSON
		status           int
		token            = ctx.Get("Authorization")
		genrePublicId, _ = url.PathUnescape(ctx.Params("publicId"))
	)
	if len(token) == 0 {
		resp, status = response.Build(errors.BadRequest, nil)
		return ctx.Status(status).JSON(resp)
	}

	resp, status = l.musicHelper.GetGenre(token, genrePublicId)
	return ctx.Status(status).JSON(resp)
}

func (l *LibraryApi) handleGetGenres(ctx *fiber.Ctx) error {
	var (
		resp   entities.JSON
		status int
		token  = ctx.Get("Authorization")
	)
	if len(token) == 0 {
		resp, status = response.Build(errors.BadRequest, nil)
		return ctx.Status(status).JSON(resp)
	}

	resp, status = l.musicHelper.GetGenres(token)
	return ctx.Status(status).JSON(resp)
}
