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
	if err := s.categoryRepository.Delete(ctx, id); err != nil {
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
	if err := s.categoryRepository.Update(ctx, id, fields); err != nil {
		return err
	}
	return nil
}

func (s Service) Share(ctx context.Context, id uint, u domain.User) error {
	if err := s.categoryRepository.Share(ctx, id, u); err != nil {
		return err
	}
	return nil
}

func (s Service) Unshare(ctx context.Context, id uint, u domain.User) error {
	if err := s.categoryRepository.Unshare(ctx, id, u); err != nil {
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
