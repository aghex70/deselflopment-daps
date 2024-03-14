package gorm

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
	common "github.com/aghex70/daps/utils"
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
	Categories        []Category `gorm:"many2many:daps_category_users;save_association:true"`
	Emails            []Email
	//Todos             []Todo
	OwnedCategories []Category `gorm:"foreignKey:OwnerID"`
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
	user := User{
		Name:              u.Name,
		Email:             u.Email,
		Password:          u.Password,
		Admin:             u.Admin,
		Active:            u.Active,
		ActivationCode:    u.ActivationCode,
		ResetPasswordCode: u.ResetPasswordCode,
		Language:          u.Language,
		AutoSuggest:       u.AutoSuggest,
		//Categories:        &categories,
		//Emails:            &emails,
		//Todos:             &todos,
		//OwnedCategories:   &ownedCategories,
	}
	user.ID = u.ID
	return user
}

func (User) TableName() string {
	return "deselflopment_users"
}

func (gr *UserRepository) Get(ctx context.Context, id uint) (domain.User, error) {
	var u User
	if result := gr.DB.First(&u, id); result.Error != nil {
		return domain.User{}, result.Error
	}
	return u.ToDto(), nil
}

func (gr *UserRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	var u User
	if result := gr.DB.Where("email = ?", email).First(&u); result.Error != nil {
		return domain.User{}, result.Error
	}
	return u.ToDto(), nil
}

func (gr *UserRepository) Create(ctx context.Context, u domain.User) (domain.User, error) {
	nu := UserFromDto(u)
	if result := gr.DB.Create(&nu); result.Error != nil {
		return domain.User{}, result.Error
	}
	return nu.ToDto(), nil
}

func (gr *UserRepository) Activate(ctx context.Context, id uint, activationCode string) error {
	var u User
	if result := gr.DB.Where("activation_code = ?", activationCode).First(&u); result.Error != nil {
		return result.Error
	}

	if u.ID != id {
		return gorm.ErrRecordNotFound
	}

	u.Active = true
	if result := gr.DB.Save(&u); result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *UserRepository) Update(ctx context.Context, id uint, filters *map[string]interface{}) error {
	var u User
	u.ID = id
	if result := gr.DB.Model(&u).Updates(*filters); result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *UserRepository) List(ctx context.Context, filters *map[string]interface{}) ([]domain.User, error) {
	var users []User
	if filters != nil {
		if result := gr.DB.Where(filters).Find(&users); result.Error != nil {
			return []domain.User{}, result.Error
		}
	} else {
		if result := gr.DB.Find(&users); result.Error != nil {
			return []domain.User{}, result.Error
		}
	}
	var us []domain.User
	for _, u := range users {
		us = append(us, u.ToDto())
	}
	return us, nil
}

func (gr *UserRepository) Delete(ctx context.Context, id uint) error {
	if result := gr.DB.Delete(&User{}, id); result.Error != nil {
		return result.Error
	}

	// Delete all category associations
	query := gr.DB.Exec(
		"DELETE FROM daps_category_users WHERE user_id = ?", id)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func (gr *UserRepository) ResetPassword(ctx context.Context, id uint, password, resetPasswordCode string) error {
	var u User
	if result := gr.DB.Where("reset_password_code = ?", resetPasswordCode).First(&u); result.Error != nil {
		return result.Error
	}

	if u.ID != id {
		return gorm.ErrRecordNotFound
	}

	u.Password = password
	u.ResetPasswordCode = common.GenerateUUID()
	if result := gr.DB.Save(&u); result.Error != nil {
		return result.Error
	}
	return nil
}

type UserRepository struct {
	*gorm.DB
}

func NewGormUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}
