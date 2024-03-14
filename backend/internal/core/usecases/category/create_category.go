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

type CreateCategoryUseCase struct {
	CategoryService category.Servicer
	UserService     user.Servicer
	logger          *log.Logger
}

func (uc *CreateCategoryUseCase) Execute(ctx context.Context, userID uint, r requests.CreateCategoryRequest) (domain.Category, error) {
	u, err := uc.UserService.Get(ctx, userID)
	if err != nil {
		return domain.Category{}, err
	}

	if !u.Active {
		return domain.Category{}, pkg.InactiveUserError
	}

	cat := domain.Category{
		Name:        r.Name,
		Description: &r.Description,
		OwnerID:     userID,
		Users:       []domain.User{u},
		Notifiable:  r.Notifiable,
		Custom:      true,
	}
	c, err := uc.CategoryService.Create(ctx, cat)
	if err != nil {
		return domain.Category{}, err
	}

	return c, nil
}

func NewCreateCategoryUseCase(s category.Servicer, u user.Servicer, logger *log.Logger) *CreateCategoryUseCase {
	return &CreateCategoryUseCase{
		CategoryService: s,
		UserService:     u,
		logger:          logger,
	}
}
