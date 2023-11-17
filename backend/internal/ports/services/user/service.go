package user

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
)

type Servicer interface {
	GetByEmail(ctx context.Context, email string) (domain.User, error)
	Create(ctx context.Context, u domain.User) (domain.User, error)
	Delete(ctx context.Context, id uint) error
	Activate(ctx context.Context, id uint, activationCode string) error
	List(ctx context.Context, fields *map[string]interface{}) ([]domain.User, error)
	Get(ctx context.Context, id uint) (domain.User, error)
	ResetPassword(ctx context.Context, userID uint, password, resetPasswordCode string) error
	Update(ctx context.Context, id uint, fields *map[string]interface{}) error
}
