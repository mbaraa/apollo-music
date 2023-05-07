package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// Subject represents the token's subject
type Subject = string

const (
	SessionToken      Subject = "SESSION_TOKEN"
	OtpToken          Subject = "OTP_TOKEN"
	CheckoutToken     Subject = "CHECKOUT_TOKEN"
	PasswordRestToken Subject = "PASSWORD_RESET_TOKEN"
	InactiveUserToken Subject = "INACTIVE_USER_TOKEN"
)

// Claims is iondsa, it's just JWT claims blyat!
type Claims[T any] struct {
	jwt.RegisteredClaims
	Payload T `json:"payload"`
}

// Signer is a wrapper to JWT signing method using the set JWT secret,
// claims are set(mostly unique) in each implementation of the thing
type Signer[T any] interface {
	Sign(data T, subject Subject, expTime time.Time) (string, error)
}

// Validator is a wrapper to JWT validation stuff, also uses the claims for that current implementation
type Validator interface {
	Validate(token string, subject Subject) error
}

// Decoder is a wrapper to JWT decoding stuff, based on the implementation's claims,
// this interface is usually implemented with the other two(Signer and Validator), because reasons...
type Decoder[T any] interface {
	Decode(token string, subject Subject) (Claims[T], error)
}

// Manager is a wrapper to JWT operations, so I don't do much shit each time I work with JWT
type Manager[T any] interface {
	Signer[T]
	Validator
	Decoder[T]
}
