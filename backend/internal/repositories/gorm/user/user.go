package user

import (
	"context"
	"database/sql"
	"github.com/aghex70/daps/internal/core/domain"
	"gorm.io/gorm"
	"log"
	"time"
)

type UserGormRepository struct {
	*gorm.DB
	SqlDb  *sql.DB
	logger *log.Logger
}

type User struct {
	ID               int       `gorm:"primaryKey;column:id"`
	Email            string    `gorm:"column:email"`
	Password         string    `gorm:"column:password"`
	AccessToken      string    `gorm:"column:access_token"`
	RefreshToken     string    `gorm:"column:refresh_token"`
	RegistrationDate time.Time `gorm:"column:registration_date;autoCreateTime"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (User) TableName() string {
	return "daps_users"
}

func (gr *UserGormRepository) Create(ctx context.Context, u domain.User) error {
	nu := fromDto(u)
	result := gr.DB.Create(&nu)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *UserGormRepository) Delete(ctx context.Context, id uint) error {
	panic("foo")
}
func (gr *UserGormRepository) GetById(ctx context.Context, id uint) (domain.User, error) {
	panic("foo")
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
		ID:           u.ID,
		Email:        u.Email,
		AccessToken:  u.AccessToken,
		RefreshToken: u.RefreshToken,
	}
}

func fromDto(u domain.User) User {
	return User{
		Email:        u.Email,
		Password:     u.Password,
		AccessToken:  u.AccessToken,
		RefreshToken: u.RefreshToken,
	}
}

func NewUserGormRepository(db *gorm.DB) (*UserGormRepository, error) {
	return &UserGormRepository{
		DB: db,
	}, nil
}
