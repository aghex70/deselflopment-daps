package todo

import (
	"context"
	"errors"
	"fmt"
	"github.com/aghex70/daps/internal/core/domain"
	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/repositories/gorm/todo"
	"gorm.io/gorm"
	"log"
)

type TodoService struct {
	logger         *log.Logger
	todoRepository *todo.TodoGormRepository
}

func (s TodoService) Create(ctx context.Context, r ports.CreateTodoRequest) error {
	preexistent := s.CheckExistentTodo(ctx, r.Name, r.Link)
	if preexistent {
		return errors.New("already existent todo")
	}

	ntd := domain.Todo{
		//Category:     r.Category,
		Description: r.Description,
		Duration:    0,
		Link:        r.Link,
		Name:        r.Name,
		Priority:    0,
	}

	err := s.todoRepository.Create(ctx, ntd)
	if err != nil {
		return err
	}
	return nil
}
func (s TodoService) Complete(ctx context.Context, r ports.CompleteTodoRequest) error {
	gr := ports.GetTodoRequest{TodoId: r.TodoId}
	_, err := s.Get(ctx, gr)
	if err != nil {
		return err
	}
	err = s.todoRepository.Complete(ctx, uint(r.TodoId))
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
	panic("foo")
	//td, err := s.todoRepository.GetById(ctx, uint(r.TodoId), uint(r.TodoId))
	//if err != nil {
	//	return domain.Todo{}, err
	//}
	//return td, nil
}
func (s TodoService) List(ctx context.Context, r ports.ListTodosRequest) ([]domain.Todo, error) {
	todos, err := s.todoRepository.Get(ctx, uint(r.TodoId))
	if err != nil {
		return []domain.Todo{}, err
	}
	fmt.Println("(s TodoService) List(ctx")
	fmt.Printf("%v", todos)
	return todos, nil
	//return td, nil
}

func (s TodoService) CheckExistentTodo(ctx context.Context, name string, link string) bool {
	_, err := s.todoRepository.GetByNameAndLink(ctx, name, link)
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

func NewTodoService(tr *todo.TodoGormRepository, logger *log.Logger) TodoService {
	return TodoService{
		logger:         logger,
		todoRepository: tr,
	}
}
