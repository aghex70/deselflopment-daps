package category

import (
	"context"
	"database/sql"
	"github.com/aghex70/daps/internal/core/domain"
	"gorm.io/gorm"
	"log"
)

type CategoryGormRepository struct {
	*gorm.DB
	SqlDb  *sql.DB
	logger *log.Logger
}

func (gr *CategoryGormRepository) Delete(ctx context.Context, id uint) error {
	panic("foo")
}
func (gr *CategoryGormRepository) Get(ctx context.Context, userId uint) ([]domain.Category, error) {
	panic("foo")
}
func (gr *CategoryGormRepository) GetById(ctx context.Context, id uint, userId uint) (domain.Category, error) {
	panic("foo")
}
func (gr *CategoryGormRepository) Save(context.Context, domain.Category) error {
	panic("foo")
}

func NewCategoryGormRepository(db *gorm.DB) (*CategoryGormRepository, error) {
	return &CategoryGormRepository{
		DB: db,
	}, nil
}
