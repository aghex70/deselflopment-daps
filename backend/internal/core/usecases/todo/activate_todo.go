package todo

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	"github.com/aghex70/daps/internal/ports/services/todo"
	utils "github.com/aghex70/daps/utils/todo"
	"log"
)

type ActivateTodoUseCase struct {
	TodoService todo.Servicer
	logger      *log.Logger
}

func (uc *ActivateTodoUseCase) Execute(ctx context.Context, todo domain.Todo) (domain.Todo, error) {
	authorized := utils.HasWritePermissions(todo, todo.CategoryID)
	if !authorized {
		return domain.Todo{}, pkg.UnauthorizedError
	}
	todo.Active = true
	t, err := uc.TodoService.Update(ctx, 0, todo)
	if err != nil {
		return domain.Todo{}, err
	}
	return t, nil
}

func NewActivateTodoUseCase(s todo.Servicer, logger *log.Logger) *ActivateTodoUseCase {
	return &ActivateTodoUseCase{
		TodoService: s,
		logger:      logger,
	}
}
