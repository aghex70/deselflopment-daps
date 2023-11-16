package todo

import (
	"context"
	"github.com/aghex70/daps/internal/core/services/todo"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	//"github.com/aghex70/daps/server"
	"log"
)

type ListTodosUseCase struct {
	TodoService todo.Service
	logger      *log.Logger
}

func (uc *ListTodosUseCase) Execute(ctx context.Context, id uint) ([]domain.Todo, error) {
	////userId, _ := server.RetrieveJWTClaims(r, req)
	authorized := utils.HasWritePermissions(todo, todo.CategoryID)
	if !authorized {
		return []domain.Todo{}, pkg.UnauthorizedError
	}
	t, err := uc.TodoService.List(ctx, nil, nil)
	if err != nil {
		return []domain.Todo{}, err
	}
	return t, nil
}

func NewListTodosUseCase(s todo.Service, logger *log.Logger) *ListTodosUseCase {
	return &ListTodosUseCase{
		TodoService: s,
		logger:      logger,
	}
}
