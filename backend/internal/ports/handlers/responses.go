package handlers

import (
	domain2 "github.com/aghex70/daps/internal/ports/domain"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	UserID      uint   `json:"user_id"`
}

type ListCategoriesResponse struct {
	Categories []domain2.Category `json:"categories"`
}

type ListUsersResponse struct {
	//Users []domain2.FilteredUser `json:"users"`
}
