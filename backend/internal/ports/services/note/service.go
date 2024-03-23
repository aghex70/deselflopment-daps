package note

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
)

type Servicer interface {
	Create(ctx context.Context, n domain.Note) (domain.Note, error)
	Get(ctx context.Context, id uint) (domain.Note, error)
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, filters *map[string]interface{}) ([]domain.Note, error)
	Update(ctx context.Context, id uint, filters *map[string]interface{}) error
	Share(ctx context.Context, id uint, u domain.User) error
	Unshare(ctx context.Context, id uint, u domain.User) error
}
