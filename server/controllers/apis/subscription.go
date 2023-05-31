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

type SubscriptionApi struct {
	helper *helpers.SubscriptionHelper
}

func NewSubscriptionApi(
	helper *helpers.SubscriptionHelper,
) *SubscriptionApi {
	return &SubscriptionApi{
		helper: helper,
	}
}

func (s *SubscriptionApi) Bind(app *fiber.App) {
	subscription := app.Group("/subscription")

	subscription.Use(middlewares.AllowJSON)
	subscription.Use(middlewares.AllowCors)

	subscription.Get("/check", s.handleCheckSubscription)
	subscription.Post("/start", s.handleStartSubscription)
	subscription.Post("/cancel", s.handleCancelSubscription)
}

func (s *SubscriptionApi) handleCheckSubscription(ctx *fiber.Ctx) error {
	var (
		resp   entities.JSON
		status int
		token  = ctx.Get("Authorization")
	)
	if len(token) == 0 {
		resp, status = response.Build(errors.BadRequest, nil)
		return ctx.Status(status).JSON(resp)
	}

	resp, status = s.helper.CheckSubscription(token)
	return ctx.Status(status).JSON(resp)
}

func (s *SubscriptionApi) handleStartSubscription(ctx *fiber.Ctx) error {
	var (
		resp   entities.JSON
		status int
		req    struct {
			Plan      enums.Plan `json:"plan"`
			CardToken string     `json:"cardToken"`
		}
		token = ctx.Get("Authorization")
	)
	err := ctx.BodyParser(&req)
	if err != nil || len(token) == 0 {
		resp, status = response.Build(errors.BadRequest, nil)
		return ctx.Status(status).JSON(resp)
	}

	resp, status = s.helper.StartSubscription(token, req.CardToken, req.Plan)
	return ctx.Status(status).JSON(resp)
}

func (s *SubscriptionApi) handleCancelSubscription(ctx *fiber.Ctx) error {
	var (
		resp   entities.JSON
		status int
		token  = ctx.Get("Authorization")
	)
	if len(token) == 0 {
		resp, status = response.Build(errors.BadRequest, nil)
		return ctx.Status(status).JSON(resp)
	}

	resp, status = s.helper.CancelSubscription(token)
	return ctx.Status(status).JSON(resp)
}
