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

//func (u User) FromDto() repository.User {
//	return repository.User{
//		Name:              u.Name,
//		Email:             u.Email,
//		Password:          u.Password,
//		Admin:             u.Admin,
//		Active:            u.Active,
//		ActivationCode:    u.ActivationCode,
//		ResetPasswordCode: u.ResetPasswordCode,
//		Language:          u.Language,
//		AutoSuggest:       u.AutoSuggest,
//		//Categories:        u.Categories,
//		//Emails:            u.Emails,
//		//Todos:             u.Todos,
//		//OwnedCategories:   u.OwnedCategories,
//	}
//}
//
//type FilteredUser struct {
//	ID               int       `json:"id"`
//	Email            string    `json:"email"`
//	Name             string    `json:"name"`
//	RegistrationDate time.Time `json:"registration_date"`
//}
