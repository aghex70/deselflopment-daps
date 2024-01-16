package todo

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
	"mime/multipart"
)

type Servicer interface {
	Create(ctx context.Context, t domain.Todo) (domain.Todo, error)
	Get(ctx context.Context, id uint) (domain.Todo, error)
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, filters *map[string]interface{}) ([]domain.Todo, error)
	Update(ctx context.Context, id uint, fields *map[string]interface{}) error
	Start(ctx context.Context, id uint) error
	Import(ctx context.Context, f multipart.File) error
	//GetSummary(ctx context.Context, r *http.Request) ([]domain.Todo, error)
	//Suggest(ctx context.Context, r *http.Request) error
	//Remind(ctx context.Context) error
}
