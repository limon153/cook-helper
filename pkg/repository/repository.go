package repository

import (
	"cook-helper/pkg/models"

	"github.com/jmoiron/sqlx"
)

// Authorization contains database methods for auth
type Authorization interface {
	CreateUser(user models.User) (int, error)
}

// Repository combines all repos
type Repository struct {
	Authorization
}

// New creates new Repository
func New(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
