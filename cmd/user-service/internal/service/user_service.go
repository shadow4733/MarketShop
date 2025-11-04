package service

import (
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"user-service/internal/dto/request"
	"user-service/internal/dto/response"
	"user-service/internal/model"
)

type UserService interface {
	// Basic CRUD operations
	GetByID(ctx context.Context, id uuid.UUID) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	Create(ctx context.Context, data *request.Create) (*response.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id uuid.UUID) error
	SoftDelete(ctx context.Context, id uuid.UUID) error

	// Profile management
	UpdateProfile(ctx context.Context, id uuid.UUID) error
	UpdateAvatar(ctx context.Context, id uuid.UUID) error

	// Financial operations
	GetBalance(ctx context.Context, id uuid.UUID) (float64, error)
	UpdateBalance(ctx context.Context, id uuid.UUID, amount float64, operationType string) error
	AddToBalance(ctx context.Context, id uuid.UUID, amount float64) error
	DeductFromBalance(ctx context.Context, id uuid.UUID, amount float64) error
	GetFinancialSummary(ctx context.Context, id uuid.UUID) error

	// Seller management
	BecomeSeller(ctx context.Context, id uuid.UUID) error
	UpdateSellerRating(ctx context.Context, id uuid.UUID, newRating float64) error
	GetSellerStats(ctx context.Context, id uuid.UUID) error

	// Search and listing
	ListUsers(ctx context.Context) ([]*model.User, int64, error)
	SearchUsers(ctx context.Context) ([]*model.User, int64, error)
	ListSellers(ctx context.Context) ([]*model.User, int64, error)

	// Bulk operations
	GetUsersByIDs(ctx context.Context, ids []uuid.UUID) ([]*model.User, error)
	GetUsersByEmails(ctx context.Context, emails []string) ([]*model.User, error)

	// Statistics and analytics
	GetUserStats(ctx context.Context, id uuid.UUID) error
	GetPlatformStats(ctx context.Context) error
}
