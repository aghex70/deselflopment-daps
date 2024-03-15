package topic

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
)

type Servicer interface {
	Create(ctx context.Context, t domain.Topic) (domain.Topic, error)
	Get(ctx context.Context, id uint) (domain.Topic, error)
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, filters *map[string]interface{}) ([]domain.Topic, error)
	Update(ctx context.Context, id uint, fields *map[string]interface{}) error
}
