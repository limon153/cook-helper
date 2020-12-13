package models

// User represents app user
type User struct {
	ID       int    `json:"-"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
