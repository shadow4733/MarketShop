package model

import (
	"github.com/google/uuid"
	"time"
)

// User represents a marketplace user entity.
type User struct {
	ID           uuid.UUID `json:"user_id" gorm:"type:uuid;primary_key"`
	Username     string    `json:"username" gorm:"uniqueIndex;not null"`
	Email        string    `json:"email" gorm:"uniqueIndex;not null"`
	PasswordHash string    `json:"-" gorm:"not null"`
	Phone        *string   `json:"phone" gorm:"index"`

	// Financial information
	Balance     float64 `json:"balance" gorm:"default:0"`
	TotalSpent  float64 `json:"total_spent" gorm:"default:0"`
	TotalEarned float64 `json:"total_earned" gorm:"default:0"`

	// Personal information
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	MiddleName  *string    `json:"middle_name"`
	DateOfBirth *time.Time `json:"date_of_birth"`
	AvatarURL   *string    `json:"avatar_url"`

	// Address information
	Country    string `json:"country"`
	City       string `json:"city"`
	Address    string `json:"address"`
	PostalCode string `json:"postal_code"`

	// Status and verification
	IsVerified  bool    `json:"is_verified" gorm:"default:false"`
	IsActive    bool    `json:"is_active" gorm:"default:true"`
	IsSeller    bool    `json:"is_seller" gorm:"default:false"`
	Rating      float64 `json:"rating" gorm:"default:0"`
	ReviewCount int     `json:"review_count" gorm:"default:0"`

	// Security
	LastPasswordChange  *time.Time `json:"last_password_change"`
	FailedLoginAttempts int        `json:"failed_login_attempts" gorm:"default:0"`
	IsLocked            bool       `json:"is_locked" gorm:"default:false"`

	// Notification settings
	EmailNotifications bool `json:"email_notifications" gorm:"default:true"`
	SMSNotifications   bool `json:"sms_notifications" gorm:"default:false"`
	PushNotifications  bool `json:"push_notifications" gorm:"default:true"`

	// Timestamps
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	LastLoginAt *time.Time `json:"last_login_at"`
	DeletedAt   *time.Time `json:"deleted_at" gorm:"index"`
}
