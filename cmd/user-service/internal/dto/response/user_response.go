package response

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID  `json:"user_id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Username    string     `json:"username"`
	Email       string     `json:"email"`
	Phone       string     `json:"phone,omitempty"`
	FirstName   string     `json:"first_name,omitempty"`
	LastName    string     `json:"last_name,omitempty"`
	MiddleName  string     `json:"middle_name,omitempty"`
	DateOfBirth *time.Time `json:"date_of_birth,omitempty"`
	Balance     float64    `json:"balance"`

	Country    string `json:"country,omitempty"`
	City       string `json:"city,omitempty"`
	Address    string `json:"address,omitempty"`
	PostalCode string `json:"postal_code,omitempty"`

	IsVerified bool    `json:"is_verified"`
	IsActive   bool    `json:"is_active"`
	IsSeller   bool    `json:"is_seller"`
	Rating     float64 `json:"rating"`

	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	LastLogin *time.Time `json:"last_login_at,omitempty"`
}
