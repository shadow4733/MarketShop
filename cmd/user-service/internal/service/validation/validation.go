package validation

import (
	"context"
	"errors"
	"strings"
	"user-service/internal/dto/request"
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
	// Базовые проверки
	if strings.TrimSpace(req.Username) == "" {
		return errors.New("username is required")
	}
	if strings.TrimSpace(req.Email) == "" {
		return errors.New("email is required")
	}
	if strings.TrimSpace(req.Password) == "" {
		return errors.New("password is required")
	}

	if len(req.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	if !strings.Contains(req.Email, "@") || !strings.Contains(req.Email, ".") {
		return errors.New("invalid email format")
	}

	if req.Phone != "" && len(req.Phone) < 5 {
		return errors.New("phone number is too short")
	}

	return nil
}
