package category

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	"github.com/aghex70/daps/internal/ports/services/category"
	utils "github.com/aghex70/daps/utils/category"
	"log"
)

type UpdateCategoryUseCase struct {
	CategoryService category.Servicer
	logger          *log.Logger
}

func (uc *UpdateCategoryUseCase) Execute(ctx context.Context, fields map[string]interface{}, id, userID uint) (domain.Category, error) {
	c, err := uc.CategoryService.Get(ctx, id)
	if err != nil {
		return domain.Category{}, err
	}
	owner := utils.IsCategoryOwner(c.OwnerID, userID)
	if !owner {
		return domain.Category{}, pkg.UnauthorizedError
	}

	c.Shared = false
	cat, err := uc.CategoryService.Update(ctx, id, c)
	if err != nil {
		return domain.Category{}, err
	}
	return cat, nil
}

func NewUpdateCategoryUseCase(s category.Servicer, logger *log.Logger) *UpdateCategoryUseCase {
	return &UpdateCategoryUseCase{
		CategoryService: s,
		logger:          logger,
	}
}
