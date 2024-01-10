package handlers

import (
	"github.com/aghex70/daps/internal/ports/domain"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	UserID      uint   `json:"user_id"`
}

type ListCategoriesResponse struct {
	Categories []domain.FilteredCategory `json:"categories"`
}

type ListUsersResponse struct {
	Users []domain.FilteredUser `json:"users"`
}

type ListTodosResponse struct {
	Todos []domain.Todo `json:"todos"`
}

type GetUserResponse struct {
	User domain.FilteredUser `json:"user"`
}
