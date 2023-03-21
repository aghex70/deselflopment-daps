package category

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/aghex70/daps/internal/core/domain"
	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/repositories/gorm/category"
	"github.com/aghex70/daps/internal/repositories/gorm/email"
	"github.com/aghex70/daps/internal/repositories/gorm/relationship"
	"github.com/aghex70/daps/server"
)

type CategoryService struct {
	logger                 *log.Logger
	categoryRepository     *category.CategoryGormRepository
	relationshipRepository *relationship.RelationshipGormRepository
	emailRepository        *email.EmailGormRepository
}

func (s CategoryService) Create(ctx context.Context, r *http.Request, req ports.CreateCategoryRequest) error {
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
	_, err = s.categoryRepository.Create(ctx, cat, int(userId))
	if err != nil {
		return err
	}

	return nil
}

func (s CategoryService) Update(ctx context.Context, r *http.Request, req ports.UpdateCategoryRequest) error {
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
		err = s.categoryRepository.Update(ctx, cat)
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
		err = s.categoryRepository.Share(ctx, cat, req.Email)
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
		err = s.categoryRepository.Unshare(ctx, cat, int(userId))
		if err != nil {
			return err
		}
	}
	return nil
}

func (s CategoryService) Get(ctx context.Context, r *http.Request, req ports.GetCategoryRequest) (domain.Category, error) {
	userId, _ := server.RetrieveJWTClaims(r, req)
	err := s.ValidateRetrieval(ctx, int(req.CategoryId), int(userId))
	if err != nil {
		return domain.Category{}, err
	}
	cat, err := s.categoryRepository.GetById(ctx, int(req.CategoryId))
	if err != nil {
		return domain.Category{}, err
	}
	return cat, nil
}

func (s CategoryService) Delete(ctx context.Context, r *http.Request, req ports.DeleteCategoryRequest) error {
	q := r.URL.Query()
	categoryId, err := strconv.Atoi(q.Get("category_id"))
	if err != nil {
		return err
	}
	payload := ports.ListTodosRequest{}
	payload.Category = categoryId

	userId, _ := server.RetrieveJWTClaims(r, req)
	err = s.ValidateRemoval(ctx, int(req.CategoryId), int(userId))
	if err != nil {
		return err
	}
	err = s.categoryRepository.Delete(ctx, int(req.CategoryId), int(userId))
	if err != nil {
		return err
	}
	return nil
}

func (s CategoryService) List(ctx context.Context, r *http.Request) ([]domain.Category, error) {
	userId, _ := server.RetrieveJWTClaims(r, nil)
	categories, err := s.categoryRepository.List(ctx, int(userId))
	if err != nil {
		return []domain.Category{}, err
	}
	return categories, nil
}

func NewCategoryService(cr *category.CategoryGormRepository, rr *relationship.RelationshipGormRepository, er *email.EmailGormRepository, logger *log.Logger) CategoryService {
	return CategoryService{
		logger:                 logger,
		categoryRepository:     cr,
		relationshipRepository: rr,
		emailRepository:        er,
	}
}
