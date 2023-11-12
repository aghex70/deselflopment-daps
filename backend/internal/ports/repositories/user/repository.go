package user

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
)

type Repository interface {
	GetByEmail(ctx context.Context, email string) (domain.User, error)
	Create(ctx context.Context, u domain.User) (domain.User, error)
	Delete(ctx context.Context, id uint) error
	Activate(ctx context.Context, activationCode string) error
	Update(ctx context.Context, u domain.User) error
	Get(ctx context.Context, id uint) (domain.User, error)
	List(ctx context.Context, filters *map[string]interface{}) ([]domain.User, error)
}
