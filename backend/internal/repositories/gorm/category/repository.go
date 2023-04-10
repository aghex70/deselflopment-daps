package category

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/aghex70/daps/internal/core/domain"
	"github.com/aghex70/daps/internal/repositories/gorm/relationship"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type GormRepository struct {
	*gorm.DB
	SqlDb *sql.DB
}

type Tabler interface {
	TableName() string
}

func (gr *GormRepository) GetByIds(ctx context.Context, ids []int) ([]domain.Category, error) {
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

func (gr *GormRepository) UserCategoryExists(ctx context.Context, conditions string) (int, error) {
	type queryResult struct {
		Id int `json:"id"`
	}
	var r queryResult
	result := gr.DB.Model(&relationship.Category{}).Select("daps_categories.id").Joins("INNER JOIN daps_category_users ON daps_categories.id = daps_category_users.category_id").Where(conditions).Find(&r)

	if result.RowsAffected == 0 {
		return r.Id, nil
	}

	if result.Error != nil {
		return r.Id, result.Error
	}
	return r.Id, nil
}

func (gr *GormRepository) Create(ctx context.Context, c domain.Category, userId int) (domain.Category, error) {
	nc := relationship.CategoryFromDto(c, userId)
	result := gr.DB.Create(&nc)
	if result.Error != nil {
		return domain.Category{}, result.Error
	}
	return nc.ToDto(), nil
}

func (gr *GormRepository) Update(ctx context.Context, c domain.Category) error {
	var nc relationship.Category
	result := gr.DB.Model(&nc).Where(relationship.Category{Id: c.Id}).Updates(map[string]interface{}{
		"name":        c.Name,
		"description": c.Description,
		"notifiable":  c.Notifiable,
	})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (gr *GormRepository) Share(ctx context.Context, c domain.Category, email string) error {
	type queryResult struct {
		Id int
	}
	var qr queryResult
	var nc relationship.Category
	query := fmt.Sprintf("SELECT deselflopment_users.id FROM deselflopment_users WHERE deselflopment_users.email = '%s'", email)
	result := gr.DB.Raw(query).Scan(&qr)
	if result.RowsAffected == 0 {
		return errors.New("invalid email")
	}

	if result.Error != nil {
		return result.Error
	}

	query = fmt.Sprintf("INSERT INTO daps_category_users (category_id, user_id) VALUES (%d, %d)", c.Id, qr.Id)
	result = gr.DB.Raw(query).Scan(&nc)
	if result.Error != nil {
		// 1062 - duplicate entry
		if result.Error.(*mysql.MySQLError).Number == 1062 {
			return errors.New("user already subscribed to that category")
		}

		return result.Error
	}

	result = gr.DB.Model(&nc).Where(relationship.Category{Id: c.Id}).Update("shared", c.Shared)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (gr *GormRepository) Unshare(ctx context.Context, c domain.Category, userId int) error {
	var cat relationship.Category
	result := gr.DB.Raw("DELETE FROM daps_category_users WHERE category_id = ? AND user_id = ?", c.Id, userId).Scan(&cat)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *GormRepository) GetById(ctx context.Context, id int) (domain.Category, error) {
	var c relationship.Category
	result := gr.DB.Where(&relationship.Category{Id: id}).First(&c)
	if result.RowsAffected == 0 {
		return domain.Category{}, gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return domain.Category{}, result.Error
	}
	return c.ToDto(), nil
}

func (gr *GormRepository) Delete(ctx context.Context, id int, userId int) error {
	var c relationship.Category
	result := gr.DB.Raw("DELETE FROM daps_todos WHERE category_id = ?", id).Scan(&c)
	if result.Error != nil {
		return result.Error
	}

	result = gr.DB.Where("id = ?", id).Delete(&c)
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

func (gr *GormRepository) List(ctx context.Context, userId int) ([]domain.Category, error) {
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

func NewGormRepository(db *gorm.DB) (*GormRepository, error) {
	return &GormRepository{
		DB: db,
	}, nil
}
