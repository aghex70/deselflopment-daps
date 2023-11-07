package category

import (
	"github.com/aghex70/daps/internal/core/domain"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name              string
	Email             string
	Password          string
	Admin             bool
	Active            bool
	ActivationCode    string
	ResetPasswordCode string
	Language          string
	AutoSuggest       bool
	Categories        *[]Category `gorm:"many2many:daps_category_users"`
	Emails            *[]Email
	Todos             *[]Todo
	OwnedCategories   *[]Category
}

func (u User) ToDto() domain.User {
	return domain.User{
		ID:                u.ID,
		CreatedAt:         u.CreatedAt,
		Name:              u.Name,
		Email:             u.Email,
		Password:          u.Password,
		Admin:             u.Admin,
		Active:            u.Active,
		ActivationCode:    u.ActivationCode,
		ResetPasswordCode: u.ResetPasswordCode,
		Language:          u.Language,
		AutoSuggest:       u.AutoSuggest,
		//Categories:        u.Categories,
		//Emails:            u.Emails,
		//Todos:             u.Todos,
		//OwnedCategories:   u.OwnedCategories,
	}
}

func (User) TableName() string {
	return "deselflopment_users"
}

func (gr *GormRepository) GetUser(id int) (domain.User, error) {
	return domain.User{}, nil
}

func (gr *GormRepository) GetUsers() ([]domain.User, error) {
	return []domain.User{}, nil
}

func (gr *GormRepository) CreateUser(u User) (domain.User, error) {
	return domain.User{}, nil
}

func (gr *GormRepository) UpdateUser(u User) error {
	return nil
}

func (gr *GormRepository) DeleteUser(id int) error {
	return nil
}
