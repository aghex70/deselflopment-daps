package category

import (
	"context"
	"fmt"
	"github.com/aghex70/daps/internal/core/domain"
	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/repositories/gorm/category"
	"github.com/aghex70/daps/internal/repositories/gorm/relationship"
	"github.com/aghex70/daps/server"
	"log"
	"net/http"
)

type CategoryService struct {
	logger                 *log.Logger
	categoryRepository     *category.CategoryGormRepository
	relationshipRepository *relationship.RelationshipGormRepository
}

func (s CategoryService) Create(ctx context.Context, r *http.Request, req ports.CreateCategoryRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)
	err := s.ValidateCreation(ctx, req.Name, int(userId))
	if err != nil {
		return err
	}

	cat := domain.Category{
		Description:       req.Description,
		Shared:            false,
		Custom:            true,
		Name:              req.Name,
		InternationalName: req.InternationalName,
	}
	nc, err := s.categoryRepository.Create(ctx, cat)
	if err != nil {
		return err
	}

	categoriesIds := []int{nc.ID}
	fmt.Println(userId)
	err = s.relationshipRepository.CreateRelationships(ctx, int(userId), categoriesIds)
	if err != nil {
		return err
	}

	return nil
}

func (s CategoryService) Get(ctx context.Context, r *http.Request, req ports.GetCategoryRequest) (domain.Category, error) {
	userId, _ := server.RetrieveJWTClaims(r, req)
	td, err := s.categoryRepository.GetById(ctx, int(req.CategoryId), int(userId))
	if err != nil {
		return domain.Category{}, err
	}
	return td, nil
}

func (s CategoryService) Update(ctx context.Context, r *http.Request, req ports.UpdateCategoryRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)
	cat := domain.Category{
		ID:                int(req.CategoryId),
		Description:       req.Description,
		Shared:            false,
		Name:              req.Name,
		InternationalName: req.InternationalName,
	}
	err := s.categoryRepository.Update(ctx, cat, int(userId))
	if err != nil {
		return err
	}
	return nil
}

func (s CategoryService) Delete(ctx context.Context, r *http.Request, req ports.DeleteCategoryRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)
	err := s.categoryRepository.Delete(ctx, int(req.CategoryId), int(userId))
	if err != nil {
		return err
	}
	return nil
}

func (s CategoryService) List(ctx context.Context, r *http.Request) ([]domain.Category, error) {
	userId, _ := server.RetrieveJWTClaims(r, nil)
	todos, err := s.categoryRepository.List(ctx, int(userId))
	if err != nil {
		return []domain.Category{}, err
	}
	return todos, nil
}

func NewCategoryService(cr *category.CategoryGormRepository, rr *relationship.RelationshipGormRepository, logger *log.Logger) CategoryService {
	return CategoryService{
		logger:                 logger,
		categoryRepository:     cr,
		relationshipRepository: rr,
	}
}
