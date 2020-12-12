package service

import (
	"cook-helper/repository"
)

// AuthService is authorization service implementation
type AuthService struct {
	repo repository.Authorization
}

// NewAuthService creates new AuthService
func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}
