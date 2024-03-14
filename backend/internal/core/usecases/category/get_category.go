package category

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	requests "github.com/aghex70/daps/internal/ports/requests/category"
	"github.com/aghex70/daps/internal/ports/services/category"
	"github.com/aghex70/daps/internal/ports/services/user"
	"log"
)

type GetCategoryUseCase struct {
	CategoryService category.Servicer
	UserService     user.Servicer
	logger          *log.Logger
}

func (uc *GetCategoryUseCase) Execute(ctx context.Context, r requests.GetCategoryRequest, userID uint) (domain.Category, error) {
	u, err := uc.UserService.Get(ctx, userID)
	if err != nil {
		return domain.Category{}, err
	}

	if !u.Active {
		return domain.Category{}, pkg.InactiveUserError
	}

	cs, err := uc.CategoryService.Get(ctx, r.CategoryID)
	if err != nil {
		return domain.Category{}, err
	}

	return cs, nil
}

func NewGetCategoryUseCase(s category.Servicer, u user.Servicer, logger *log.Logger) *GetCategoryUseCase {
	return &GetCategoryUseCase{
		CategoryService: s,
		UserService:     u,
		logger:          logger,
	}
}
