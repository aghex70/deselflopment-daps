package category

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	requests "github.com/aghex70/daps/internal/ports/requests/category"
	"github.com/aghex70/daps/internal/ports/services/category"
	"github.com/aghex70/daps/internal/ports/services/user"
	common "github.com/aghex70/daps/utils"
	utils "github.com/aghex70/daps/utils/category"
	"log"
)

type UpdateCategoryUseCase struct {
	CategoryService category.Servicer
	UserService     user.Servicer
	logger          *log.Logger
}

func (uc *UpdateCategoryUseCase) Execute(ctx context.Context, r requests.UpdateCategoryRequest, userID uint) error {
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

	filters := common.StructToMap(r, "category_id")
	if err = uc.CategoryService.Update(ctx, c.ID, &filters); err != nil {
		return err
	}
	return nil
}

func NewUpdateCategoryUseCase(s category.Servicer, u user.Servicer, logger *log.Logger) *UpdateCategoryUseCase {
	return &UpdateCategoryUseCase{
		CategoryService: s,
		UserService:     u,
		logger:          logger,
	}
}
