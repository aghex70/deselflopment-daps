package category

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	requests "github.com/aghex70/daps/internal/ports/requests/category"
	"github.com/aghex70/daps/internal/ports/services/category"
	"github.com/aghex70/daps/internal/ports/services/user"
	"log"
)

type UnsubscribeCategoryUseCase struct {
	CategoryService category.Servicer
	UserService     user.Servicer
	logger          *log.Logger
}

func (uc *UnsubscribeCategoryUseCase) Execute(ctx context.Context, r requests.UnsubscribeCategoryRequest, userID uint) error {
	u, err := uc.UserService.Get(ctx, userID)
	if err != nil {
		return err
	}

	if !u.Active {
		return pkg.InactiveUserError
	}

	if err = uc.CategoryService.Unshare(ctx, r.CategoryID, u); err != nil {
		return err
	}
	return nil
}

func NewUnsubscribeCategoryUseCase(s category.Servicer, u user.Servicer, logger *log.Logger) *UnsubscribeCategoryUseCase {
	return &UnsubscribeCategoryUseCase{
		CategoryService: s,
		UserService:     u,
		logger:          logger,
	}
}
