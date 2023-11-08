package email

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
)

type Repository interface {
	Get(ctx context.Context, id uint) (domain.Email, error)
	List(ctx context.Context, filters *map[string]interface{}) ([]domain.Email, error)
	Create(ctx context.Context, e domain.Email) (domain.Email, error)
	Update(ctx context.Context, e domain.Email) error
	Delete(ctx context.Context, id uint) error
}
