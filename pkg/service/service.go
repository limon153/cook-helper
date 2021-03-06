package service

import (
	"cook-helper/pkg/models"
	"cook-helper/pkg/repository"
)

// Authorization container all service methods
type Authorization interface {
	CreateUser(user models.User) (int, error)
}

// Service contains all services
type Service struct {
	Authorization
}

// New creates new Service
func New(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
