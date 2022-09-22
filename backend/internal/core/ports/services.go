package ports

import (
	"context"
	"github.com/aghex70/daps/internal/core/domain"
)

type CategoryServicer interface {
	Create(ctx context.Context, r CreateCategoryRequest) error
	Delete(ctx context.Context, r DeleteCategoryRequest) error
	Get(ctx context.Context, r GetCategoryRequest) (domain.Category, error)
	List(ctx context.Context, r ListCategoriesRequest) ([]domain.Category, error)
}

type TodoServicer interface {
	Create(ctx context.Context, r CreateTodoRequest) error
	Complete(ctx context.Context, r CompleteTodoRequest) error
	Delete(ctx context.Context, r DeleteTodoRequest) error
	Get(ctx context.Context, r GetTodoRequest) (domain.Todo, error)
	List(ctx context.Context, r ListTodosRequest) ([]domain.Todo, error)
}

type UserServicer interface {
	Login(ctx context.Context, r LoginUserRequest) (string, error)
	Logout(ctx context.Context, r LogoutUserRequest) error
	RefreshToken(ctx context.Context, r RefreshTokenRequest) (string, error)
	Register(ctx context.Context, r CreateUserRequest) error
	Remove(ctx context.Context, r DeleteUserRequest) error
}
