package validation

import (
	"context"
	"errors"
	"strings"
	"user-service/internal/dto/request"
	"user-service/internal/repository"
)

type UserValidation struct {
	countryRepo *repository.CountryPhoneRepository
}

func NewUserValidation(countryRepo *repository.CountryPhoneRepository) *UserValidation {
	return &UserValidation{
		countryRepo: countryRepo,
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

	// Длина пароля
	if len(req.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}

	// Простая проверка email
	if !strings.Contains(req.Email, "@") || !strings.Contains(req.Email, ".") {
		return errors.New("invalid email format")
	}

	// Телефон опционален, но если указан - проверяем длину
	if req.Phone != "" && len(req.Phone) < 5 {
		return errors.New("phone number is too short")
	}

	return nil
}
