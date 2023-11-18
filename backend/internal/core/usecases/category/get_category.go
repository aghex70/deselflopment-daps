package category

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
	requests "github.com/aghex70/daps/internal/ports/requests/category"
	"github.com/aghex70/daps/internal/ports/services/category"
	"log"
)

type GetCategoryUseCase struct {
	CategoryService category.Servicer
	logger          *log.Logger
}

func (uc *GetCategoryUseCase) Execute(ctx context.Context, r requests.GetCategoryRequest, userID uint) (domain.Category, error) {
	cs, err := uc.CategoryService.Get(ctx, r.CategoryID)
	if err != nil {
		return domain.Category{}, err
	}

	//c, err := utils.CanRetrieveCategory([]domain.Category{cs}, userID)
	//if err != nil {
	//	return domain.Category{}, pkg.UnauthorizedError
	//}

	return cs, nil
}

func NewGetCategoryUseCase(s category.Servicer, logger *log.Logger) *GetCategoryUseCase {
	return &GetCategoryUseCase{
		CategoryService: s,
		logger:          logger,
	}
}
