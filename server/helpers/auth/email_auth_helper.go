package auth

import (
	"log"
	"net/mail"
	"time"

	"github.com/mbaraa/apollo-music/data"
	"github.com/mbaraa/apollo-music/entities"
	"github.com/mbaraa/apollo-music/enums"
	"github.com/mbaraa/apollo-music/errors"
	"github.com/mbaraa/apollo-music/helpers/response"
	"github.com/mbaraa/apollo-music/models"
	"github.com/mbaraa/apollo-music/utils/jwt"
	"github.com/mbaraa/apollo-music/utils/mailer"
	"github.com/mbaraa/apollo-music/utils/strings"
	"golang.org/x/crypto/bcrypt"
)

type EmailHelper struct {
	repo             data.CRUDRepo[models.User]
	verificationRepo data.CRUDRepo[models.Verification]
	jwtUtil          jwt.Manager[entities.JSON]
}

func NewEmailHelper(repo data.CRUDRepo[models.User], verificationRepo data.CRUDRepo[models.Verification], jwtUtil jwt.Manager[entities.JSON]) *EmailHelper {
	return &EmailHelper{
		repo:             repo,
		verificationRepo: verificationRepo,
		jwtUtil:          jwtUtil,
	}
}

func (e *EmailHelper) SigninUser(user entities.User) (entities.JSON, int) {
	addr, err := mail.ParseAddress(user.Email)
	if len(user.Email) == 0 || err != nil || addr.Address != user.Email {
		return response.Build(errors.InvalidCredentials, nil)
	}
	if len(user.Password) < 8 {
		return response.Build(errors.PasswordTooShort, nil)
	}

	dbUser, err := e.repo.GetByConds("email = ?", user.Email)
	if err != nil {
		log.Println(err)
		return response.Build(errors.InvalidCredentials, nil)
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser[0].Password), []byte(user.Password))
	if err != nil {
		log.Println(err)
		return response.Build(errors.InvalidCredentials, nil)
	}

	expirationTime := time.Now().UTC().Add(time.Hour * 24 * 30)
	sub := jwt.SessionToken
	if dbUser[0].Status != enums.ActiveStatus {
		sub = jwt.InactiveUserToken
		expirationTime = time.Now().UTC().Add(time.Minute * 30)
	}
	if dbUser[0].Status == enums.UnverifiedEmailStatus {
		otp := strings.GenerateOTP()
		hashedOTP, err := bcrypt.GenerateFromPassword([]byte(otp), bcrypt.MinCost)
		if err != nil {
			log.Println(err)
			return response.Build(errors.InternalServerError, nil)
		}

		verification := models.Verification{
			UserId:      dbUser[0].Id,
			ValidBefore: expirationTime,
			OTP:         string(hashedOTP),
		}
		err = e.verificationRepo.Add(&verification)
		if err != nil {
			log.Println(err)
			return response.Build(errors.InternalServerError, nil)
		}

		err = mailer.SendOTP(otp, dbUser[0].Email)
		if err != nil {
			log.Println(err)
			return response.Build(errors.InternalServerError, nil)
		}
		sub = jwt.OtpToken
	}

	token, err := e.jwtUtil.Sign(entities.JSON{
		"email":    dbUser[0].Email,
		"publicId": dbUser[0].PublicId,
	}, sub, expirationTime)
	if err != nil {
		log.Println(err)
		return response.Build(errors.InternalServerError, nil)
	}

	return response.Build(errors.None, map[string]string{
		"token": token,
	})
}

func (e *EmailHelper) SignupUser(user entities.User) (entities.JSON, int) {
	if len(user.FullName) == 0 {
		return response.Build(errors.EmptyName, nil)
	}
	addr, err := mail.ParseAddress(user.Email)
	if len(user.Email) == 0 || err != nil || addr.Address != user.Email {
		return response.Build(errors.InvalidEmail, nil)
	}
	if len(user.Password) < 8 {
		return response.Build(errors.PasswordTooShort, nil)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		return response.Build(errors.InternalServerError, nil)
	}

	dbUser := models.User{
		FullName: user.FullName,
		Email:    user.Email,
		Password: string(hashedPassword),
		Status:   enums.UnverifiedEmailStatus,
	}
	err = e.repo.Add(&dbUser)
	if err != nil {
		log.Println(err)
		return response.Build(errors.EmailExists, nil)
	}

	otp := strings.GenerateOTP()
	hashedOTP, err := bcrypt.GenerateFromPassword([]byte(otp), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		return response.Build(errors.InternalServerError, nil)
	}

	expirationTime := time.Now().UTC().Add(time.Minute * 30)
	verification := models.Verification{
		UserId:      dbUser.Id,
		ValidBefore: expirationTime,
		OTP:         string(hashedOTP),
	}
	err = e.verificationRepo.Add(&verification)
	if err != nil {
		log.Println(err)
		return response.Build(errors.InternalServerError, nil)
	}

	err = mailer.SendOTP(otp, dbUser.Email)
	if err != nil {
		log.Println(err)
		return response.Build(errors.InternalServerError, nil)
	}

	token, err := e.jwtUtil.Sign(entities.JSON{
		"email":    dbUser.Email,
		"publicId": dbUser.PublicId,
	}, jwt.OtpToken, expirationTime)
	if err != nil {
		log.Println(err)
		return response.Build(errors.InternalServerError, nil)
	}

	return response.Build(errors.None, map[string]string{
		"token": token,
	})
}
