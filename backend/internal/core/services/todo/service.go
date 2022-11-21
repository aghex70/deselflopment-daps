package todo

import (
	"context"
	"errors"
	"fmt"
	"github.com/aghex70/daps/internal/core/domain"
	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/repositories/gorm/relationship"
	"github.com/aghex70/daps/internal/repositories/gorm/todo"
	"github.com/aghex70/daps/server"
	"log"
	"net/http"
)

type TodoService struct {
	logger                 *log.Logger
	todoRepository         *todo.TodoGormRepository
	relationshipRepository *relationship.RelationshipGormRepository
}

func (s TodoService) Create(ctx context.Context, r *http.Request, req ports.CreateTodoRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)
	err := s.CheckCategoryPermissions(ctx, int(userId), req.Category)
	if err != nil {
		return err
	}
	_, preexistent := s.CheckExistentTodo(ctx, req.Name, req.Category)
	if preexistent {
		return errors.New("already existent todo with that info")
	}

	ntd := domain.Todo{
		Category:    req.Category,
		Description: req.Description,
		Link:        req.Link,
		Name:        req.Name,
		Priority:    domain.Priority(req.Priority),
		Recurring:   req.Recurring,
	}

	err = s.todoRepository.Create(ctx, ntd)
	if err != nil {
		return err
	}

	return nil
}

func (s TodoService) Update(ctx context.Context, r *http.Request, req ports.UpdateTodoRequest) error {
	fmt.Println("12312312312313123")
	userId, _ := server.RetrieveJWTClaims(r, req)
	err := s.CheckCategoryPermissions(ctx, int(userId), req.Category)
	if err != nil {
		return err
	}

	fmt.Println("342345345345345")
	ntd := domain.Todo{
		ID:          int(req.TodoId),
		Category:    req.Category,
		Description: req.Description,
		Link:        req.Link,
		Name:        req.Name,
		Priority:    domain.Priority(req.Priority),
		Recurring:   req.Recurring,
	}

	err = s.todoRepository.Update(ctx, ntd)
	fmt.Println("6666666666666664564564546")
	if err != nil {
		return err
	}
	return nil
}

func (s TodoService) Complete(ctx context.Context, r *http.Request, req ports.CompleteTodoRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)
	err := s.CheckCategoryPermissions(ctx, int(userId), req.Category)
	if err != nil {
		return err
	}
	err = s.todoRepository.Complete(ctx, int(req.TodoId))
	if err != nil {
		return err
	}
	return nil
}

func (s TodoService) Activate(ctx context.Context, r *http.Request, req ports.ActivateTodoRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)
	err := s.CheckCategoryPermissions(ctx, int(userId), req.Category)
	if err != nil {
		return err
	}
	err = s.todoRepository.Activate(ctx, int(req.TodoId))
	if err != nil {
		return err
	}
	return nil
}

func (s TodoService) Start(ctx context.Context, r *http.Request, req ports.StartTodoRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)
	err := s.CheckCategoryPermissions(ctx, int(userId), req.Category)
	if err != nil {
		return err
	}
	err = s.todoRepository.Start(ctx, int(req.TodoId))
	if err != nil {
		return err
	}
	return nil
}

func (s TodoService) Get(ctx context.Context, r *http.Request, req ports.GetTodoRequest) (domain.TodoInfo, error) {
	userId, _ := server.RetrieveJWTClaims(r, req)
	err := s.CheckCategoryPermissions(ctx, int(userId), req.Category)
	if err != nil {
		return domain.TodoInfo{}, err
	}
	td, err := s.todoRepository.GetById(ctx, int(req.TodoId), int(userId))
	if err != nil {
		return domain.TodoInfo{}, err
	}
	return td, nil
}

func (s TodoService) List(ctx context.Context, r *http.Request, req ports.ListTodosRequest) ([]domain.Todo, error) {
	userId, _ := server.RetrieveJWTClaims(r, nil)
	err := s.CheckCategoryPermissions(ctx, int(userId), req.Category)
	fmt.Println("3333333333333333333333333333")
	if err != nil {
		fmt.Println("44444444444444444444444444")
		return []domain.Todo{}, err
	}
	todos, err := s.todoRepository.List(ctx, req.Category)
	fmt.Println("55555555555555555555555555")
	if err != nil {
		fmt.Println("6666666666666666666666666666666")
		return []domain.Todo{}, err
	}
	fmt.Println("7777777777777777777777777777777")
	return todos, nil
}

func (s TodoService) ListRecurring(ctx context.Context, r *http.Request) ([]domain.Todo, error) {
	userId, _ := server.RetrieveJWTClaims(r, nil)
	categoryIds, err := s.CheckCategoriesPermissions(ctx, int(userId))
	if err != nil {
		return []domain.Todo{}, err
	}
	todos, err := s.todoRepository.ListRecurring(ctx, categoryIds)
	if err != nil {
		return []domain.Todo{}, err
	}
	return todos, nil
}

func (s TodoService) ListCompleted(ctx context.Context, r *http.Request) ([]domain.Todo, error) {
	userId, _ := server.RetrieveJWTClaims(r, nil)
	categoryIds, err := s.CheckCategoriesPermissions(ctx, int(userId))
	if err != nil {
		return []domain.Todo{}, err
	}
	todos, err := s.todoRepository.ListCompleted(ctx, categoryIds)
	if err != nil {
		return []domain.Todo{}, err
	}
	return todos, nil
}

func (s TodoService) Delete(ctx context.Context, r *http.Request, req ports.DeleteTodoRequest) error {
	userId, _ := server.RetrieveJWTClaims(r, req)
	err := s.CheckCategoryPermissions(ctx, int(userId), req.Category)
	if err != nil {
		return err
	}
	err = s.todoRepository.Delete(ctx, int(req.TodoId))
	if err != nil {
		return err
	}
	return nil
}

func (s TodoService) Summary(ctx context.Context, r *http.Request) ([]domain.CategorySummary, error) {
	userId, _ := server.RetrieveJWTClaims(r, nil)
	summary, err := s.todoRepository.GetSummary(ctx, int(userId))
	if err != nil {
		return []domain.CategorySummary{}, err
	}
	return summary, nil
}

func NewtodoService(tr *todo.TodoGormRepository, rr *relationship.RelationshipGormRepository, logger *log.Logger) TodoService {
	return TodoService{
		logger:                 logger,
		todoRepository:         tr,
		relationshipRepository: rr,
	}
}
