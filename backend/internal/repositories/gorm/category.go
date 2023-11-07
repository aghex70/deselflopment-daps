package category

import (
	"context"
	"github.com/aghex70/daps/internal/core/domain"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name        string
	Description *string
	OwnerID     *uint
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

func (Category) TableName() string {
	return "daps_categories"
}

func (gr *GormRepository) GetCategory(ctx context.Context, id int) (domain.Category, error) {
	return domain.Category{}, nil
}

func (gr *GormRepository) GetCategories(ctx context.Context, ids []int) ([]domain.Category, error) {
	return nil, nil
}

func (gr *GormRepository) CreateCategory(ctx context.Context, c domain.Category, userId int) (domain.Category, error) {
	return domain.Category{}, nil
}

func (gr *GormRepository) UpdateCategory(ctx context.Context, c domain.Category) error {
	return nil
}

func (gr *GormRepository) DeleteCategory(ctx context.Context, c domain.Category, email string) error {
	return nil
}
