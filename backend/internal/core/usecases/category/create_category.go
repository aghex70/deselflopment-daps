package category

import (
	"context"
	"github.com/aghex70/daps/internal/core/services/category"
	"github.com/aghex70/daps/internal/ports/domain"
	"github.com/aghex70/daps/server"
	"log"
)

type CreateCategoryUseCase struct {
	CategoryService category.Service
	logger          *log.Logger
}

func (uc *CreateCategoryUseCase) Execute(ctx context.Context, c domain.Category) (domain.Category, error) {
	userID, _ := server.RetrieveJWTClaims(r, req)
	u := domain.User{ID: userID}
	cat := domain.Category{
		OwnerID: userID,
		Name:    c.Name,
		//Description: req.Description,
		Custom:     true,
		Notifiable: c.Notifiable,
		Users:      &[]domain.User{u},
	}
	err := uc.CategoryService.Create(ctx, cat)
	if err != nil {
		return domain.Category{}, err
	}

	return domain.Category{}, nil
}

func NewCreateCategoryUseCase(s category.Service, logger *log.Logger) *CreateCategoryUseCase {
	return &CreateCategoryUseCase{
		CategoryService: s,
		logger:          logger,
	}
}
