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

type ListCategoryUsersUseCase struct {
	CategoryService category.Servicer
	UserService     user.Servicer
	logger          *log.Logger
}

func (uc *ListCategoryUsersUseCase) Execute(ctx context.Context, r requests.GetCategoryRequest, userID uint) ([]domain.CategoryUser, error) {
	u, err := uc.UserService.Get(ctx, userID)
	if err != nil {
		return []domain.CategoryUser{}, err
	}

	if !u.Active {
		return []domain.CategoryUser{}, pkg.InactiveUserError
	}

	if !u.Admin {
		return []domain.CategoryUser{}, pkg.UnauthorizedError
	}

	us, err := uc.CategoryService.ListCategoryUsers(ctx, r.CategoryID)
	if err != nil {
		return []domain.CategoryUser{}, err
	}
	return us, nil
}

func NewListCategoryUsersUseCase(categoryService category.Servicer, userService user.Servicer, logger *log.Logger) *ListCategoryUsersUseCase {
	return &ListCategoryUsersUseCase{
		CategoryService: categoryService,
		UserService:     userService,
		logger:          logger,
	}
}
