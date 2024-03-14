package handlers

import (
	"github.com/aghex70/daps/internal/ports/domain"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	Admin       bool   `json:"admin"`
	UserID      uint   `json:"user_id"`
}

type ListCategoriesResponse struct {
	Categories []domain.FilteredCategory `json:"categories"`
}

type ListCategoryUsersResponse struct {
	Users []domain.CategoryUser `json:"users"`
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

type ListNotesResponse struct {
	Notes []domain.Note `json:"notes"`
}

type CreateNoteResponse struct {
	ID uint `json:"id"`
}
