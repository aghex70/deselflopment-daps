package todo

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	"github.com/aghex70/daps/internal/ports/services/todo"
	utils "github.com/aghex70/daps/utils/todo"
	"log"
)

type CreateTodoUseCase struct {
	TodoService todo.Servicer
	logger      *log.Logger
}

func (uc *CreateTodoUseCase) Execute(ctx context.Context, todo domain.Todo) (domain.Todo, error) {
	authorized := utils.HasWritePermissions(todo, todo.CategoryID)
	if !authorized {
		return domain.Todo{}, pkg.UnauthorizedError
	}
	t, err := uc.TodoService.Create(ctx, todo)
	if err != nil {
		return domain.Todo{}, err
	}
	return t, nil
}

func NewCreateTodoUseCase(s todo.Servicer, logger *log.Logger) *CreateTodoUseCase {
	return &CreateTodoUseCase{
		TodoService: s,
		logger:      logger,
	}
}
