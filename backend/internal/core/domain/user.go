package domain

import "time"

type User struct {
	ID               int
	Email            string
	Password         string
	AccessToken      string
	RefreshToken     string
	RegistrationDate time.Time
}
