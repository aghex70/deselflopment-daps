package category

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
	"github.com/aghex70/daps/internal/ports/repositories/category"
	"log"
)

type Service struct {
	logger             *log.Logger
	categoryRepository category.Repository
}

func (s Service) Create(ctx context.Context, c domain.Category) (domain.Category, error) {
	cat, err := s.categoryRepository.Create(ctx, c)
	if err != nil {
		return domain.Category{}, err
	}
	return cat, nil
}

func (s Service) Get(ctx context.Context, id uint) (domain.Category, error) {
	cat, err := s.categoryRepository.Get(ctx, id)
	if err != nil {
		return domain.Category{}, err
	}
	return cat, nil
}

func (s Service) Delete(ctx context.Context, id uint) error {
	err := s.categoryRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) List(ctx context.Context, ids *[]uint, fields *map[string]interface{}) ([]domain.Category, error) {
	categories, err := s.categoryRepository.List(ctx, ids, fields)
	if err != nil {
		return []domain.Category{}, err
	}
	return categories, nil
}

func (s Service) Update(ctx context.Context, id uint, fields *map[string]interface{}) error {
	err := s.categoryRepository.Update(ctx, id, fields)
	if err != nil {
		return err
	}
	return nil
}

func NewCategoryService(r category.Repository, logger *log.Logger) Service {
	return Service{
		logger:             logger,
		categoryRepository: r,
	}
}
