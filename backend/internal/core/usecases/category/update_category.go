package category

import (
	"context"
	"github.com/aghex70/daps/internal/core/services/category"
	"github.com/aghex70/daps/internal/ports/domain"
	"github.com/aghex70/daps/server"
	"log"
)

type UpdateCategoryUseCase struct {
	CategoryService category.Service
	logger          *log.Logger
}

func (uc *UpdateCategoryUseCase) Execute(ctx context.Context, fields map[string]interface{}, id uint) ([]domain.Category, error) {
	userID, _ := server.RetrieveJWTClaims(r, req)

	switch {
	case req.Shared == nil:
		err := s.ValidateModification(ctx, uint(int(req.CategoryID)), uint(int(userID)))
		if err != nil {
			return err
		}
		cat := domain.Category{
			ID:          req.CategoryID,
			Description: req.Description,
			Notifiable:  req.Notifiable,
			Name:        req.Name,
		}
		err = s.repository.UpdateCategory(ctx, cat)
		if err != nil {
			return err
		}

	case *req.Shared:
		err := s.ValidateShare(ctx, uint(int(req.CategoryID)), uint(int(userID)))
		if err != nil {
			return err
		}
		cat := domain.Category{
			ID:     req.CategoryID,
			Shared: *req.Shared,
		}
		err = s.repository.Share(ctx, cat, req.Email)
		if err != nil {
			return err
		}

	case !*req.Shared:
		err := s.ValidateUnshare(ctx, uint(int(req.CategoryID)), uint(int(userID)))
		if err != nil {
			return err
		}
		cat := domain.Category{
			ID:     req.CategoryID,
			Shared: *req.Shared,
		}
		err = s.repository.Unshare(ctx, cat, int(userID))
		if err != nil {
			return err
		}
	}
}
