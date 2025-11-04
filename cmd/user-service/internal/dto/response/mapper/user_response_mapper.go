package mapper

import "user-service/internal/model"
import "user-service/internal/dto/response"

func FromModel(u *model.User) *response.User {
	if u == nil {
		return nil
	}
	return &response.User{
		ID:          u.ID,
		Username:    u.Username,
		Email:       u.Email,
		Phone:       u.Phone,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		MiddleName:  u.MiddleName,
		DateOfBirth: u.DateOfBirth,
		Balance:     u.Balance,
		Country:     u.Country,
		City:        u.City,
		Address:     u.Address,
		PostalCode:  u.PostalCode,
		IsVerified:  u.IsVerified,
		IsActive:    u.IsActive,
		IsSeller:    u.IsSeller,
		Rating:      u.Rating,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
		LastLogin:   u.LastLoginAt,
	}
}
