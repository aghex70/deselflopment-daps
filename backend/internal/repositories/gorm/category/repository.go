package category

import (
	"context"
	"database/sql"
	"errors"
	"github.com/aghex70/daps/internal/core/domain"
	"github.com/aghex70/daps/internal/repositories/gorm/relationship"
	"gorm.io/gorm"
	"log"
)

type CategoryGormRepository struct {
	*gorm.DB
	SqlDb  *sql.DB
	logger *log.Logger
}

type Tabler interface {
	TableName() string
}

func (gr *CategoryGormRepository) GetByIds(ctx context.Context, ids []int) ([]domain.Category, error) {
	var cs []relationship.Category
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

func (gr *CategoryGormRepository) UserCategoryExists(ctx context.Context, conditions string) (int, error) {
	type queryResult struct {
		ID int `json:"id"`
	}
	var r queryResult
	result := gr.DB.Model(&relationship.Category{}).Select("daps_categories.id").Joins("INNER JOIN daps_category_users ON daps_categories.id = daps_category_users.category_id").Where(conditions).Find(&r)

	if result.RowsAffected == 0 {
		return r.ID, nil
	}

	if result.Error != nil {
		return r.ID, result.Error
	}
	return r.ID, nil
}

func (gr *CategoryGormRepository) Create(ctx context.Context, c domain.Category, userId int) error {
	nc := relationship.CategoryFromDto(c, userId)
	result := gr.DB.Create(&nc)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *CategoryGormRepository) Update(ctx context.Context, c domain.Category) error {
	var nc relationship.Category
	result := gr.DB.Model(&nc).Where(relationship.Category{ID: c.ID}).Updates(map[string]interface{}{
		"name":               c.Name,
		"international_name": c.InternationalName,
		"description":        c.Description,
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (gr *CategoryGormRepository) Share(ctx context.Context, c domain.Category, email string) error {
	var nc relationship.Category
	result := gr.DB.Model(&nc).Where(relationship.Category{ID: c.ID}).Update("shared", c.Shared)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (gr *CategoryGormRepository) GetById(ctx context.Context, id int) (domain.Category, error) {
	var c relationship.Category
	result := gr.DB.Where(&relationship.Category{ID: id}).First(&c)
	if result.RowsAffected == 0 {
		return domain.Category{}, gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return domain.Category{}, result.Error
	}
	return c.ToDto(), nil
}

func (gr *CategoryGormRepository) Delete(ctx context.Context, id int, userId int) error {
	var c relationship.Category
	result := gr.DB.Where("id = ?", id).Delete(&c)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("cannot delete category")
	}

	result = gr.DB.Raw("DELETE FROM daps_category_users WHERE category_id = ? AND user_id = ?", id, userId).Scan(&c)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (gr *CategoryGormRepository) ListCustom(ctx context.Context, userId int) ([]domain.Category, error) {
	var cs []relationship.Category
	var cats []domain.Category
	result := gr.DB.Where(&relationship.Category{Shared: true}).Find(&cs)
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
	var c relationship.Category
	result := gr.DB.Where(&relationship.Category{Name: name, Shared: false}).First(&c)
	if result.Error != nil {
		return domain.Category{}, result.Error
	}
	return c.ToDto(), nil
}

func (gr *CategoryGormRepository) List(ctx context.Context, userId int) ([]domain.Category, error) {
	var cs []relationship.Category
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

func NewCategoryGormRepository(db *gorm.DB) (*CategoryGormRepository, error) {
	return &CategoryGormRepository{
		DB: db,
	}, nil
}
