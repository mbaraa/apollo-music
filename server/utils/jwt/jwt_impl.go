package jwt

import (
	"errors"
	"time"

	"github.com/mbaraa/apollo-music/entities"
	aerrors "github.com/mbaraa/apollo-music/errors"

	"github.com/mbaraa/apollo-music/config/env"
	"github.com/mbaraa/apollo-music/models"

	"github.com/golang-jwt/jwt/v4"
)

// interface checker
var _ Manager[entities.JSON] = &JWTImpl{}

// Claims is iondsa, it's just JWT claims blyat!
type Claims[T any] struct {
	jwt.RegisteredClaims
	Payload entities.JSON `json:"payload"`
}

// JWTImpl implements JWTManager to verify session tokens
type JWTImpl struct{}

// NewJWTImpl returns a new JWTImpl instance,
// and since session tokens are to validate users the working type is models.User
func NewJWTImpl() *JWTImpl {
	return &JWTImpl{}
}

// Sign returns a JWT string(which will be the session token) based on the set JWT secret,
// using HS256 algorithm, and validity for 30 days
// and an occurring error
func (s *JWTImpl) Sign(data entities.JSON, subject Subject, expTime time.Time) (string, error) {
	expirationDate := jwt.NumericDate{Time: expTime}
	currentTime := jwt.NumericDate{Time: time.Now().UTC()}

	claims := Claims[models.User]{
		Payload: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &expirationDate,
			Issuer:    "Apollo Music",
			Subject:   subject,
			NotBefore: &currentTime,
			IssuedAt:  &currentTime,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)

	// error is ignored here, because it's either json marshaling or signing method(using library's built-in method) errors
	//which are correct as of this version of golang-jwt, god knows what could go wrong in the future :)
	tokenString, _ := token.SignedString(env.JWTSecret())

	return tokenString, nil
}

// Validate checks the validity of the JWT string, and returns an occurring error
// nil means you're all good
// possible errors are errors.ErrInvalidToken, and errors.ErrTokenExpired
func (s *JWTImpl) Validate(token string, subject Subject) error {
	_, err := s.Decode(token, subject)
	if err != nil {
		return err
	}

	return nil
}

// Decode decodes the given token using the set JWT secret
func (s *JWTImpl) Decode(token string, subject Subject) (entities.JSON, error) {
	if len(token) == 0 {
		return nil, errors.New("empty token")
	}

	claims := Claims[models.User]{}

	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		if claims.Subject != subject {
			return nil, aerrors.ErrInvalidToken.New("")
		}
		if claims.ExpiresAt.Time.Before(time.Now().UTC()) {
			return nil, errors.New("token was expired")
		}

		return env.JWTSecret(), nil
	})

	if err != nil {
		return nil, err
	}

	return claims.Payload, nil
}
