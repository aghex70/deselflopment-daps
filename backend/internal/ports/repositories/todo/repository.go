package todo

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
)

type Repository interface {
	Get(ctx context.Context, id uint) (domain.Todo, error)
	List(ctx context.Context, filters *map[string]interface{}) ([]domain.Todo, error)
	Create(ctx context.Context, t domain.Todo) (domain.Todo, error)
	Update(ctx context.Context, t domain.Todo) error
	Delete(ctx context.Context, id uint) error
}
