package category

import "github.com/aghex70/daps/internal/ports/domain"

type ListCategoryUsersResponse struct {
	Users []domain.CategoryUser `json:"users"`
}

type GetCategoryResponse struct {
	domain.FilteredCategory
}
