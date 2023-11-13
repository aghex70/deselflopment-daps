package category

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
	requests "github.com/aghex70/daps/internal/ports/requests/category"
	"net/http"
)

type Servicer interface {
	Create(ctx context.Context, r *http.Request, req requests.CreateCategoryRequest) error
	Update(ctx context.Context, r *http.Request, req requests.UpdateCategoryRequest) error
	Get(ctx context.Context, r *http.Request, req requests.GetCategoryRequest) (domain.Category, error)
	Delete(ctx context.Context, r *http.Request, req requests.DeleteCategoryRequest) error
	List(ctx context.Context, r *http.Request) ([]domain.Category, error)
	GetOwner(ctx context.Context, id uint) (domain.User, error)
}
