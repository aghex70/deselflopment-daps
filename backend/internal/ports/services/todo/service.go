package todo

import (
	"context"
	domain2 "github.com/aghex70/daps/internal/ports/domain"
	"github.com/aghex70/daps/internal/ports/requests/todo"
	"net/http"
)

type Servicer interface {
	Create(ctx context.Context, r *http.Request, req requests.CreateTodoRequest) error
	Complete(ctx context.Context, r *http.Request, req requests.CompleteTodoRequest) error
	Activate(ctx context.Context, r *http.Request, req requests.ActivateTodoRequest) error
	Delete(ctx context.Context, r *http.Request, req requests.DeleteTodoRequest) error
	//Get(ctx context.Context, r *http.Request, req requests.GetTodoRequest) (domain2.TodoInfo, error)
	List(ctx context.Context, r *http.Request, req requests.ListTodosRequest) ([]domain2.Todo, error)
	ListCompleted(ctx context.Context, r *http.Request) ([]domain2.Todo, error)
	ListRecurring(ctx context.Context, r *http.Request) ([]domain2.Todo, error)
	//ListSuggested(ctx context.Context, r *http.Request) ([]domain2.TodoInfo, error)
	Suggest(ctx context.Context, r *http.Request) error
	Update(ctx context.Context, r *http.Request, req requests.UpdateTodoRequest) error
	Start(ctx context.Context, r *http.Request, req requests.StartTodoRequest) error
	Restart(ctx context.Context, r *http.Request, req requests.StartTodoRequest) error
	//Summary(ctx context.Context, r *http.Request) ([]domain2.CategorySummary, error)
	Remind(ctx context.Context) error
}
