package jwt

import (
	"errors"
	"time"

	"github.com/mbaraa/apollo-music/config"
	"github.com/mbaraa/apollo-music/models"

	"github.com/golang-jwt/jwt/v4"
)

// interface checker
var _ Manager[models.User] = &SessionJWTManager{}

// Claims is iondsa, it's just JWT claims blyat!
type Claims[T any] struct {
	jwt.RegisteredClaims
	models.User
}

// SessionJWTManager implements JWTManager to verify session tokens
type SessionJWTManager struct{}

// NewSessionJWTManager returns a new SessionJWTManager instance,
// and since session tokens are to validate users the working type is models.User
func NewSessionJWTManager() *SessionJWTManager {
	return &SessionJWTManager{}
}

// Sign returns a JWT string(which will be the session token) based on the set JWT secret,
// using HS256 algorithm, and validity for 30 days
// and an occurring error
func (s *SessionJWTManager) Sign(user models.User, subject Subject) (string, error) {
	if len(user.Email) == 0 {
		return "", errors.New("empty email")
	}

	expirationDate := jwt.NumericDate{Time: time.Now().UTC().Add(time.Hour * 24 * 30)} // you get 30 days of this, enjoy your life
	currentTime := jwt.NumericDate{Time: time.Now().UTC()}

	claims := Claims[models.User]{
		User: user,
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
	tokenString, _ := token.SignedString(config.JWTSecret())

	return tokenString, nil
}

// Validate checks the validity of the JWT string, and returns an occurring error
// nil means you're all good
// possible errors are errors.ErrInvalidToken, and errors.ErrTokenExpired
func (s *SessionJWTManager) Validate(token string) error {
	_, err := s.Decode(token)
	if err != nil {
		return err
	}

	return nil
}

// Decode decodes the given token using the set JWT secret
func (s *SessionJWTManager) Decode(token string) (models.User, error) {
	if len(token) == 0 {
		return models.User{}, errors.New("empty token")
	}

	claims := Claims[models.User]{}

	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		if claims.ExpiresAt.Time.Before(time.Now().UTC()) {
			return nil, errors.New("token was expired")
		}

		return config.JWTSecret(), nil
	})

	if err != nil {
		return models.User{}, err
	}

	return claims.User, nil
}
