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
		ID:          c.ID,
		CreatedAt:   c.CreatedAt,
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

func CategoryFromDto(c domain.Category) Category {
	return Category{
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

type CategoryRepository struct {
	*gorm.DB
}

func NewGormCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db}
}

func (gr *CategoryRepository) Create(ctx context.Context, c domain.Category) (domain.Category, error) {
	nc := CategoryFromDto(c)
	result := gr.DB.Create(&nc)
	if result.Error != nil {
		return domain.Category{}, result.Error
	}
	return nc.ToDto(), nil
}

func (gr *CategoryRepository) Get(ctx context.Context, id uint) (domain.Category, error) {
	var c Category
	result := gr.DB.First(&c, id)
	if result.Error != nil {
		return domain.Category{}, result.Error
	}
	return c.ToDto(), nil
}

func (gr *CategoryRepository) Delete(ctx context.Context, id uint) error {
	result := gr.DB.Delete(&Category{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *CategoryRepository) List(ctx context.Context, ids *[]uint, filters *map[string]interface{}) ([]domain.Category, error) {
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

func (gr *CategoryRepository) Update(ctx context.Context, id uint, c domain.Category) (domain.Category, error) {
	var cat Category
	result := gr.DB.First(&cat, id)
	if result.Error != nil {
		return domain.Category{}, result.Error
	}
	cat.Name = c.Name
	cat.Description = c.Description
	cat.OwnerID = c.OwnerID
	cat.Shared = c.Shared
	cat.Notifiable = c.Notifiable
	cat.Custom = c.Custom
	result = gr.DB.Save(&cat)
	if result.Error != nil {
		return domain.Category{}, result.Error
	}
	return cat.ToDto(), nil
}
