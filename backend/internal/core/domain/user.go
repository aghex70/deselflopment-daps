package domain

import "time"

type User struct {
	Id                int        `json:"id"`
	Email             string     `json:"email"`
	Name              string     `json:"name"`
	Password          string     `json:"password"`
	RegistrationDate  time.Time  `json:"registration_date"`
	Categories        []Category `json:"categories"`
	Admin             bool       `json:"admin"`
	ActivationCode    string     `json:"activation_code"`
	Active            bool       `json:"active"`
	ResetPasswordCode string     `json:"reset_password_code"`
}

type FilteredUser struct {
	Id               int       `json:"id"`
	Email            string    `json:"email"`
	Name             string    `json:"name"`
	RegistrationDate time.Time `json:"registration_date"`
}
