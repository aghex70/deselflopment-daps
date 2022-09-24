package ports

import (
	"context"
	"github.com/aghex70/daps/internal/core/domain"
	"net/http"
)

type CategoryServicer interface {
	Create(ctx context.Context, r CreateCategoryRequest) error
	Delete(ctx context.Context, r DeleteCategoryRequest) error
	Get(ctx context.Context, r GetCategoryRequest) (domain.Category, error)
	List(ctx context.Context, r ListCategoriesRequest) ([]domain.Category, error)
}

type TodoServicer interface {
	Create(ctx context.Context, r *http.Request, req CreateTodoRequest) error
	Complete(ctx context.Context, r *http.Request, req CompleteTodoRequest) error
	Delete(ctx context.Context, r *http.Request, req DeleteTodoRequest) error
	Get(ctx context.Context, r *http.Request, req GetTodoRequest) (domain.Todo, error)
	List(ctx context.Context, r *http.Request) ([]domain.Todo, error)
	Update(ctx context.Context, r *http.Request, req UpdateTodoRequest) error
	Start(ctx context.Context, r *http.Request, req StartTodoRequest) error
}

type UserServicer interface {
	Login(ctx context.Context, r LoginUserRequest) (string, error)
	Logout(ctx context.Context) error
	RefreshToken(ctx context.Context, r *http.Request) (string, error)
	Register(ctx context.Context, r CreateUserRequest) error
	Remove(ctx context.Context, r *http.Request) error
}
