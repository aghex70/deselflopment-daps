package category

import (
	"context"
	domain2 "github.com/aghex70/daps/internal/ports/domain"
	requests "github.com/aghex70/daps/internal/ports/requests/category"
	"net/http"
)

type Servicer interface {
	Create(ctx context.Context, r *http.Request, req requests.CreateCategoryRequest) error
	Update(ctx context.Context, r *http.Request, req requests.UpdateCategoryRequest) error
	Get(ctx context.Context, r *http.Request, req requests.GetCategoryRequest) (domain2.Category, error)
	Delete(ctx context.Context, r *http.Request, req requests.DeleteCategoryRequest) error
	List(ctx context.Context, r *http.Request) ([]domain2.Category, error)
}
