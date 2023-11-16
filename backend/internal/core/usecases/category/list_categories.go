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

func (uc *ListCategoriesUseCase) Execute(ctx context.Context, ids *[]uint, filters *map[string]interface{}) ([]domain.Category, error) {
	categories, err := uc.CategoryService.List(ctx, ids, filters)
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
