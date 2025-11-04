package repository

import (
	"context"
	"user-service/internal/model"

	"gorm.io/gorm"
)

type CountryPhoneRepository struct {
	db *gorm.DB
}

func NewCountryPhoneRepository(db *gorm.DB) *CountryPhoneRepository {
	return &CountryPhoneRepository{db: db}
}

func (r *CountryPhoneRepository) GetByCountryCode(ctx context.Context, countryCode string) (*model.CountryPhoneFormat, error) {
	var format model.CountryPhoneFormat
	err := r.db.WithContext(ctx).
		Where("country_code = ? AND is_active = ?", countryCode, true).
		First(&format).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &format, nil
}

func (r *CountryPhoneRepository) IsCountrySupported(ctx context.Context, countryCode string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).
		Model(&model.CountryPhoneFormat{}).
		Where("country_code = ? AND is_active = ?", countryCode, true).
		Count(&count).Error

	return count > 0, err
}
