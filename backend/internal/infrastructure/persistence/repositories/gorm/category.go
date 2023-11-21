package gorm

import (
	"context"
	"fmt"
	"github.com/aghex70/daps/internal/ports/domain"
	"gorm.io/gorm"
	"strings"
	"time"
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
	var createdAt time.Time
	if !c.CreatedAt.IsZero() {
		createdAt = c.CreatedAt
	}

	var todos []domain.Todo
	if c.Todos != nil {
		for _, todo := range *c.Todos {
			todos = append(todos, todo.ToDto())
		}
	}

	var users []domain.User
	if c.Users != nil {
		for _, user := range *c.Users {
			users = append(users, user.ToDto())
		}
	}

	return domain.Category{
		ID:          c.ID,
		CreatedAt:   createdAt,
		Name:        c.Name,
		Description: c.Description,
		OwnerID:     c.OwnerID,
		Users:       &users,
		Todos:       &todos,
		Shared:      c.Shared,
		Notifiable:  c.Notifiable,
		Custom:      c.Custom,
	}
}

func CategoryFromDto(c domain.Category) Category {
	var users []User
	if c.Users != nil {
		for _, userDTO := range *c.Users {
			user := UserFromDto(userDTO)
			users = append(users, user)
		}
	}

	var todos []Todo
	if c.Todos != nil {
		for _, todoDTO := range *c.Todos {
			todo := TodoFromDto(todoDTO)
			todos = append(todos, todo)
		}
	}

	return Category{
		Name:        c.Name,
		Description: c.Description,
		OwnerID:     c.OwnerID,
		Users:       &users,
		Todos:       &todos,
		Shared:      c.Shared,
		Notifiable:  c.Notifiable,
		Custom:      c.Custom,
	}
}

func (Category) TableName() string {
	return "daps_categories"
}

func (gr *CategoryRepository) Create(ctx context.Context, c domain.Category) (domain.Category, error) {
	nc := CategoryFromDto(c)
	result := gr.DB.Create(&nc)
	if result.Error != nil {
		return domain.Category{}, result.Error
	}

	// Hack to get around the fact that GORM doesn't support many-to-many relationships
	if nc.Users != nil {
		err := gr.DB.Association("Users").Append(*nc.Users)
		if err != nil {
			return domain.Category{}, err
		}
	}
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

	query := gr.DB
	if filters != nil {
		// Convert map[string]interface{} to a slice of arguments
		var args []interface{}
		var conditions []string
		for key, value := range *filters {
			conditions = append(conditions, fmt.Sprintf("%s = ?", key))
			args = append(args, value)
		}
		query = query.Where(strings.Join(conditions, " AND "), args...)
	}

	result := query.Find(&cs)
	if result.Error != nil {
		return []domain.Category{}, result.Error
	}

	for _, c := range cs {
		cs := c.ToDto()
		cats = append(cats, cs)
	}
	return cats, nil
}

func (gr *CategoryRepository) Update(ctx context.Context, id uint, filters *map[string]interface{}) error {
	var c Category
	c.ID = id
	result := gr.DB.Model(&c).Updates(*filters)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

type CategoryRepository struct {
	*gorm.DB
}

func NewGormCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db}
}
