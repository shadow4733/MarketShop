package impl

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"time"
	"user-service/internal/dto/request"
	"user-service/internal/dto/response"
	"user-service/internal/dto/response/mapper"
	"user-service/internal/model"
	"user-service/internal/repository"
	"user-service/internal/service"
	"user-service/internal/service/validation"
)

type UserServiceImpl struct {
	userRepo   *repository.UserRepository
	db         *gorm.DB
	validation validation.UserValidation
}

func NewUserServiceImpl(db *gorm.DB) service.UserService {
	userRepo := repository.NewUserRepository(db)
	return &UserServiceImpl{
		userRepo: userRepo,
		db:       db,
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
	existing, err := u.userRepo.GetByEmailOrUsername(ctx, data.Email, data.Username)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		if existing.Email == data.Email {
			return nil, errors.New("Email уже занят")
		}
		return nil, errors.New("Username уже занят")
	}

	if err := u.validation.ValidateCreateRequest(ctx, data); err != nil {
		return nil, err
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
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) SoftDelete(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) UpdateProfile(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) UpdateAvatar(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) GetBalance(ctx context.Context, id uuid.UUID) (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) UpdateBalance(ctx context.Context, id uuid.UUID, amount float64, operationType string) error {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) AddToBalance(ctx context.Context, id uuid.UUID, amount float64) error {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) DeductFromBalance(ctx context.Context, id uuid.UUID, amount float64) error {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) GetFinancialSummary(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) BecomeSeller(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) UpdateSellerRating(ctx context.Context, id uuid.UUID, newRating float64) error {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) GetSellerStats(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) ListUsers(ctx context.Context) ([]*model.User, int64, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) SearchUsers(ctx context.Context) ([]*model.User, int64, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) ListSellers(ctx context.Context) ([]*model.User, int64, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) GetUsersByIDs(ctx context.Context, ids []uuid.UUID) ([]*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) GetUsersByEmails(ctx context.Context, emails []string) ([]*model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) GetUserStats(ctx context.Context, id uuid.UUID) error {
	//TODO implement me
	panic("implement me")
}

func (u UserServiceImpl) GetPlatformStats(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}
