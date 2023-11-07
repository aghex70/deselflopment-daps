package category

import (
	"context"
	"log"
	"net/http"

	"github.com/aghex70/daps/internal/core/domain"
	"github.com/aghex70/daps/internal/core/ports"
	repository "github.com/aghex70/daps/internal/repositories/gorm"
	"github.com/aghex70/daps/server"
)

type Service struct {
	logger     *log.Logger
	repository *repository.GormRepository
}

func (s Service) Create(ctx context.Context, r *http.Request, req ports.CreateCategoryRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)
	err := s.ValidateCreation(ctx, req.Name, int(userId))
	if err != nil {
		return err
	}
	u := domain.User{Id: int(userId)}
	cat := domain.Category{
		OwnerId:     int(userId),
		Description: req.Description,
		Custom:      true,
		Name:        req.Name,
		Notifiable:  req.Notifiable,
		Users:       []domain.User{u},
	}
	_, err = s.repository.CreateCategory(ctx, cat, int(userId))
	if err != nil {
		return err
	}

	return nil
}

func (s Service) Update(ctx context.Context, r *http.Request, req ports.UpdateCategoryRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)

	switch {
	case req.Shared == nil:
		err := s.ValidateModification(ctx, int(req.CategoryId), int(userId))
		if err != nil {
			return err
		}
		cat := domain.Category{
			Id:          int(req.CategoryId),
			Description: req.Description,
			Notifiable:  req.Notifiable,
			Name:        req.Name,
		}
		err = s.repository.UpdateCategory(ctx, cat)
		if err != nil {
			return err
		}

	case *req.Shared:
		err := s.ValidateShare(ctx, int(req.CategoryId), int(userId))
		if err != nil {
			return err
		}
		cat := domain.Category{
			Id:     int(req.CategoryId),
			Shared: req.Shared,
		}
		err = s.repository.Share(ctx, cat, req.Email)
		if err != nil {
			return err
		}

	case !*req.Shared:
		err := s.ValidateUnshare(ctx, int(req.CategoryId), int(userId))
		if err != nil {
			return err
		}
		cat := domain.Category{
			Id:     int(req.CategoryId),
			Shared: req.Shared,
		}
		err = s.repository.Unshare(ctx, cat, int(userId))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s Service) Get(ctx context.Context, r *http.Request, req ports.GetCategoryRequest) (domain.Category, error) {
	userId, _ := server.RetrieveJWTClaims(r, req)
	err := s.ValidateRetrieval(ctx, int(req.CategoryId), int(userId))
	if err != nil {
		return domain.Category{}, err
	}
	cat, err := s.repository.GetCategory(ctx, int(req.CategoryId))
	if err != nil {
		return domain.Category{}, err
	}
	return cat, nil
}

func (s Service) Delete(ctx context.Context, r *http.Request, req ports.DeleteCategoryRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)
	err := s.ValidateRemoval(ctx, int(req.CategoryId), int(userId))
	if err != nil {
		return err
	}
	err = s.repository.DeleteCategory(ctx, int(req.CategoryId), int(userId))
	if err != nil {
		return err
	}
	return nil
}

func (s Service) List(ctx context.Context, r *http.Request) ([]domain.Category, error) {
	userId, _ := server.RetrieveJWTClaims(r, nil)
	categories, err := s.repository.GetCategories(ctx, int(userId))
	if err != nil {
		return []domain.Category{}, err
	}
	return categories, nil
}

func NewCategoryService(r *repository.GormRepository, logger *log.Logger) Service {
	return Service{
		logger:     logger,
		repository: r,
	}
}
