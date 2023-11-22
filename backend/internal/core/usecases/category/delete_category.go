package category

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	requests "github.com/aghex70/daps/internal/ports/requests/category"
	"github.com/aghex70/daps/internal/ports/services/category"
	"github.com/aghex70/daps/internal/ports/services/user"
	utils "github.com/aghex70/daps/utils/category"
	"log"
)

type DeleteCategoryUseCase struct {
	CategoryService category.Servicer
	UserService     user.Servicer
	logger          *log.Logger
}

func (uc *DeleteCategoryUseCase) Execute(ctx context.Context, r requests.DeleteCategoryRequest, userID uint) error {
	u, err := uc.UserService.Get(ctx, userID)
	if err != nil {
		return err
	}

	if !u.Active {
		return pkg.InactiveUserError
	}

	c, err := uc.CategoryService.Get(ctx, r.CategoryID)
	if err != nil {
		return err
	}
	owner := utils.IsCategoryOwner(c.OwnerID, userID)
	if !owner {
		return pkg.UnauthorizedError
	}

	if err = uc.CategoryService.Delete(ctx, r.CategoryID); err != nil {
		return err
	}
	return nil
}

func NewDeleteCategoryUseCase(s category.Servicer, u user.Servicer, logger *log.Logger) *DeleteCategoryUseCase {
	return &DeleteCategoryUseCase{
		CategoryService: s,
		UserService:     u,
		logger:          logger,
	}
}
