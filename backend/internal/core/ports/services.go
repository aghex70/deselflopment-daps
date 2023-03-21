package ports

import (
	"context"
	"mime/multipart"
	"net/http"

	"github.com/aghex70/daps/internal/core/domain"
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
	ListSuggested(ctx context.Context, r *http.Request) ([]domain.TodoInfo, error)
	Suggest(ctx context.Context, r *http.Request) error
	Update(ctx context.Context, r *http.Request, req UpdateTodoRequest) error
	Start(ctx context.Context, r *http.Request, req StartTodoRequest) error
	Restart(ctx context.Context, r *http.Request, req StartTodoRequest) error
	Summary(ctx context.Context, r *http.Request) ([]domain.CategorySummary, error)
	Remind(ctx context.Context) error
}

type UserServicer interface {
	Register(ctx context.Context, r CreateUserRequest) error
	Login(ctx context.Context, r LoginUserRequest) (string, int, error)
	RefreshToken(ctx context.Context, r *http.Request) (string, error)
	CheckAdmin(ctx context.Context, r *http.Request) (int, error)
	Delete(ctx context.Context, r *http.Request, req DeleteUserRequest) error
	Get(ctx context.Context, r *http.Request, req GetUserRequest) (domain.User, error)
	ProvisionDemoUser(ctx context.Context, r *http.Request, req ProvisionDemoUserRequest) error
	List(ctx context.Context, r *http.Request) ([]domain.User, error)
	ImportCSV(ctx context.Context, r *http.Request, f multipart.File) error
	Activate(ctx context.Context, r ActivateUserRequest) error
	SendResetLink(ctx context.Context, r ResetLinkRequest) error
	ResetPassword(ctx context.Context, r ResetPasswordRequest) error
}

type UserConfigServicer interface {
	Update(ctx context.Context, r *http.Request, req UpdateUserConfigRequest) error
	Get(ctx context.Context, r *http.Request) (domain.Profile, error)
}

type EmailServicer interface {
	Send(ctx context.Context, r *http.Request, req SendEmailRequest) error
}
