package helpers

import (
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

type OTPHelper struct {
	repo     data.CRUDRepo[models.Verification]
	userRepo data.CRUDRepo[models.User]
	jwtUtil  jwt.Manager[entities.JSON]
}

func NewOTPHelper(
	repo data.CRUDRepo[models.Verification],
	userRepo data.CRUDRepo[models.User],
	jwtUtil jwt.Manager[entities.JSON],
) *OTPHelper {
	return &OTPHelper{
		repo:     repo,
		userRepo: userRepo,
		jwtUtil:  jwtUtil,
	}
}

func (o *OTPHelper) VerifyOTP(token, verificationCode string) (entities.JSON, int) {
	claims, err := o.jwtUtil.Decode(token, jwt.OtpToken)
	if err != nil {
		return response.Build(errors.InvalidToken, nil)
	}

	if claims.Subject != jwt.OtpToken {
		return response.Build(errors.InvalidToken, nil)
	}

	dbUser, err := o.userRepo.GetByConds("email = ?", claims.Payload["email"].(string))
	if err != nil {
		return response.Build(errors.NotFound, nil)
	}

	verification, err := o.repo.GetByConds("user_id = ?", dbUser[0].Id)
	if err != nil {
		return response.Build(errors.InvalidOTP, nil)
	}

	err = bcrypt.CompareHashAndPassword([]byte(verification[len(verification)-1].OTP), []byte(verificationCode))
	if err != nil {
		return response.Build(errors.InvalidOTP, nil)
	}

	err = o.userRepo.Update(&models.User{
		Status: enums.InactiveStatus,
	}, "id = ?", dbUser[0].Id)
	if err != nil {
		return response.Build(errors.InternalServerError, nil)
	}

	checkoutToken, err := o.jwtUtil.Sign(entities.JSON{
		"email":    dbUser[0].Email,
		"publicId": dbUser[0].PublicId,
	}, jwt.CheckoutToken, time.Now().UTC().Add(time.Hour*24*30))
	if err != nil {
		return response.Build(errors.InternalServerError, nil)
	}

	err = o.repo.Delete("id = ?", verification[len(verification)-1].Id)
	if err != nil {
		return response.Build(errors.InternalServerError, nil)
	}

	return response.Build(errors.None, entities.JSON{
		"token": checkoutToken,
	})
}

func (o *OTPHelper) ResendOTP(token string) (entities.JSON, int) {
	claims, err := o.jwtUtil.Decode(token, jwt.OtpToken)
	if err != nil {
		return response.Build(errors.InvalidToken, nil)
	}

	if claims.Subject != jwt.OtpToken {
		return response.Build(errors.InvalidToken, nil)
	}

	dbUser, err := o.userRepo.GetByConds("email = ?", claims.Payload["email"].(string))
	if err != nil {
		return response.Build(errors.NotFound, nil)
	}

	otp := strings.GenerateOTP()
	hashedOTP, err := bcrypt.GenerateFromPassword([]byte(otp), bcrypt.MinCost)
	if err != nil {
		return response.Build(errors.InternalServerError, nil)
	}

	_ = o.repo.Delete("user_id = ?", dbUser[0].Id)

	expirationTime := time.Now().UTC().Add(time.Minute * 30)
	verification := models.Verification{
		UserId:      dbUser[0].Id,
		ValidBefore: expirationTime,
		OTP:         string(hashedOTP),
	}
	err = o.repo.Add(&verification)
	if err != nil {
		return response.Build(errors.InternalServerError, nil)
	}

	err = mailer.SendOTP(otp, dbUser[0].Email)
	if err != nil {
		return response.Build(errors.InternalServerError, nil)
	}

	return response.Build(errors.None, nil)
}
