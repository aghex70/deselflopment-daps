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

type UnshareCategoryUseCase struct {
	CategoryService category.Servicer
	UserService     user.Servicer
	logger          *log.Logger
}

func (uc *UnshareCategoryUseCase) Execute(ctx context.Context, r requests.UpdateCategoryRequest, userID uint) error {
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

	fields := map[string]interface{}{"shared": false}
	if err = uc.CategoryService.Update(ctx, c.ID, &fields); err != nil {
		return err
	}
	return nil
}

func NewUnshareCategoryUseCase(s category.Servicer, u user.Servicer, logger *log.Logger) *UnshareCategoryUseCase {
	return &UnshareCategoryUseCase{
		CategoryService: s,
		UserService:     u,
		logger:          logger,
	}
}