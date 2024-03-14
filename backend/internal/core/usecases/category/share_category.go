package category

import (
	"context"
	"fmt"
	"github.com/aghex70/daps/internal/pkg"
	requests "github.com/aghex70/daps/internal/ports/requests/category"
	"github.com/aghex70/daps/internal/ports/services/category"
	"github.com/aghex70/daps/internal/ports/services/user"
	utils "github.com/aghex70/daps/utils/category"
	"log"
)

type ShareCategoryUseCase struct {
	CategoryService category.Servicer
	UserService     user.Servicer
	logger          *log.Logger
}

func (uc *ShareCategoryUseCase) Execute(ctx context.Context, r requests.ShareCategoryRequest, userID uint) error {
	u, err := uc.UserService.Get(ctx, userID)
	if err != nil {
		return err
	}

	if !u.Active {
		return pkg.InactiveUserError
	}

	nu, err := uc.UserService.GetByEmail(ctx, r.Email)
	if err != nil {
		fmt.Printf("Error getting user by email: %v\n", err)
		return err
	}

	c, err := uc.CategoryService.Get(ctx, r.CategoryID)
	if err != nil {
		return err
	}
	if owner := utils.IsCategoryOwner(c.OwnerID, userID); !owner {
		return pkg.UnauthorizedError
	}

	if err = uc.CategoryService.Share(ctx, c.ID, nu); err != nil {
		return err
	}
	return nil
}

func NewShareCategoryUseCase(s category.Servicer, u user.Servicer, logger *log.Logger) *ShareCategoryUseCase {
	return &ShareCategoryUseCase{
		CategoryService: s,
		UserService:     u,
		logger:          logger,
	}
}
