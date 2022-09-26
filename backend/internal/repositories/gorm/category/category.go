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
	UserId            *int   `gorm:"column:user_id"`
	Custom            bool   `gorm:"column:custom"`
	Description       string `gorm:"column:description"`
	Name              string `gorm:"column:name"`
	InternationalName string `gorm:"column:international_name"`
}

type Tabler interface {
	TableName() string
}

func (Category) TableName() string {
	return "daps_categories"
}

func (gr *CategoryGormRepository) Delete(ctx context.Context, id int, userId int) error {
	c := Category{ID: id, UserId: &userId}
	result := gr.DB.Delete(&c)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *CategoryGormRepository) ListCustom(ctx context.Context, userId int) ([]domain.Category, error) {
	var cs []Category
	var cats []domain.Category
	result := gr.DB.Where(&Category{UserId: &userId, Custom: true}).Find(&cs)
	if result.Error != nil {
		return []domain.Category{}, result.Error
	}

	for _, c := range cs {
		cs := c.ToDto()
		cats = append(cats, cs)
	}
	return cats, nil
}

func (gr *CategoryGormRepository) GetBaseCategory(ctx context.Context, name string) (domain.Category, error) {
	var c Category
	result := gr.DB.Where(&Category{Name: name, Custom: false}).First(&c)
	if result.Error != nil {
		return domain.Category{}, result.Error
	}
	return c.ToDto(), nil
}

func (gr *CategoryGormRepository) List(ctx context.Context, userId int) ([]domain.Category, error) {
	var cs []Category
	var cats []domain.Category
	result := gr.DB.Where(gr.DB.Where("user_id = ?", &userId).Where("custom = ?", true)).Or("custom = ?", false).Find(&cs)
	if result.Error != nil {
		return []domain.Category{}, result.Error
	}

	for _, c := range cs {
		cs := c.ToDto()
		cats = append(cats, cs)
	}
	return cats, nil
}

func (gr *CategoryGormRepository) GetById(ctx context.Context, id int) (domain.Category, error) {
	var c Category
	result := gr.DB.Where(&Category{ID: id}).First(&c)
	if result.Error != nil {
		return domain.Category{}, result.Error
	}
	return c.ToDto(), nil
}
func (gr *CategoryGormRepository) Create(ctx context.Context, c domain.Category) error {
	nc := fromDto(c)
	result := gr.DB.Create(&nc)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *CategoryGormRepository) Update(ctx context.Context, c domain.Category) error {
	nc := fromDto(c)
	result := gr.DB.Model(&nc).Where(Category{ID: nc.ID, Custom: true}).Updates(map[string]interface{}{
		"name":               nc.Name,
		"international_name": nc.InternationalName,
		"description":        nc.Description,
	})

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewCategoryGormRepository(db *gorm.DB) (*CategoryGormRepository, error) {
	return &CategoryGormRepository{
		DB: db,
	}, nil
}

func (c Category) ToDto() domain.Category {
	return domain.Category{
		ID:                c.ID,
		Description:       c.Description,
		Custom:            c.Custom,
		Name:              c.Name,
		InternationalName: c.InternationalName,
		User:              c.UserId,
	}
}

func fromDto(c domain.Category) Category {
	return Category{
		ID:                c.ID,
		Custom:            c.Custom,
		Description:       c.Description,
		Name:              c.Name,
		InternationalName: c.InternationalName,
		UserId:            c.User,
	}
}
