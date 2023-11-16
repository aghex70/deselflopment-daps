package category

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
)

type Repository interface {
	Create(ctx context.Context, c domain.Category) (domain.Category, error)
	Get(ctx context.Context, id uint) (domain.Category, error)
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, ids *[]uint, filters *map[string]interface{}) ([]domain.Category, error)
	Update(ctx context.Context, id uint, c domain.Category) (domain.Category, error)
}
