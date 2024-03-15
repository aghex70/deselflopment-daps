package category

import "github.com/aghex70/daps/internal/ports/domain"

type ListCategoriesResponse struct {
	Categories []domain.FilteredCategory `json:"categories"`
}

type ListCategoryUsersResponse struct {
	Users []domain.CategoryUser `json:"users"`
}
