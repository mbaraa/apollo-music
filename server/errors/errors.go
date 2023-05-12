package errors

import "fmt"

// ErrorCode represents enums for all of the possible errors
type ErrorCode uint

const (
	None ErrorCode = iota
	NotFound
	Unauthorized
	BadRequest
	InternalServerError
	InsufficientStorage
	UnsupportedFileFormat
	InvalidEmail
	EmailExists
	PasswordTooShort
	EmptyName
	InvalidOAuthToken
	InvalidOTP
	InvalidToken
	InvalidCredentials
	InsufficientFunds
	PaymentError
)

// String returns the error's corresponidng string
func (e ErrorCode) String(msg ...string) string {
	if len(msg) > 0 {
		return fmt.Sprintf(errorMessages[e], ": "+msg[0])
	}
	return fmt.Sprintf(errorMessages[e], "")
}

// StatusCode returns the error's corresponidng http status code
func (e ErrorCode) StatusCode() int {
	return errorHttpStatuses[e]
}

var errorMessages = map[ErrorCode]string{
	None:                  "None%s",
	NotFound:              "Not found%s",
	Unauthorized:          "Unauthorized%s",
	BadRequest:            "Bad request%s",
	InternalServerError:   "Internal server error%s",
	InsufficientStorage:   "Insufficient storage%s",
	UnsupportedFileFormat: "Unsupported file format%s",
	InvalidEmail:          "Invalid email%s",
	EmailExists:           "Email exists%s",
	PasswordTooShort:      "Password too short%s",
	EmptyName:             "Empty name%s",
	InvalidOAuthToken:     "Invalid OAuth token%s",
	InvalidOTP:            "Invalid OTP%s",
	InvalidToken:          "Invalid token%s",
	InvalidCredentials:    "Invalid credentials%s",
	InsufficientFunds:     "Insufficient funds%s",
	PaymentError:          "Payment error%s",
}

var errorHttpStatuses = map[ErrorCode]int{
	None:                  200,
	NotFound:              404,
	Unauthorized:          401,
	BadRequest:            400,
	InternalServerError:   500,
	InsufficientStorage:   507,
	UnsupportedFileFormat: 415,
	InvalidEmail:          400,
	EmailExists:           400,
	PasswordTooShort:      400,
	EmptyName:             400,
	InvalidOAuthToken:     400,
	InvalidOTP:            401,
	InvalidToken:          401,
	InvalidCredentials:    401,
	InsufficientFunds:     402,
	PaymentError:          402,
}
