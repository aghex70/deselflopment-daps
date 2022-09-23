package todo

import (
	"context"
	"errors"
	"github.com/aghex70/daps/internal/core/domain"
	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/repositories/gorm/todo"
	"github.com/aghex70/daps/server"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type TodoService struct {
	logger         *log.Logger
	todoRepository *todo.TodoGormRepository
}

func (s TodoService) Create(ctx context.Context, r *http.Request, req ports.CreateTodoRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)
	preexistent := s.CheckExistentTodo(ctx, req.Name, req.Link, int(userId))
	if preexistent {
		return errors.New("already existent todo")
	}

	ntd := domain.Todo{
		Category:    7,
		User:        int(userId),
		Description: req.Description,
		Duration:    0,
		Link:        req.Link,
		Name:        req.Name,
		Priority:    0,
	}

	err := s.todoRepository.Create(ctx, ntd)
	if err != nil {
		return err
	}
	return nil
}
func (s TodoService) Complete(ctx context.Context, r *http.Request, req ports.CompleteTodoRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)
	err := s.todoRepository.Complete(ctx, uint(req.TodoId), int(userId))
	if err != nil {
		return err
	}
	return nil
}

func (s TodoService) Delete(ctx context.Context, r ports.DeleteTodoRequest) error {
	err := s.todoRepository.Delete(ctx, uint(r.TodoId))
	if err != nil {
		return err
	}
	return nil
}
func (s TodoService) Get(ctx context.Context, r ports.GetTodoRequest) (domain.Todo, error) {
	td, err := s.todoRepository.GetById(ctx, uint(r.TodoId), int(r.TodoId))
	if err != nil {
		return domain.Todo{}, err
	}
	return td, nil
}
func (s TodoService) List(ctx context.Context, r ports.ListTodosRequest) ([]domain.Todo, error) {
	todos, err := s.todoRepository.List(ctx, uint(r.TodoId))
	if err != nil {
		return []domain.Todo{}, err
	}
	return todos, nil
}

func (s TodoService) CheckExistentTodo(ctx context.Context, name string, link string, userId int) bool {
	_, err := s.todoRepository.GetByNameAndLink(ctx, name, link, userId)
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

func NewTodoService(tr *todo.TodoGormRepository, logger *log.Logger) TodoService {
	return TodoService{
		logger:         logger,
		todoRepository: tr,
	}
}
