package user

import (
	"context"
	"database/sql"
	"github.com/aghex70/daps/internal/core/domain"
	"github.com/aghex70/daps/internal/repositories/gorm/relationship"
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
		return domain.User{}, result.Error
	}
	return nu.ToDto(), nil
}

func (gr *UserGormRepository) Delete(ctx context.Context, id int) error {
	u := relationship.User{ID: id}
	result := gr.DB.Where("owner_id = ?", id).Delete(&relationship.Category{})
	if result.Error != nil {
		return result.Error
	}
	result = gr.DB.Select(clause.Associations).Delete(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *UserGormRepository) Get(ctx context.Context, id int) (domain.User, error) {
	var u relationship.User
	result := gr.DB.Where(&relationship.User{ID: int(id)}).First(&u)
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

	return u.ToDto(), nil
}

func NewUserGormRepository(db *gorm.DB) (*UserGormRepository, error) {
	return &UserGormRepository{
		DB: db,
	}, nil
}