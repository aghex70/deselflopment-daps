package category

import (
	"context"
	"github.com/aghex70/daps/internal/core/services/category"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	"github.com/aghex70/daps/server"
	utils "github.com/aghex70/daps/utils/category"
	"log"
)

type DeleteCategoryUseCase struct {
	CategoryService category.Service
	logger          *log.Logger
}

func (uc *DeleteCategoryUseCase) Execute(ctx context.Context, id uint) (domain.Category, error) {
	userID, _ := server.RetrieveJWTClaims(r, req)
	u, err := uc.CategoryService.GetOwner(ctx, id)
	if err != nil {
		return domain.Category{}, err
	}
	owner := utils.IsCategoryOwner(u.OwnerID, userID)
	if !owner {
		return domain.Category{}, pkg.UnauthorizedError
	}

	err = uc.CategoryService.Delete(ctx, id)
	if err != nil {
		return domain.Category{}, err
	}
	return domain.Category{}, nil
}

func NewDeleteCategoryUseCase(s category.Service, logger *log.Logger) *DeleteCategoryUseCase {
	return &DeleteCategoryUseCase{
		CategoryService: s,
		logger:          logger,
	}
}
