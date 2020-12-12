package repository

import (
	"github.com/jmoiron/sqlx"
)

// AuthPostgres is postgres implementation for authorization repository
type AuthPostgres struct {
	db *sqlx.DB
}

// NewAuthPostgres create new AuthPostgres
func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}
