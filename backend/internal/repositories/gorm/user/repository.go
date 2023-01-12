package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/aghex70/daps/internal/core/domain"
	"github.com/aghex70/daps/internal/repositories/gorm/relationship"
	"github.com/aghex70/daps/pkg"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

type UserGormRepository struct {
	*gorm.DB
	SqlDb  *sql.DB
	logger *log.Logger
}

type Tabler interface {
	TableName() string
}

func (gr *UserGormRepository) Create(ctx context.Context, u domain.User) (domain.User, error) {
	nu := relationship.UserFromDto(u)
	result := gr.DB.Omit("Categories").Create(&nu)
	if result.Error != nil {
		fmt.Println("result.Error", result.Error)
		return domain.User{}, result.Error
	}
	return nu.ToDto(), nil
}

func (gr *UserGormRepository) Delete(ctx context.Context, adminId, id int) error {
	if adminId == id {
		return errors.New("admin user cannot be deleted")
	}
	type empty struct{}
	var Empty empty
	result := gr.DB.Raw("DELETE FROM daps_user_configs WHERE user_id = ?", id).Scan(&Empty)
	if result.Error != nil {
		return result.Error
	}

	var categoriesList []int
	result = gr.DB.Raw("SELECT id FROM daps_categories WHERE owner_id = ?", id).Scan(&categoriesList)
	if result.Error != nil {
		return result.Error
	}

	result = gr.DB.Raw("DELETE FROM daps_todos WHERE category_id IN ?", categoriesList).Scan(&Empty)
	if result.Error != nil {
		return result.Error
	}

	result = gr.DB.Raw("DELETE FROM daps_category_users WHERE user_id = ?", id).Scan(&Empty)
	if result.Error != nil {
		return result.Error
	}

	result = gr.DB.Raw("DELETE FROM daps_categories WHERE owner_id = ?", id).Scan(&Empty)
	if result.Error != nil {
		return result.Error
	}

	u := relationship.User{Id: id}
	result = gr.DB.Select(clause.Associations).Delete(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *UserGormRepository) Get(ctx context.Context, id int) (domain.User, error) {
	var u relationship.User
	result := gr.DB.Where(&relationship.User{Id: id}).First(&u)
	if result.Error != nil {
		return domain.User{}, result.Error
	}

	return u.ToDto(), nil
}

func (gr *UserGormRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	var u relationship.User
	result := gr.DB.Where(&relationship.User{Email: email}).First(&u)
	if result.Error != nil {
		return domain.User{}, result.Error
	}

	if u.Active == false {
		return domain.User{}, errors.New("user is not activated")
	}
	return u.ToDto(), nil
}

func (gr *UserGormRepository) ActivateUser(ctx context.Context, code string) error {
	var nu relationship.User
	var u relationship.User
	result := gr.DB.Where(&relationship.User{ActivationCode: code}).First(&u)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	if u.Active {
		return errors.New("user is already active")
	}

	result = gr.DB.Model(&nu).Where(relationship.User{ActivationCode: code}).Updates(map[string]interface{}{
		"active": true,
	})

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *UserGormRepository) CreateResetLink(ctx context.Context, email string) (domain.User, error) {
	var u relationship.User
	result := gr.DB.Where(&relationship.User{Email: email}).First(&u)
	if result.RowsAffected == 0 {
		return domain.User{}, gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return domain.User{}, result.Error
	}

	return u.ToDto(), nil
}

func (gr *UserGormRepository) ResetPassword(ctx context.Context, password, code string) error {
	var u relationship.User
	result := gr.DB.Where(&relationship.User{ResetPasswordCode: code}).First(&u)
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return result.Error
	}

	var nu relationship.User
	newUUID := pkg.GenerateUUID()
	result = gr.DB.Model(&nu).Where(relationship.User{ResetPasswordCode: code}).Updates(map[string]interface{}{
		"reset_password_code": newUUID,
		"password":            password})

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (gr *UserGormRepository) ProvisionDemoUser(ctx context.Context, e string) (domain.User, error) {
	nu := relationship.User{
		Name:     "Demo user",
		Email:    e,
		IsAdmin:  false,
		Password: "demopassword123",
	}
	result := gr.DB.Omit("Categories").Create(&nu)

	if result.Error != nil {
		return domain.User{}, result.Error
	}

	return nu.ToDto(), nil
}

func (gr *UserGormRepository) List(ctx context.Context) ([]domain.User, error) {
	var dbUsers []relationship.User
	var users []domain.User
	result := gr.DB.Find(&dbUsers)
	if result.Error != nil {
		return []domain.User{}, result.Error
	}

	for _, u := range dbUsers {
		cs := u.ToDto()
		users = append(users, cs)
	}
	return users, nil
}

func NewUserGormRepository(db *gorm.DB) (*UserGormRepository, error) {
	return &UserGormRepository{
		DB: db,
	}, nil
}
