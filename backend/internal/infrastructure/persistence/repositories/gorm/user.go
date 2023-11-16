package gorm

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
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

type UserRepository struct {
	*gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (gr *UserRepository) Get(ctx context.Context, id uint) (domain.User, error) {
	var u User
	result := gr.DB.First(&u, id)
	if result.Error != nil {
		return domain.User{}, result.Error
	}
	return u.ToDto(), nil
}

func (gr *UserRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	var u User
	result := gr.DB.Where("email = ?", email).First(&u)
	if result.Error != nil {
		return domain.User{}, result.Error
	}
	return u.ToDto(), nil
}

func (gr *UserRepository) List(ctx context.Context, filters *map[string]interface{}) ([]domain.User, error) {
	return []domain.User{}, nil
}

func (gr *UserRepository) Create(ctx context.Context, u domain.User) (domain.User, error) {
	nu := UserFromDto(u)
	result := gr.DB.Create(&nu)
	if result.Error != nil {
		return domain.User{}, result.Error
	}
	return u, nil
}

func (gr *UserRepository) Activate(ctx context.Context, id uint, activationCode string) error {
	var u User
	result := gr.DB.Where("activation_code = ?", activationCode).First(&u)
	if result.Error != nil {
		return result.Error
	}

	if u.ID != id {
		return gorm.ErrRecordNotFound
	}

	u.Active = true
	result = gr.DB.Save(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *UserRepository) Update(ctx context.Context, u domain.User, filters *map[string]interface{}) (domain.User, error) {
	return domain.User{}, nil
}

func (gr *UserRepository) Delete(ctx context.Context, id uint) error {
	return nil
}

func (gr *UserRepository) ResetPassword(ctx context.Context, id uint, password, resetPasswordCode string) error {
	var u User
	result := gr.DB.Where("reset_password_code = ?", resetPasswordCode).First(&u)
	if result.Error != nil {
		return result.Error
	}

	if u.ID != id {
		return gorm.ErrRecordNotFound
	}

	u.Password = password
	result = gr.DB.Save(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

//func NewGormUserRepository(db *gorm.DB, sqlDb *sql.DB) UserRepository {
//	return &GormUserRepository{
//		DB:         db,
//		SqlDb:      sqlDb,
//		repository: &GormUserRepository{},
//	}
//}
