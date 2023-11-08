package category

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
)

type Repository interface {
	Get(ctx context.Context, id uint) (domain.Category, error)
	List(ctx context.Context, filters *map[string]interface{}) ([]domain.Category, error)
	Create(ctx context.Context, c domain.Category) (domain.Category, error)
	Update(ctx context.Context, c domain.Category) error
	Delete(ctx context.Context, c domain.Category, email string) error
}
