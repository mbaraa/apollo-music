package jwt

// Subject represents the token's subject
type Subject = string

const (
	SessionToken      Subject = "SESSION_TOKEN"
	OtpToken          Subject = "OTP_TOKEN"
	CheckoutToken     Subject = "CHECKOUT_TOKEN"
	PasswordRestToken Subject = "PASSWORD_RESET_TOKEN"
)

// Signer is a wrapper to JWT signing method using the set JWT secret,
// claims are set(mostly unique) in each implementation of the thing
type Signer[T any] interface {
	Sign(data T, subject Subject) (string, error)
}

// Validator is a wrapper to JWT validation stuff, also uses the claims for that current implementation
type Validator interface {
	Validate(token string) error
}

// Decoder is a wrapper to JWT decoding stuff, based on the implementation's claims,
// this interface is usually implemented with the other two(Signer and Validator), because reasons...
type Decoder[T any] interface {
	Decode(token string) (T, error)
}

// Manager is a wrapper to JWT operations, so I don't do much shit each time I work with JWT
type Manager[T any] interface {
	Signer[T]
	Validator
	Decoder[T]
}
