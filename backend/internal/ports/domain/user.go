package domain

import (
	"time"
)

type User struct {
	ID                uint
	CreatedAt         time.Time
	Name              string
	Email             string
	Password          string
	Admin             bool
	Active            bool
	ActivationCode    string
	ResetPasswordCode string
	Language          string
	AutoSuggest       bool
	Categories        *[]Category
	Emails            *[]Email
	Todos             *[]Todo
	OwnedCategories   *[]Category
}

//type FilteredUser struct {
//	ID               int       `json:"id"`
//	Email            string    `json:"email"`
//	Name             string    `json:"name"`
//	RegistrationDate time.Time `json:"registration_date"`
//}
