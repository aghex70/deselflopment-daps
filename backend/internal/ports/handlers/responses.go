package handlers

import (
	"github.com/aghex70/daps/internal/ports/domain"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	UserID      uint   `json:"user_id"`
}

type ListCategoriesResponse struct {
	Categories []domain.Category `json:"categories"`
}

type ListUsersResponse struct {
	Users []domain.User `json:"users"`
	//Users []domain2.FilteredUser `json:"users"`
}

type ListTodosResponse struct {
	Todos []domain.Todo `json:"todos"`
}

type GetUserResponse struct {
	User domain.User `json:"user"`
}
