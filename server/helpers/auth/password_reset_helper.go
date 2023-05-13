package auth

import (
	"fmt"
	"time"

	"github.com/mbaraa/apollo-music/data"
	"github.com/mbaraa/apollo-music/entities"
	"github.com/mbaraa/apollo-music/errors"
	"github.com/mbaraa/apollo-music/helpers/response"
	"github.com/mbaraa/apollo-music/models"
	"github.com/mbaraa/apollo-music/utils/jwt"
	"github.com/mbaraa/apollo-music/utils/mailer"
	"golang.org/x/crypto/bcrypt"
)

type PasswordResetHelper struct {
	repo    data.CRUDRepo[models.User]
	jwtUtil jwt.Manager[entities.JSON]
}

func NewPasswordResetHelper(
	repo data.CRUDRepo[models.User],
	jwtUtil jwt.Manager[entities.JSON],
) *PasswordResetHelper {
	return &PasswordResetHelper{
		repo:    repo,
		jwtUtil: jwtUtil,
	}
}

func (p *PasswordResetHelper) ResetPassword(email, origin string) (entities.JSON, int) {
	dbUser, err := p.repo.GetByConds("email = ?", email)
	if err != nil {
		return response.Build(errors.NotFound, nil)
	}

	passwordResetToken, err := p.jwtUtil.Sign(entities.JSON{
		"email": email,
	}, jwt.PasswordRestToken, time.Now().UTC().Add(time.Minute*30))
	if err != nil {
		return response.Build(errors.InternalServerError, nil)
	}

	passwordResetLink := fmt.Sprintf("%s/password-reset/%s", origin, passwordResetToken)
	err = mailer.SendPasswordReset(passwordResetLink, dbUser[0].Email)
	if err != nil {
		return response.Build(errors.InternalServerError, nil)
	}

	return response.Build(errors.None, nil)
}

func (p *PasswordResetHelper) UpdatePassword(token, newPassword string) (entities.JSON, int) {
	claims, err := p.jwtUtil.Decode(token, jwt.PasswordRestToken)
	if err != nil {
		return response.Build(errors.InvalidToken, nil)
	}

	if len(newPassword) < 8 {
		return response.Build(errors.PasswordTooShort, nil)
	}

	dbUser, err := p.repo.GetByConds("email = ?", claims.Payload["email"])
	if err != nil {
		return response.Build(errors.NotFound, nil)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.MinCost)
	if err != nil {
		return response.Build(errors.InternalServerError, nil)
	}

	err = p.repo.Update(&models.User{
		Password: string(hashedPassword),
	}, "id = ?", dbUser[0].Id)
	if err != nil {
		return response.Build(errors.InternalServerError, nil)
	}

	return response.Build(errors.None, nil)
}
