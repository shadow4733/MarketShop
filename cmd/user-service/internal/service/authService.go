package service

import (
	"user-service/internal/dto/request"
	"user-service/internal/dto/response"
)

type AuthService interface {
	Authorization(req request.RegisterRequest) (*response.RegisterResponse, error)
}
