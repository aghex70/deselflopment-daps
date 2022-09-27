package todo

import (
	"context"
	"errors"
	"github.com/aghex70/daps/internal/core/domain"
	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/repositories/gorm/todo"
	"github.com/aghex70/daps/server"
	"log"
	"net/http"
)

type TodoService struct {
	logger         *log.Logger
	todoRepository *todo.TodoGormRepository
}

func (s TodoService) Create(ctx context.Context, r *http.Request, req ports.CreateTodoRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)
	t, preexistent := s.CheckExistentTodo(ctx, req.Name, req.Link, int(userId))
	if preexistent && t.Active {
		return errors.New("already existent and active todo")
	}

	ntd := domain.Todo{
		Category:    req.Category,
		User:        int(userId),
		Description: req.Description,
		Duration:    req.Duration,
		Link:        req.Link,
		Name:        req.Name,
		Priority:    domain.Priority(req.Priority),
	}

	if preexistent && !t.Active {
		err := s.todoRepository.Update(ctx, ntd)
		if err != nil {
			return err
		}
	} else {
		err := s.todoRepository.Create(ctx, ntd)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s TodoService) Complete(ctx context.Context, r *http.Request, req ports.CompleteTodoRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)
	err := s.todoRepository.Complete(ctx, int(req.TodoId), int(userId))
	if err != nil {
		return err
	}
	return nil
}

func (s TodoService) Delete(ctx context.Context, r *http.Request, req ports.DeleteTodoRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)
	err := s.todoRepository.Delete(ctx, int(req.TodoId), int(userId))
	if err != nil {
		return err
	}
	return nil
}

func (s TodoService) Get(ctx context.Context, r *http.Request, req ports.GetTodoRequest) (domain.Todo, error) {
	userId, _ := server.RetrieveJWTClaims(r, req)
	td, err := s.todoRepository.GetById(ctx, int(req.TodoId), int(userId))
	if err != nil {
		return domain.Todo{}, err
	}
	return td, nil
}

func (s TodoService) List(ctx context.Context, r *http.Request, sorting string, filters string) ([]domain.Todo, error) {
	userId, _ := server.RetrieveJWTClaims(r, nil)
	todos, err := s.todoRepository.List(ctx, int(userId), sorting, filters)
	if err != nil {
		return []domain.Todo{}, err
	}
	return todos, nil
}

func (s TodoService) Update(ctx context.Context, r *http.Request, req ports.UpdateTodoRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)

	ntd := domain.Todo{
		ID:          int(req.TodoId),
		Category:    req.Category,
		User:        int(userId),
		Description: req.Description,
		Duration:    req.Duration,
		Link:        req.Link,
		Name:        req.Name,
		Priority:    domain.Priority(req.Priority),
	}

	err := s.todoRepository.Update(ctx, ntd)
	if err != nil {
		return err
	}
	return nil
}

func (s TodoService) Start(ctx context.Context, r *http.Request, req ports.StartTodoRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)
	err := s.todoRepository.Start(ctx, int(req.TodoId), int(userId))
	if err != nil {
		return err
	}
	return nil
}

func NewtodoService(tr *todo.TodoGormRepository, logger *log.Logger) TodoService {
	return TodoService{
		logger:         logger,
		todoRepository: tr,
	}
}
