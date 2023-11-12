package todo

import (
	"context"
	domain2 "github.com/aghex70/daps/internal/ports/domain"
	"github.com/aghex70/daps/internal/ports/repositories/todo"
	"github.com/aghex70/daps/internal/ports/requests/todo"
	"log"
	"net/http"
)

type Service struct {
	logger     *log.Logger
	repository todo.Repository
}

func (s Service) Create(ctx context.Context, r *http.Request, req requests.CreateTodoRequest) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) Complete(ctx context.Context, r *http.Request, req requests.CompleteTodoRequest) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) Activate(ctx context.Context, r *http.Request, req requests.ActivateTodoRequest) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) Delete(ctx context.Context, r *http.Request, req requests.DeleteTodoRequest) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) Get(ctx context.Context, r *http.Request, req requests.GetTodoRequest) (interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) List(ctx context.Context, r *http.Request, req requests.ListTodosRequest) ([]domain2.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) ListCompleted(ctx context.Context, r *http.Request) ([]domain2.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) ListRecurring(ctx context.Context, r *http.Request) ([]domain2.Todo, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) ListSuggested(ctx context.Context, r *http.Request) ([]interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) Suggest(ctx context.Context, r *http.Request) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) Update(ctx context.Context, r *http.Request, req requests.UpdateTodoRequest) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) Start(ctx context.Context, r *http.Request, req requests.StartTodoRequest) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) Restart(ctx context.Context, r *http.Request, req requests.StartTodoRequest) error {
	//TODO implement me
	panic("implement me")
}

func (s Service) Summary(ctx context.Context, r *http.Request) ([]interface{}, error) {
	//TODO implement me
	panic("implement me")
}

func (s Service) Remind(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func NewTodoService(r todo.Repository, logger *log.Logger) Service {
	return Service{
		logger:     logger,
		repository: r,
	}
}
