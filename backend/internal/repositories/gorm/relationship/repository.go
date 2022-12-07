package relationship

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/aghex70/daps/internal/core/domain"
	"gorm.io/gorm"
	"log"
	"time"
)

type RelationshipGormRepository struct {
	*gorm.DB
	SqlDb  *sql.DB
	logger *log.Logger
}

type User struct {
	ID               int        `gorm:"primaryKey;column:id"`
	Email            string     `gorm:"column:email"`
	IsAdmin          bool       `gorm:"column:is_admin"`
	Password         string     `gorm:"column:password"`
	RegistrationDate time.Time  `gorm:"column:registration_date;autoCreateTime"`
	Categories       []Category `gorm:"many2many:daps_category_users"`
}

type Category struct {
	ID                int    `gorm:"primaryKey;column:id"`
	OwnerID           int    `gorm:"column:owner_id"`
	Shared            bool   `gorm:"column:shared"`
	Custom            bool   `gorm:"column:custom"`
	Description       string `gorm:"column:description"`
	Name              string `gorm:"column:name"`
	InternationalName string `gorm:"column:international_name"`
	Users             []User `gorm:"many2many:daps_category_users"`
}

type UserCategory struct {
	UserID     int `gorm:"column:user_id"`
	CategoryID int `gorm:"column:category_id"`
}

type Tabler interface {
	TableName() string
}

func (Category) TableName() string {
	return "daps_categories"
}

func (User) TableName() string {
	return "daps_users"
}

func (UserCategory) TableName() string {
	return "daps_category_users"
}

func (gr *RelationshipGormRepository) GetUserCategory(ctx context.Context, userId, categoryId int) error {
	type queryResult struct {
		Id int
	}
	var qr queryResult
	query := fmt.Sprintf("SELECT daps_categories.id FROM daps_categories INNER JOIN daps_category_users ON daps_categories.id = daps_category_users.category_id INNER JOIN daps_users ON daps_users.id = daps_category_users.user_id WHERE daps_category_users.user_id = %d AND daps_category_users.category_id = %d", userId, categoryId)
	result := gr.DB.Raw(query).Scan(&qr)

	if result.RowsAffected == 0 {
		return errors.New("user not linked to category")
	}

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *RelationshipGormRepository) ListUserCategories(ctx context.Context, userId int) ([]int, error) {
	var uc []UserCategory
	result := gr.DB.Where(&UserCategory{UserID: userId}).Find(&uc)

	if result.RowsAffected == 0 {
		return []int{}, errors.New("user not linked to category")
	}

	if result.Error != nil {
		return []int{}, result.Error
	}

	categoryIds := UserCategoryToList(uc)
	return categoryIds, nil
}

func NewRelationshipGormRepository(db *gorm.DB) (*RelationshipGormRepository, error) {
	return &RelationshipGormRepository{
		DB: db,
	}, nil
}

func (c Category) ToDto() domain.Category {
	return domain.Category{
		ID:                c.ID,
		OwnerID:           c.OwnerID,
		Description:       c.Description,
		Shared:            &c.Shared,
		Custom:            c.Custom,
		Name:              c.Name,
		InternationalName: c.InternationalName,
	}
}

func CategoryFromDto(c domain.Category, userId int) Category {
	return Category{
		ID:                c.ID,
		OwnerID:           c.OwnerID,
		Custom:            c.Custom,
		Description:       c.Description,
		Name:              c.Name,
		InternationalName: c.InternationalName,
		Users:             []User{{ID: userId}},
	}
}

func (u User) ToDto() domain.User {
	return domain.User{
		ID:         u.ID,
		Email:      u.Email,
		Categories: CategoryDBDomain(u.Categories),
		Password:   u.Password,
	}
}

func UserFromDto(u domain.User) User {
	return User{
		Email:    u.Email,
		Password: u.Password,
		//Categories: CategoryDomainDB(u.Categories, 0),
	}
}

func CategoryDomainDB(categories []domain.Category, userId int) []Category {
	var c []Category
	for _, category := range categories {
		nc := CategoryFromDto(category, userId)
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

func UserCategoryToList(userCategories []UserCategory) []int {
	var c []int
	for _, uc := range userCategories {
		c = append(c, uc.CategoryID)
	}
	return c
}
