package domain

import "time"

type User struct {
	ID               int        `json:"id"`
	Email            string     `json:"email"`
	Password         string     `json:"password"`
	RegistrationDate time.Time  `json:"registrationDate"`
	Categories       []Category `json:"categories"`
}
