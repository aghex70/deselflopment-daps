package usercategory

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/aghex70/daps/internal/core/domain"
	"gorm.io/gorm"
)

type CategoryGormRepository struct {
	*gorm.DB
	SqlDb *sql.DB
}

type Category struct {
	Id          int    `gorm:"primaryKey;column:id"`
	CategoryId  int    `gorm:"column:category_id"`
	UserId      int    `gorm:"column:user_id"`
	Shared      bool   `gorm:"column:shared"`
	Custom      bool   `gorm:"column:custom"`
	Description string `gorm:"column:description"`
	Name        string `gorm:"column:name"`
}

type Tabler interface {
	TableName() string
}

func (Category) TableName() string {
	return "daps_categories"
}

func (gr *CategoryGormRepository) GetUserCategory(ctx context.Context, name string, userId int) (domain.Category, error) {
	var c Category
	query := fmt.Sprintf("SELECT daps_categories.id FROM daps_categories INNER JOIN daps_categories_users_relationships ON daps_categories.id = daps_categories_users_relationships.category_id WHERE daps_categories_users_relationships.user_id = %d AND daps_categories.name = '%s'", userId, name)
	result := gr.DB.Raw(query).Scan(&c)

	if result.RowsAffected == 0 {
		return domain.Category{}, errors.New("category not updated")
	}

	if result.Error != nil {
		return domain.Category{}, result.Error
	}
	return c.ToDto(), nil
}

func (gr *CategoryGormRepository) GetById(ctx context.Context, id int, userId int) (domain.Category, error) {
	var c Category
	query := fmt.Sprintf("SELECT * FROM daps_categories INNER JOIN daps_categories_users_relationships ON daps_categories.id = daps_categories_users_relationships.category_id WHERE daps_categories_users_relationships.user_id = %d AND daps_categories_users_relationships.category_id = %d", userId, id)
	result := gr.DB.Raw(query).Scan(&c)
	if result.RowsAffected == 0 {
		return domain.Category{}, gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return domain.Category{}, result.Error
	}
	return c.ToDto(), nil
}

func (gr *CategoryGormRepository) Update(ctx context.Context, c domain.Category, userId int) error {
	var nc Category
	query := fmt.Sprintf("SELECT * FROM daps_categories INNER JOIN daps_categories_users_relationships ON daps_categories.id = daps_categories_users_relationships.category_id WHERE daps_categories_users_relationships.user_id = %d AND daps_categories_users_relationships.category_id = %d", userId, c.Id)

	result := gr.DB.Raw(query).Scan(&nc)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return result.Error
	}

	result = gr.DB.Model(&nc).Where(Category{Id: c.Id}).Updates(map[string]interface{}{
		"name":        c.Name,
		"description": c.Description,
	})

	if result.RowsAffected == 0 {
		return errors.New("category not updated")
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (gr *CategoryGormRepository) Delete(ctx context.Context, id int, userId int) error {
	var c Category
	result := gr.DB.Where("id = ?", id).Where("user_id = ?", userId).Delete(&c)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("cannot delete category")
	}
	return nil
}

func (gr *CategoryGormRepository) ListCustom(ctx context.Context, userId int) ([]domain.Category, error) {
	var cs []Category
	var cats []domain.Category
	result := gr.DB.Where(&Category{Shared: true}).Find(&cs)
	if result.Error != nil {
		return []domain.Category{}, result.Error
	}

	for _, c := range cs {
		cs := c.ToDto()
		cats = append(cats, cs)
	}
	return cats, nil
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

func (gr *CategoryGormRepository) Create(ctx context.Context, c domain.Category) (domain.Category, error) {
	nc := fromDto(c)
	result := gr.DB.Create(&nc)
	if result.Error != nil {
		return domain.Category{}, result.Error
	}
	return nc.ToDto(), nil
}

func NewCategoryGormRepository(db *gorm.DB) (*CategoryGormRepository, error) {
	return &CategoryGormRepository{
		DB: db,
	}, nil
}

func (c Category) ToDto() domain.Category {
	return domain.Category{
		Id:          c.Id,
		Description: c.Description,
		Shared:      &c.Shared,
		Custom:      c.Custom,
		Name:        c.Name,
	}
}

func fromDto(c domain.Category) Category {
	return Category{
		Id:          c.Id,
		Shared:      *c.Shared,
		Custom:      c.Custom,
		Description: c.Description,
		Name:        c.Name,
	}
}
