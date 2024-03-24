package category

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	"github.com/aghex70/daps/internal/ports/services/category"
	"github.com/aghex70/daps/internal/ports/services/user"
	"log"
)

type GetSummaryUseCase struct {
	CategoryService category.Servicer
	UserService     user.Servicer
	logger          *log.Logger
}

func (uc *GetSummaryUseCase) Execute(ctx context.Context, userID uint) ([]domain.CategorySummary, error) {
	u, err := uc.UserService.Get(ctx, userID)
	if err != nil {
		return []domain.CategorySummary{}, err
	}

	if !u.Active {
		return []domain.CategorySummary{}, pkg.InactiveUserError
	}

	cs, err := uc.CategoryService.GetSummary(ctx, userID)
	if err != nil {
		return []domain.CategorySummary{}, err
	}
	return cs, nil
}

func NewGetSummaryUseCase(s category.Servicer, u user.Servicer, logger *log.Logger) *GetSummaryUseCase {
	return &GetSummaryUseCase{
		CategoryService: s,
		UserService:     u,
		logger:          logger,
	}
}
