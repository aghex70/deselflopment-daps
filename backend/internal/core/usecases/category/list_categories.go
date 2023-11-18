package category

import (
	"context"
	"github.com/aghex70/daps/internal/core/services/category"
	"github.com/aghex70/daps/internal/ports/domain"
	//"github.com/aghex70/daps/server"
	"log"
)

type ListCategoriesUseCase struct {
	CategoryService category.Service
	logger          *log.Logger
}

func (uc *ListCategoriesUseCase) Execute(ctx context.Context, fields *map[string]interface{}, userID uint) ([]domain.Category, error) {
	// Set the user ID into the fields map
	if fields == nil {
		fields = &map[string]interface{}{}
		(*fields)["owner_id"] = userID
	} else {
		(*fields)["owner_id"] = userID
	}
	categories, err := uc.CategoryService.List(ctx, nil, fields)
	if err != nil {
		return []domain.Category{}, err
	}
	return categories, nil
}

func NewListCategoriesUseCase(s category.Service, logger *log.Logger) *ListCategoriesUseCase {
	return &ListCategoriesUseCase{
		CategoryService: s,
		logger:          logger,
	}
}
