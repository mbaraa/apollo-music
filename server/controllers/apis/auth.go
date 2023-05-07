package apis

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mbaraa/apollo-music/entities"
	"github.com/mbaraa/apollo-music/helpers"
	"github.com/mbaraa/apollo-music/middlewares"
)

type AuthApi struct {
	helper *helpers.EmailHelper
}

func NewAuthApi(helper *helpers.EmailHelper) *AuthApi {
	return &AuthApi{helper}
}

func (a *AuthApi) Bind(app *fiber.App) {
	auth := app.Group("/auth")

	auth.Use(middlewares.AllowJSON)
	auth.Use(middlewares.AllowCors)

	signinLogin := auth.Group("/signin")
	signinLogin.Post("/email", a.handleEmailSignin)

	signupLogin := auth.Group("/signup")
	signupLogin.Post("/email", a.handleEmailSignup)
}

func (a *AuthApi) handleEmailSignin(ctx *fiber.Ctx) error {
	var user entities.User
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.SendStatus(http.StatusBadRequest)
	}
	resp, status := a.helper.SigninUser(user)

	return ctx.Status(status).JSON(resp)
}

func (a *AuthApi) handleEmailSignup(ctx *fiber.Ctx) error {
	var user entities.User
	err := ctx.BodyParser(&user)
	if err != nil {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	resp, status := a.helper.SignupUser(user)

	return ctx.Status(status).JSON(resp)
}
