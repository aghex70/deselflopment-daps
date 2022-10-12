package user

import (
	"context"
	"database/sql"
	"github.com/aghex70/daps/internal/core/domain"
	"github.com/aghex70/daps/internal/repositories/gorm/category"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"time"
)

type UserGormRepository struct {
	*gorm.DB
	SqlDb  *sql.DB
	logger *log.Logger
}

type User struct {
	ID               int                 `gorm:"primaryKey;column:id"`
	Email            string              `gorm:"column:email"`
	IsAdmin          bool                `gorm:"column:is_admin"`
	Password         string              `gorm:"column:password"`
	RegistrationDate time.Time           `gorm:"column:registration_date;autoCreateTime"`
	Categories       []category.Category `gorm:"many2many:daps_category_users"`
}

type Tabler interface {
	TableName() string
}

func (User) TableName() string {
	return "daps_users"
}

func (gr *UserGormRepository) Create(ctx context.Context, u domain.User) (domain.User, error) {
	nu := fromDto(u)
	result := gr.DB.Create(&nu)
	if result.Error != nil {
		return domain.User{}, result.Error
	}
	return nu.ToDto(), nil
}

func (gr *UserGormRepository) Delete(ctx context.Context, id int) error {
	u := User{ID: id}
	result := gr.DB.Select(clause.Associations).Delete(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *UserGormRepository) Get(ctx context.Context, id int) (domain.User, error) {
	var u User
	result := gr.DB.Where(&User{ID: int(id)}).First(&u)
	if result.Error != nil {
		return domain.User{}, result.Error
	}

	return u.ToDto(), nil
}

func (gr *UserGormRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	var u User
	result := gr.DB.Where(&User{Email: email}).First(&u)
	if result.Error != nil {
		return domain.User{}, result.Error
	}

	return u.ToDto(), nil
}

func (u User) ToDto() domain.User {
	return domain.User{
		ID:         u.ID,
		Email:      u.Email,
		Categories: category.CategoryDBDomain(u.Categories),
	}
}

func fromDto(u domain.User) User {
	return User{
		Email:      u.Email,
		Password:   u.Password,
		Categories: category.CategoryDomainDB(u.Categories),
	}
}

func NewUserGormRepository(db *gorm.DB) (*UserGormRepository, error) {
	return &UserGormRepository{
		DB: db,
	}, nil
}
