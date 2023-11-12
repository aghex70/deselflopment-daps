package gorm

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string
	Description *string
	OwnerID     uint
	Users       *[]User `gorm:"many2many:daps_category_users"`
	Todos       *[]Todo
	Shared      bool
	Notifiable  bool
	Custom      bool
}

func (c Category) ToDto() domain.Category {
	return domain.Category{
		//ID:          c.ID,
		//CreatedAt:   c.CreatedAt,
		Name:        c.Name,
		Description: c.Description,
		OwnerID:     c.OwnerID,
		//Users:       c.Users,
		//Todos:       c.Todos,
		Shared:     c.Shared,
		Notifiable: c.Notifiable,
		Custom:     c.Custom,
	}
}

func (Category) TableName() string {
	return "daps_categories"
}

type GormCategoryRepository struct {
	*gorm.DB
}

func NewGormCategoryRepository(db *gorm.DB) *GormCategoryRepository {
	return &GormCategoryRepository{db}
}

func (gr *GormCategoryRepository) Get(ctx context.Context, id uint) (domain.Category, error) {
	var c Category
	result := gr.DB.First(&c, id)
	if result.Error != nil {
		return domain.Category{}, result.Error
	}
	return domain.Category{}, nil
}

func (gr *GormCategoryRepository) List(ctx context.Context, filters *map[string]interface{}) ([]domain.Category, error) {
	var cs []Category
	var cats []domain.Category

	if filters != nil {
		result := gr.DB.Where(filters).Find(&cs)
		if result.Error != nil {
			return []domain.Category{}, result.Error
		}
	} else {
		result := gr.DB.Find(&cs)
		if result.Error != nil {
			return []domain.Category{}, result.Error
		}
	}

	for _, c := range cs {
		cs := c.ToDto()
		cats = append(cats, cs)
	}
	return cats, nil
}

func (gr *GormCategoryRepository) ListByIds(ctx context.Context, ids []uint) ([]domain.Category, error) {
	var cs []Category
	var cats []domain.Category

	result := gr.DB.Find(&cs, ids)
	if result.Error != nil {
		return []domain.Category{}, result.Error
	}

	for _, c := range cs {
		cs := c.ToDto()
		cats = append(cats, cs)
	}
	return cats, nil
}

func (gr *GormCategoryRepository) Create(ctx context.Context, c domain.Category) (domain.Category, error) {
	result := gr.DB.Create(&c)
	if result.Error != nil {
		return domain.Category{}, result.Error
	}
	return c, nil
}

func (gr *GormCategoryRepository) Update(ctx context.Context, c domain.Category) error {
	return nil
}

func (gr *GormCategoryRepository) Delete(ctx context.Context, c domain.Category, email string) error {
	return nil
}
