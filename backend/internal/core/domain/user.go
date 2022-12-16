package domain

import "time"

type User struct {
	ID               int        `json:"id"`
	Email            string     `json:"email"`
	Name             string     `json:"name"`
	Password         string     `json:"password"`
	RegistrationDate time.Time  `json:"registration_date"`
	Categories       []Category `json:"categories"`
	IsAdmin          bool       `json:"is_admin"`
}

type FilteredUser struct {
	ID               int       `json:"id"`
	Email            string    `json:"email"`
	Name             string    `json:"name"`
	RegistrationDate time.Time `json:"registration_date"`
}
