package repository

import (
	"github.com/jmoiron/sqlx"
)

// Authorization contains database methods for auth
type Authorization interface {
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
