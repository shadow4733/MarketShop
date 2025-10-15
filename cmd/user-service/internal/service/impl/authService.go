package impl

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"user-service/internal/dto/request"
	"user-service/internal/dto/response"
	"user-service/internal/model"
	"user-service/internal/repository"
	"user-service/internal/service"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// AuthServiceImpl реализует service.AuthService
type AuthServiceImpl struct {
	userRepo  *repository.UserRepository
	jwtSecret string
}

func NewAuthService(db *gorm.DB) service.AuthService {
	userRepo := repository.NewUserRepository(db)
	return &AuthServiceImpl{
		userRepo: userRepo,
	}
}

func (s *AuthServiceImpl) Authorization(req request.RegisterRequest) (*response.RegisterResponse, error) {
	// существующая реализация...
	emailExists, err := s.userRepo.EmailExists(req.Email)
	if err != nil {
		return nil, fmt.Errorf("ошибка проверки email: %w", err)
	}
	if emailExists {
		return nil, errors.New("пользователь с таким email уже существует")
	}

	usernameExists, err := s.userRepo.UsernameExists(req.Username)
	if err != nil {
		return nil, fmt.Errorf("ошибка проверки username: %w", err)
	}
	if usernameExists {
		return nil, errors.New("пользователь с таким username уже существует")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("ошибка хеширования пароля: %w", err)
	}

	now := time.Now()
	user := &model.User{
		ID:           uuid.New(),
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(passwordHash),
		Phone:        req.Phone,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		Country:      req.Country,
		City:         req.City,
		IsActive:     true,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	if err := s.userRepo.Create(user); err != nil {
		return nil, fmt.Errorf("ошибка создания пользователя: %w", err)
	}

	return &response.RegisterResponse{
		UserID:    user.ID.String(),
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt,
	}, nil
}
