package gorm

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
	"gorm.io/gorm"
	"log"
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
	OwnedCategories   *[]Category `gorm:"foreignKey:OwnerID"`
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

func UserFromDto(u domain.User) User {
	return User{
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

type GormUserRepository struct {
	*gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *GormUserRepository {
	return &GormUserRepository{DB: db}
}

func (gr *GormUserRepository) Get(ctx context.Context, id uint) (domain.User, error) {
	var u User
	result := gr.DB.First(&u, id)
	if result.Error != nil {
		return domain.User{}, result.Error
	}
	log.Printf("User: %+v", u)
	return u.ToDto(), nil
}

func (gr *GormUserRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	var u User
	result := gr.DB.Where("email = ?", email).First(&u)
	if result.Error != nil {
		return domain.User{}, result.Error
	}
	return u.ToDto(), nil
}

func (gr *GormUserRepository) List(ctx context.Context, filters *map[string]interface{}) ([]domain.User, error) {
	return []domain.User{}, nil
}

func (gr *GormUserRepository) Create(ctx context.Context, u domain.User) (domain.User, error) {
	nu := UserFromDto(u)
	result := gr.DB.Create(&nu)
	if result.Error != nil {
		return domain.User{}, result.Error
	}
	return u, nil
}

func (gr *GormUserRepository) Update(ctx context.Context, u domain.User) error {
	return nil
}

func (gr *GormUserRepository) Delete(ctx context.Context, uid uint) error {
	return nil
}

//func NewGormUserRepository(db *gorm.DB, sqlDb *sql.DB) UserRepository {
//	return &GormUserRepository{
//		DB:         db,
//		SqlDb:      sqlDb,
//		repository: &GormUserRepository{},
//	}
//}
