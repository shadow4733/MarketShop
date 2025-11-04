package model

import (
	"time"
)

type CountryPhoneFormat struct {
	ID          uint      `json:"country_phone_format_id" gorm:"type:uuid;primary_key"`
	CountryCode string    `json:"country_code" gorm:"uniqueIndex;size:2;not null"`
	CountryName string    `json:"country_name" gorm:"not null"`
	PhoneRegex  string    `json:"phone_regex" gorm:"not null"`
	PhoneMask   string    `json:"phone_mask" gorm:"not null"`
	MinLength   int       `json:"min_length" gorm:"default:0"`
	MaxLength   int       `json:"max_length" gorm:"default:20"`
	IsActive    bool      `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
