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

type Category struct {
	ID                int    `gorm:"primaryKey;column:id"`
	Custom            bool   `gorm:"column:custom"`
	Description       bool   `gorm:"column:description"`
	Name              string `gorm:"column:name"`
	InternationalName string `gorm:"column:international_name"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (Category) TableName() string {
	return "daps_todos"
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
