package service

import (
	"cook-helper/repository"
)

// Authorization container all service methods
type Authorization interface {
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
