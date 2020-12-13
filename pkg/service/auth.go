package service

import (
	"cook-helper/pkg/models"
	"cook-helper/pkg/repository"

	"golang.org/x/crypto/bcrypt"
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

// CreateUser hashes password and passes user to repo.
// Return created user id
func (s *AuthService) CreateUser(u models.User) (int, error) {
	hash, err := generatePasswordHash(u.Password)
	if err != nil {
		return 0, err
	}

	u.Password = hash

	id, err := s.repo.CreateUser(u)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func generatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
