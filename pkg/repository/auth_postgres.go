package repository

import (
	"cook-helper/pkg/models"
	"fmt"

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

// CreateUser inserts user into database and returns his id
func (r *AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (email, password_hash) VALUES ($1, $2) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Email, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
