package domain

import "time"

type User struct {
	Email            string
	Password         string
	RegistrationDate time.Time
	Token            string
}
