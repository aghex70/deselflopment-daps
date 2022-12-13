package ports

import (
	"context"
	"github.com/aghex70/daps/internal/core/domain"
	"net/http"
)

type CategoryServicer interface {
	Create(ctx context.Context, r *http.Request, req CreateCategoryRequest) error
	Update(ctx context.Context, r *http.Request, req UpdateCategoryRequest) error
	Get(ctx context.Context, r *http.Request, req GetCategoryRequest) (domain.Category, error)
	Delete(ctx context.Context, r *http.Request, req DeleteCategoryRequest) error
	List(ctx context.Context, r *http.Request) ([]domain.Category, error)
}

type TodoServicer interface {
	Create(ctx context.Context, r *http.Request, req CreateTodoRequest) error
	Complete(ctx context.Context, r *http.Request, req CompleteTodoRequest) error
	Activate(ctx context.Context, r *http.Request, req ActivateTodoRequest) error
	Delete(ctx context.Context, r *http.Request, req DeleteTodoRequest) error
	Get(ctx context.Context, r *http.Request, req GetTodoRequest) (domain.TodoInfo, error)
	List(ctx context.Context, r *http.Request, req ListTodosRequest) ([]domain.Todo, error)
	ListCompleted(ctx context.Context, r *http.Request) ([]domain.Todo, error)
	ListRecurring(ctx context.Context, r *http.Request) ([]domain.Todo, error)
	Update(ctx context.Context, r *http.Request, req UpdateTodoRequest) error
	Start(ctx context.Context, r *http.Request, req StartTodoRequest) error
	Summary(ctx context.Context, r *http.Request) ([]domain.CategorySummary, error)
}

type UserServicer interface {
	Register(ctx context.Context, r CreateUserRequest) error
	Login(ctx context.Context, r LoginUserRequest) (string, int, error)
	RefreshToken(ctx context.Context, r *http.Request) (string, error)
	Remove(ctx context.Context, r *http.Request) error
	ProvisionDemoUser(ctx context.Context, r *http.Request, req ProvisionDemoUserRequest) error
}

type UserConfigServicer interface {
	Update(ctx context.Context, r *http.Request, req UpdateUserConfigRequest) error
	Get(ctx context.Context, r *http.Request) (domain.Profile, error)
}
