package service

import (
	"context"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
	"user-service/internal/dto/request"
	"user-service/internal/dto/response"
	"user-service/internal/dto/response/mapper"
	appErrors "user-service/internal/errors"
	"user-service/internal/model"
	"user-service/internal/repository"
	"user-service/internal/service/validation"
)

type UserService interface {
	GetByID(ctx context.Context, id uuid.UUID) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	Create(ctx context.Context, data *request.Create) (*response.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type UserServiceImpl struct {
	userRepo   *repository.UserRepository
	db         *gorm.DB
	validation *validation.UserValidation
}

func NewUserServiceImpl(db *gorm.DB) UserService {
	userRepo := repository.NewUserRepository(db)
	countryRepo := repository.NewCountryPhoneRepository(db)
	userValidation := validation.NewUserValidation(countryRepo)

	return &UserServiceImpl{
		userRepo:   userRepo,
		db:         db,
		validation: userValidation,
	}
}

func (u UserServiceImpl) GetByID(ctx context.Context, id uuid.UUID) (*model.User, error) {

	panic("implement me")
}

func (u UserServiceImpl) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	panic("implement me")
}

func (u UserServiceImpl) GetByEmail(ctx context.Context, email string) (*model.User, error) {

	panic("implement me")
}

func (u *UserServiceImpl) Create(ctx context.Context, data *request.Create) (*response.User, error) {
	if err := u.validation.ValidateCreateRequest(ctx, data); err != nil {
		return nil, err
	}

	existing, err := u.userRepo.GetByEmailOrUsername(ctx, data.Email, data.Username)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		if existing.Email == data.Email {
			return nil, appErrors.ErrEmailAlreadyExists
		}
		return nil, appErrors.ErrUsernameAlreadyExists
	}

	// Хэшируем пароль
	hash, err := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		ID:           uuid.New(),
		Username:     data.Username,
		Email:        data.Email,
		PasswordHash: string(hash),
		Phone:        data.Phone,
		FirstName:    data.FirstName,
		LastName:     data.LastName,
		MiddleName:   data.MiddleName,
		DateOfBirth:  data.DateOfBirth,
		Country:      data.Country,
		City:         data.City,
		Address:      data.Address,
		PostalCode:   data.PostalCode,
		IsActive:     true,
		IsVerified:   false,
		IsSeller:     false,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := u.userRepo.CreateNewUser(ctx, user); err != nil {
		return nil, err
	}

	return mapper.FromModel(user), nil
}

func (u UserServiceImpl) Update(ctx context.Context, user *model.User) error {
	panic("implement me")
}

func (u UserServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	panic("implement me")
}
