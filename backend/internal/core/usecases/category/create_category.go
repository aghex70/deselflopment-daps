package category

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
	requests "github.com/aghex70/daps/internal/ports/requests/category"
	"github.com/aghex70/daps/internal/ports/services/category"
	"log"
)

type CreateCategoryUseCase struct {
	CategoryService category.Servicer
	logger          *log.Logger
}

func (uc *CreateCategoryUseCase) Execute(ctx context.Context, userID uint, r requests.CreateCategoryRequest) (domain.Category, error) {
	u := domain.User{ID: userID}
	cat := domain.Category{
		Name:        r.Name,
		Description: &r.Description,
		OwnerID:     userID,
		Users:       &[]domain.User{u},
		Notifiable:  true,
		Custom:      true,
	}
	c, err := uc.CategoryService.Create(ctx, cat)
	if err != nil {
		return domain.Category{}, err
	}

	return c, nil
}

func NewCreateCategoryUseCase(s category.Servicer, logger *log.Logger) *CreateCategoryUseCase {
	return &CreateCategoryUseCase{
		CategoryService: s,
		logger:          logger,
	}
}
