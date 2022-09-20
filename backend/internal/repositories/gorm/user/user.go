package user

import (
	"context"
	"database/sql"
	"github.com/aghex70/daps/internal/core/domain"
	"gorm.io/gorm"
	"log"
)

type UserGormRepository struct {
	*gorm.DB
	SqlDb  *sql.DB
	logger *log.Logger
}

func (gr *UserGormRepository) Delete(ctx context.Context, id uint) error {
	panic("foo")
}
func (gr *UserGormRepository) GetById(ctx context.Context, id uint) (domain.User, error) {
	panic("foo")
}
func (gr *UserGormRepository) Save(context.Context, domain.User) error {
	panic("foo")
}

func NewUserGormRepository(db *gorm.DB) (*UserGormRepository, error) {
	return &UserGormRepository{
		DB: db,
	}, nil
}
