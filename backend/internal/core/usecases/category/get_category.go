package category

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	"github.com/aghex70/daps/internal/ports/services/category"
	utils "github.com/aghex70/daps/utils/category"
	"log"
)

type GetCategoryUseCase struct {
	CategoryService category.Servicer
	logger          *log.Logger
}

func (uc *GetCategoryUseCase) Execute(ctx context.Context, id uint) (domain.Category, error) {
	cs, err := uc.CategoryService.List(ctx, nil, nil)
	if err != nil {
		return domain.Category{}, err
	}

	c, err := utils.CanRetrieveCategory(cs, id)
	if err != nil {
		return domain.Category{}, pkg.UnauthorizedError
	}

	return c, nil
}

func NewGetCategoryUseCase(s category.Servicer, logger *log.Logger) *GetCategoryUseCase {
	return &GetCategoryUseCase{
		CategoryService: s,
		logger:          logger,
	}
}
