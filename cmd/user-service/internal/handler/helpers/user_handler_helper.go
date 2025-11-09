package helpers

import (
	"errors"
	"net/http"
	appErrors "user-service/internal/errors"
)

func GetStatusCode(err error) int {
	switch {
	case errors.Is(err, appErrors.ErrEmailAlreadyExists),
		errors.Is(err, appErrors.ErrUsernameAlreadyExists):
		return http.StatusConflict
	case errors.Is(err, appErrors.ErrInvalidInput):
		return http.StatusBadRequest
	case errors.Is(err, appErrors.ErrUserNotFound):
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}
