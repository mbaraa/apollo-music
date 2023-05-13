package apis

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mbaraa/apollo-music/config/env"
	"github.com/mbaraa/apollo-music/helpers"
	"github.com/mbaraa/apollo-music/middlewares"
)

type StorageApi struct {
	helper *helpers.StorageHelper
}

func NewStorageApi(helper *helpers.StorageHelper) *StorageApi {
	return &StorageApi{
		helper: helper,
	}
}

func (s *StorageApi) Bind(app *fiber.App) {
	storage := app.Group("/storage")
	storage.Use(middlewares.AllowCors)
	storage.Get("/:userPublicId/:storageType/:fileName", s.handleGetStatic)
	storage.Get("/", s.handleGetDetails)
}

func (s *StorageApi) handleGetStatic(ctx *fiber.Ctx) error {
	var (
		userPublicId = ctx.Params("userPublicId")
		storageType  = ctx.Params("storageType")
		fileName     = ctx.Params("fileName")
		token        = ctx.Get("Authorization")
	)
	if len(userPublicId) == 0 || len(storageType) == 0 ||
		len(fileName) == 0 || len(token) == 0 {
		return ctx.SendStatus(http.StatusNotFound)
	}

	if !s.helper.CheckUserAndStorage(token, userPublicId, storageType) {
		return ctx.SendStatus(http.StatusUnauthorized)
	}

	filePath := fmt.Sprintf("%s/%s/%s/%s", env.MusicDirectory(), userPublicId, storageType, fileName)
	return ctx.SendFile(filePath)
}

func (s *StorageApi) handleGetDetails(ctx *fiber.Ctx) error {
	var (
		token = ctx.Get("Authorization")
	)
	if len(token) == 0 {
		return ctx.SendStatus(http.StatusUnauthorized)
	}

	resp, status := s.helper.GetDetails(token)
	return ctx.Status(status).JSON(resp)
}
