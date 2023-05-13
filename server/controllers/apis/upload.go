package apis

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mbaraa/apollo-music/entities"
	"github.com/mbaraa/apollo-music/enums"
	"github.com/mbaraa/apollo-music/errors"
	"github.com/mbaraa/apollo-music/helpers"
	"github.com/mbaraa/apollo-music/helpers/response"
	"github.com/mbaraa/apollo-music/middlewares"
)

type UploadApi struct {
	helper *helpers.UploadHelper
}

func NewUploadApi(helper *helpers.UploadHelper) *UploadApi {
	return &UploadApi{
		helper: helper,
	}
}

func (u *UploadApi) Bind(app *fiber.App) {
	upload := app.Group("/upload")
	upload.Use(middlewares.AllowCors)
	upload.Use(middlewares.AllowMultipartForm)

	upload.Post("/file/:audioType", u.handleUploadFile)
}

func (u *UploadApi) handleUploadFile(ctx *fiber.Ctx) error {
	var (
		token           = ctx.Get("Authorization")
		fileHeader, err = ctx.FormFile("audioFile")
		audioType       = enums.GetAudioType(ctx.Params("audioType"))
		resp            entities.JSON
		status          int
	)
	if len(token) == 0 || err != nil {
		resp, status = response.Build(errors.BadRequest, nil)
		return ctx.Status(status).JSON(resp)
	}

	resp, status = u.helper.UploadFile(token, audioType, fileHeader)
	return ctx.Status(status).JSON(resp)
}
