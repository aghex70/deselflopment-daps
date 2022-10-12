package category

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
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
	Shared            bool   `gorm:"column:shared"`
	Custom            bool   `gorm:"column:custom"`
	Description       string `gorm:"column:description"`
	Name              string `gorm:"column:name"`
	InternationalName string `gorm:"column:international_name"`
	//Users             []user.User `gorm:"many2many:daps_category_users"`
}

type Tabler interface {
	TableName() string
}

func (Category) TableName() string {
	return "daps_categories"
}

func (gr *CategoryGormRepository) GetByIds(ctx context.Context, ids []int) ([]domain.Category, error) {
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

func (gr *CategoryGormRepository) GetUserCategory(ctx context.Context, name string, userId int) (domain.Category, error) {
	var c Category
	query := fmt.Sprintf("SELECT daps_categories.id FROM daps_categories INNER JOIN daps_categories_users_relationships ON daps_categories.id = daps_categories_users_relationships.category_id WHERE daps_categories_users_relationships.user_id = %d AND daps_categories.name = '%s'", userId, name)
	fmt.Println(query)
	result := gr.DB.Raw(query).Scan(&c)

	if result.RowsAffected == 0 {
		fmt.Println("000000000000000000")
		return domain.Category{}, errors.New("category not updated")
	}

	fmt.Println("1111111111111111111")
	if result.Error != nil {
		fmt.Println("2222222222222222222222")
		return domain.Category{}, result.Error
	}
	fmt.Println("333333333333333333")
	return c.ToDto(), nil
}

func (gr *CategoryGormRepository) GetById(ctx context.Context, id int, userId int) (domain.Category, error) {
	var c Category
	query := fmt.Sprintf("SELECT * FROM daps_categories INNER JOIN daps_categories_users_relationships ON daps_categories.id = daps_categories_users_relationships.category_id WHERE daps_categories_users_relationships.user_id = %d AND daps_categories_users_relationships.category_id = %d", userId, id)
	fmt.Println(query)
	result := gr.DB.Raw(query).Scan(&c)
	if result.RowsAffected == 0 {
		fmt.Println("000000000000000000")
		return domain.Category{}, gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return domain.Category{}, result.Error
	}
	return c.ToDto(), nil
}

func (gr *CategoryGormRepository) Update(ctx context.Context, c domain.Category, userId int) error {
	var nc Category
	query := fmt.Sprintf("SELECT * FROM daps_categories INNER JOIN daps_categories_users_relationships ON daps_categories.id = daps_categories_users_relationships.category_id WHERE daps_categories_users_relationships.user_id = %d AND daps_categories_users_relationships.category_id = %d", userId, c.ID)

	//tx := db.Table("books").
	//	Joins("INNER JOIN user_liked_books ulb ON ulb.book_id = books.id").
	//	Select("books.id, books.name, count(ulb.user_id) as likes_count").
	//	Group("books.id, books.name").
	//	Order("likes_count desc").
	//	Limit(50).
	//	Find(&books)
	fmt.Println(query)
	result := gr.DB.Raw(query).Scan(&nc)
	if result.RowsAffected == 0 {
		fmt.Println("000000000000000000")
		return gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return result.Error
	}

	//nc.Name = c.Name
	//nc.Description = c.Description
	//nc.InternationalName = c.InternationalName
	fmt.Printf("\n\n------------> nc %+v", nc)
	result = gr.DB.Model(&nc).Where(Category{ID: c.ID}).Updates(map[string]interface{}{
		"name":               c.Name,
		"international_name": c.InternationalName,
		"description":        c.Description,
	})

	if result.RowsAffected == 0 {
		fmt.Println("111111111111111111111111111111111")
		fmt.Println("111111111111111111111111111111111")
		fmt.Println("111111111111111111111111111111111")
		fmt.Println("111111111111111111111111111111111")
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

func (gr *CategoryGormRepository) GetBaseCategory(ctx context.Context, name string) (domain.Category, error) {
	var c Category
	result := gr.DB.Where(&Category{Name: name, Shared: false}).First(&c)
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
		ID:                c.ID,
		Description:       c.Description,
		Shared:            c.Shared,
		Custom:            c.Custom,
		Name:              c.Name,
		InternationalName: c.InternationalName,
	}
}

func fromDto(c domain.Category) Category {
	return Category{
		ID:                c.ID,
		Shared:            c.Shared,
		Custom:            c.Custom,
		Description:       c.Description,
		Name:              c.Name,
		InternationalName: c.InternationalName,
	}
}

func CategoryDomainDB(categories []domain.Category) []Category {
	var c []Category
	for _, category := range categories {
		nc := fromDto(category)
		c = append(c, nc)
	}
	return c
}

func CategoryDBDomain(categories []Category) []domain.Category {
	var c []domain.Category
	for _, category := range categories {
		nc := category.ToDto()
		c = append(c, nc)
	}
	return c
}
