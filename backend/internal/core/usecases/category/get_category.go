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

type GetCategoryUseCase struct {
	CategoryService category.Service
	logger          *log.Logger
}

func (uc *GetCategoryUseCase) Execute(ctx context.Context, id uint) (domain.Category, error) {
	userID, _ := server.RetrieveJWTClaims(r, req)
	cs, err := uc.CategoryService.List(ctx, nil)
	if err != nil {
		return domain.Category{}, err
	}

	c, err := utils.CanRetrieveCategory(cs, id)
	if err != nil {
		return domain.Category{}, pkg.UnauthorizedError
	}

	return c, nil
}

func NewGetCategoryUseCase(s category.Service, logger *log.Logger) *GetCategoryUseCase {
	return &GetCategoryUseCase{
		CategoryService: s,
		logger:          logger,
	}
}
