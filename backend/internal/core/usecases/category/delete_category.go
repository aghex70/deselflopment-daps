package category

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/services/category"
	utils "github.com/aghex70/daps/utils/category"
	"log"
)

type DeleteCategoryUseCase struct {
	CategoryService category.Servicer
	logger          *log.Logger
}

func (uc *DeleteCategoryUseCase) Execute(ctx context.Context, id, userID uint) error {
	c, err := uc.CategoryService.Get(ctx, id)
	if err != nil {
		return err
	}
	owner := utils.IsCategoryOwner(c.OwnerID, userID)
	if !owner {
		return pkg.UnauthorizedError
	}

	err = uc.CategoryService.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func NewDeleteCategoryUseCase(s category.Servicer, logger *log.Logger) *DeleteCategoryUseCase {
	return &DeleteCategoryUseCase{
		CategoryService: s,
		logger:          logger,
	}
}
