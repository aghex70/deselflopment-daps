package category

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	requests "github.com/aghex70/daps/internal/ports/requests/category"
	"github.com/aghex70/daps/internal/ports/services/category"
	utils "github.com/aghex70/daps/utils/category"
	"log"
)

type ShareCategoryUseCase struct {
	CategoryService category.Servicer
	logger          *log.Logger
}

func (uc *ShareCategoryUseCase) Execute(ctx context.Context, r requests.UpdateCategoryRequest, userID uint) error {
	c, err := uc.CategoryService.Get(ctx, r.CategoryID)
	if err != nil {
		return err
	}
	owner := utils.IsCategoryOwner(c.OwnerID, userID)
	if !owner {
		return pkg.UnauthorizedError
	}

	fields := map[string]interface{}{"shared": true}
	err = uc.CategoryService.Update(ctx, c.ID, &fields)
	if err != nil {
		return err
	}
	return nil
}

func NewShareCategoryUseCase(s category.Servicer, logger *log.Logger) *ShareCategoryUseCase {
	return &ShareCategoryUseCase{
		CategoryService: s,
		logger:          logger,
	}
}
