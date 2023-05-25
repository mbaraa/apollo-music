package apis

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/mbaraa/apollo-music/entities"
	"github.com/mbaraa/apollo-music/enums"
	"github.com/mbaraa/apollo-music/errors"
	"github.com/mbaraa/apollo-music/helpers"
	"github.com/mbaraa/apollo-music/helpers/response"
	"github.com/mbaraa/apollo-music/middlewares"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/checkout/session"
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

	subscription.Post("/sth", s.handleSomethingElse)
	subscription.Get("/check", s.handleCheckSubscription)
	subscription.Post("/start", s.handleStartSubscription)
	subscription.Post("/cancel", s.handleCancelSubscription)
}

func (s *SubscriptionApi) handleSomethingElse(ctx *fiber.Ctx) error {
	var req struct {
		Plan enums.Plan `json:"plan"`
	}

	err := ctx.BodyParser(&req)
	if err != nil {
		return ctx.SendStatus(400)
	}

	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String(req.Plan.PlanStripeProductId()),
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String("subscription"),
		SuccessURL: stripe.String("http://localhost:5173/success"),
		CancelURL:  stripe.String("http://localhost:5173/success"),
	}
	sess, err := session.New(params)
	if err != nil {
		return ctx.Status(500).SendString(err.Error())
	}

	fmt.Println("product", req.Plan.PlanStripeProductId())

	return ctx.Status(200).JSON(map[string]any{
		"url": sess.URL,
	})
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
