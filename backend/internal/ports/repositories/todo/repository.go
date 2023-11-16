package todo

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
)

type Repository interface {
	Create(ctx context.Context, t domain.Todo) (domain.Todo, error)
	Get(ctx context.Context, id uint) (domain.Todo, error)
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, ids *[]uint, filters *map[string]interface{}) ([]domain.Todo, error)
	Update(ctx context.Context, id uint, t domain.Todo) (domain.Todo, error)
}
