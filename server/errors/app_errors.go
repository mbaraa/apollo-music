package errors

import "github.com/joomcode/errorx"

var (
	ErrorsNS = errorx.NewNamespace("Apollo errors")
)

var (
	ErrNotFound              = ErrorsNS.NewType(NotFound.String())
	ErrUnauthorized          = ErrorsNS.NewType(Unauthorized.String())
	ErrBadRequest            = ErrorsNS.NewType(BadRequest.String())
	ErrInternalServerError   = ErrorsNS.NewType(InsufficientStorage.String())
	ErrInsufficientStorage   = ErrorsNS.NewType(InsufficientStorage.String())
	ErrUnsupportedFileFormat = ErrorsNS.NewType(UnsupportedFileFormat.String())
	ErrInvalidEmail          = ErrorsNS.NewType(InvalidEmail.String())
	ErrEmailExists           = ErrorsNS.NewType(EmailExists.String())
	ErrPasswordTooShort      = ErrorsNS.NewType(PasswordTooShort.String())
	ErrEmptyName             = ErrorsNS.NewType(EmptyName.String())
	ErrInvalidOAuthToken     = ErrorsNS.NewType(InvalidOAuthToken.String())
	ErrInvalidOTP            = ErrorsNS.NewType(InvalidOTP.String())
	ErrInvalidToken          = ErrorsNS.NewType(InvalidToken.String())
	ErrInvalidCredentials    = ErrorsNS.NewType(InvalidCredentials.String())
	ErrInsufficientFunds     = ErrorsNS.NewType(InsufficientFunds.String())
	ErrPaymentError          = ErrorsNS.NewType(PaymentError.String())
)
