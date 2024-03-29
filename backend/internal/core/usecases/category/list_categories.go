package category

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	"github.com/aghex70/daps/internal/ports/services/category"
	"github.com/aghex70/daps/internal/ports/services/user"

	//"github.com/aghex70/daps/server"
	"log"
)

type ListCategoriesUseCase struct {
	CategoryService category.Servicer
	UserService     user.Servicer
	logger          *log.Logger
}

func (uc *ListCategoriesUseCase) Execute(ctx context.Context, filters *map[string]interface{}, userID uint) ([]domain.Category, error) {
	u, err := uc.UserService.Get(ctx, userID)
	if err != nil {
		return []domain.Category{}, err
	}

	if !u.Active {
		return []domain.Category{}, pkg.InactiveUserError
	}

	// Set the user ID into the filters map
	if filters == nil {
		filters = &map[string]interface{}{}
		(*filters)["owner_id"] = userID
	} else {
		(*filters)["owner_id"] = userID
	}
	categories, err := uc.CategoryService.List(ctx, nil, filters)
	if err != nil {
		return []domain.Category{}, err
	}
	return categories, nil
}

func NewListCategoriesUseCase(s category.Servicer, u user.Servicer, logger *log.Logger) *ListCategoriesUseCase {
	return &ListCategoriesUseCase{
		CategoryService: s,
		UserService:     u,
		logger:          logger,
	}
}
