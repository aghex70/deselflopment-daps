package category

import (
	"context"
	"github.com/aghex70/daps/internal/core/domain"
	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/repositories/gorm/category"
	"log"
)

type CategoryService struct {
	logger             *log.Logger
	categoryRepository *category.CategoryGormRepository
}

func (s CategoryService) Create(context.Context, ports.CreateCategoryRequest) error {
	panic("foo")
}

func (s CategoryService) Delete(context.Context, ports.DeleteCategoryRequest) error {
	panic("foo")
}

func (s CategoryService) Get(context.Context, ports.GetCategoryRequest) (domain.Category, error) {
	panic("foo")
}

func (s CategoryService) List(context.Context, ports.ListCategoriesRequest) ([]domain.Category, error) {
	panic("foo")
}

func NewCategoryService(cr *category.CategoryGormRepository, logger *log.Logger) CategoryService {
	return CategoryService{
		logger:             logger,
		categoryRepository: cr,
	}
}
