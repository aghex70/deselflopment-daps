package domain

import "time"

type User struct {
	ID               int
	Email            string
	Password         string
	RegistrationDate time.Time
}
