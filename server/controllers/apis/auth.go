package apis

import (
	"github.com/gofiber/fiber/v2"
	"github.com/mbaraa/apollo-music/config/env"
	"github.com/mbaraa/apollo-music/entities"
	"github.com/mbaraa/apollo-music/errors"
	"github.com/mbaraa/apollo-music/helpers/auth"
	"github.com/mbaraa/apollo-music/helpers/response"
	"github.com/mbaraa/apollo-music/middlewares"
)

type AuthApi struct {
	helper              *auth.EmailHelper
	otpHelper           *auth.OTPHelper
	passwordResetHelper *auth.PasswordResetHelper
	sessionHelper       *auth.SessionHelper
}

func NewAuthApi(
	helper *auth.EmailHelper,
	otpHelper *auth.OTPHelper,
	passwordResetHelper *auth.PasswordResetHelper,
	sessionHelper *auth.SessionHelper,
) *AuthApi {
	return &AuthApi{
		helper:              helper,
		otpHelper:           otpHelper,
		passwordResetHelper: passwordResetHelper,
		sessionHelper:       sessionHelper,
	}
}

func (a *AuthApi) Bind(app *fiber.App) {
	auth := app.Group("/auth")

	auth.Use(middlewares.AllowJSON)
	auth.Use(middlewares.AllowCors)

	signinLogin := auth.Group("/signin")
	signinLogin.Post("/email", a.handleEmailSignin)

	signupLogin := auth.Group("/signup")
	signupLogin.Post("/email", a.handleEmailSignup)

	otp := auth.Group("/otp")
	otp.Post("/verify", a.handleVerifyOTP)
	otp.Get("/resend", a.handleResendOTP)

	passwordReset := auth.Group("/password")
	passwordReset.Post("/reset", a.handleResetPassword)
	passwordReset.Post("/update", a.handleUpdatePassword)

	session := auth.Group("/session")
	session.Get("/check", a.handleCheckSession)
}

func (a *AuthApi) handleEmailSignin(ctx *fiber.Ctx) error {
	var (
		resp   entities.JSON
		status int
		user   entities.User
	)
	err := ctx.BodyParser(&user)
	if err != nil {
		resp, status = response.Build(errors.BadRequest, nil)
		return ctx.Status(status).JSON(resp)
	}

	resp, status = a.helper.SigninUser(user)
	return ctx.Status(status).JSON(resp)
}

func (a *AuthApi) handleEmailSignup(ctx *fiber.Ctx) error {
	var (
		resp   entities.JSON
		status int
		user   entities.User
	)
	err := ctx.BodyParser(&user)
	if err != nil {
		resp, status = response.Build(errors.BadRequest, nil)
		return ctx.Status(status).JSON(resp)
	}

	resp, status = a.helper.SignupUser(user)
	return ctx.Status(status).JSON(resp)
}

func (a *AuthApi) handleVerifyOTP(ctx *fiber.Ctx) error {
	var (
		resp   entities.JSON
		status int
		code   struct {
			VerificationCode string `json:"verificationCode"`
		}
		token = ctx.Get("Authorization")
	)
	err := ctx.BodyParser(&code)
	if err != nil || len(token) == 0 {
		resp, status = response.Build(errors.BadRequest, nil)
		return ctx.Status(status).JSON(resp)
	}

	resp, status = a.otpHelper.VerifyOTP(token, code.VerificationCode)
	return ctx.Status(status).JSON(resp)
}

func (a *AuthApi) handleResendOTP(ctx *fiber.Ctx) error {
	var (
		resp   entities.JSON
		status int
		token  = ctx.Get("Authorization")
	)
	if len(token) == 0 {
		resp, status = response.Build(errors.BadRequest, nil)
		return ctx.Status(status).JSON(resp)
	}

	resp, status = a.otpHelper.ResendOTP(token)
	return ctx.Status(status).JSON(resp)
}

func (a *AuthApi) handleResetPassword(ctx *fiber.Ctx) error {
	var (
		resp   entities.JSON
		status int
		body   struct {
			UserEmail string `json:"userEmail"`
		}
	)
	origin := ctx.Get("Origin")
	if len(origin) == 0 {
		origin = env.FrontendAddress()
	}

	err := ctx.BodyParser(&body)
	if err != nil {
		resp, status = response.Build(errors.BadRequest, nil)
		return ctx.Status(status).JSON(resp)
	}

	resp, status = a.passwordResetHelper.ResetPassword(body.UserEmail, origin)
	return ctx.Status(status).JSON(resp)
}

func (a *AuthApi) handleUpdatePassword(ctx *fiber.Ctx) error {
	var (
		resp   entities.JSON
		status int
		body   struct {
			NewPassword string `json:"newPassword"`
		}
		token = ctx.Get("Authorization")
	)
	err := ctx.BodyParser(&body)
	if len(token) == 0 && err != nil {
		resp, status = response.Build(errors.BadRequest, nil)
		return ctx.Status(status).JSON(resp)
	}

	resp, status = a.passwordResetHelper.UpdatePassword(token, body.NewPassword)
	return ctx.Status(status).JSON(resp)
}

func (a *AuthApi) handleCheckSession(ctx *fiber.Ctx) error {
	var (
		resp   entities.JSON
		status int
		token  = ctx.Get("Authorization")
	)
	if len(token) == 0 {
		resp, status = response.Build(errors.BadRequest, nil)
		return ctx.Status(status).JSON(resp)
	}

	resp, status = a.sessionHelper.CheckSession(token)
	return ctx.Status(status).JSON(resp)
}
