package user

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
)

type Repository interface {
	Get(ctx context.Context, id uint) (domain.User, error)
	GetByEmail(ctx context.Context, email string) (domain.User, error)
	List(ctx context.Context, filters *map[string]interface{}) ([]domain.User, error)
	Create(ctx context.Context, u domain.User) (domain.User, error)
	Activate(ctx context.Context, id uint, activationCode string) error
	Update(ctx context.Context, id uint, filters *map[string]interface{}) error
	Delete(ctx context.Context, id uint) error
	ResetPassword(ctx context.Context, id uint, password, resetPasswordCode string) error
}
