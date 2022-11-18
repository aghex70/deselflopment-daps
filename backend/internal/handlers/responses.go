package handlers

import (
	"github.com/aghex70/daps/internal/core/domain"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	UserId      int    `json:"user_id"`
}

type ListCategoriesResponse struct {
	Categories []domain.Category `json:"categories"`
}
