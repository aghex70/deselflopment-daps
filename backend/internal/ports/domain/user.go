package domain

import (
	"time"
)

type User struct {
	ID                uint       `json:"id"`
	CreatedAt         time.Time  `json:"created_at"`
	Name              string     `json:"name"`
	Email             string     `json:"email"`
	Password          string     `json:"password"`
	Admin             bool       `json:"admin"`
	Active            bool       `json:"active"`
	ActivationCode    string     `json:"activation_code"`
	ResetPasswordCode string     `json:"reset_password_code"`
	Language          string     `json:"language"`
	AutoSuggest       bool       `json:"auto_suggest"`
	Categories        []Category `json:"categories"`
	Emails            []Email    `json:"emails"`
	Todos             []Todo     `json:"todos"`
	OwnedCategories   []Category `json:"owned_categories"`
}

type FilteredUser struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Admin     bool      `json:"admin"`
	Active    bool      `json:"active"`
}

type CategoryUser struct {
	UserID  uint   `json:"user_id"`
	Email   string `json:"email"`
	IsOwner bool   `json:"is_owner"`
}
