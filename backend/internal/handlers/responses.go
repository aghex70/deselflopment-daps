package handlers

import (
	"github.com/aghex70/daps/internal/core/domain"
)

type ListTodosResponse struct {
	Todos []domain.Todo `json:"todos"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

type ListCategoriesResponse struct {
	Categories []domain.Category `json:"categories"`
}

type SummaryResponse struct {
	Summary []domain.CategorySummary `json:"summary"`
}
