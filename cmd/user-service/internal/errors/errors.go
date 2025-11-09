package errors

import "errors"

var (
	ErrEmailAlreadyExists    = errors.New("email уже занят")
	ErrUsernameAlreadyExists = errors.New("username уже занят")
	ErrInvalidCredentials    = errors.New("неверные учетные данные")
	ErrUserNotFound          = errors.New("пользователь не найден")
	ErrInvalidInput          = errors.New("неверные данные запроса")
)
