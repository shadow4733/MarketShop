package errors

import "errors"

var (
	ErrEmailAlreadyExists    = errors.New("email is already taken")
	ErrUsernameAlreadyExists = errors.New("username is already taken")
	ErrInvalidCredentials    = errors.New("incorrect credentials")
	ErrUserNotFound          = errors.New("user not found")
	ErrInvalidInput          = errors.New("invalid request data")

	ErrUsernameRequired   = errors.New("username is required")
	ErrEmailRequired      = errors.New("email is required")
	ErrPasswordRequired   = errors.New("password is required")
	ErrPasswordTooShort   = errors.New("password must be at least 6 characters")
	ErrInvalidEmailFormat = errors.New("invalid email format")
	ErrPhoneTooShort      = errors.New("phone number is too short")
)
