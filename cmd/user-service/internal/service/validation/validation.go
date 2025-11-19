package validation

import (
	"context"
	"strings"
	"user-service/internal/dto/request"
	"user-service/internal/errors"
	"user-service/internal/repository"
)

type UserValidation struct {
	userRepo *repository.UserRepository
}

func NewUserValidation(countryRepo *repository.UserRepository) *UserValidation {
	return &UserValidation{
		userRepo: countryRepo,
	}
}

func (v *UserValidation) ValidateCreateRequest(ctx context.Context, req *request.Create) error {
	if strings.TrimSpace(req.Username) == "" {
		return errors.ErrUsernameRequired
	}
	if strings.TrimSpace(req.Email) == "" {
		return errors.ErrEmailRequired
	}
	if strings.TrimSpace(req.Password) == "" {
		return errors.ErrPasswordRequired
	}

	if len(req.Password) < 6 {
		return errors.ErrPasswordTooShort
	}

	if !strings.Contains(req.Email, "@") || !strings.Contains(req.Email, ".") {
		return errors.ErrInvalidEmailFormat
	}

	if req.Phone != "" && len(req.Phone) < 5 {
		return errors.ErrPhoneTooShort
	}

	return nil
}
