package relationship

import (
	"database/sql"
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

type Tabler interface {
	TableName() string
}

func (Category) TableName() string {
	return "daps_categories"
}

func (User) TableName() string {
	return "daps_users"
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
