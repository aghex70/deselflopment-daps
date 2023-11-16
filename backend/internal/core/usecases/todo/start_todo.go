package todo

import (
	"context"
	"github.com/aghex70/daps/internal/core/services/todo"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	utils "github.com/aghex70/daps/utils/todo"
	"log"
)

type RestartTodoUseCase struct {
	TodoService todo.Service
	logger      *log.Logger
}

func (uc *RestartTodoUseCase) Execute(ctx context.Context, todo domain.Todo) (domain.Todo, error) {
	//userId, _ := server.RetrieveJWTClaims(r, req)
	authorized := utils.HasWritePermissions(todo, todo.CategoryID)
	if !authorized {
		return domain.Todo{}, pkg.UnauthorizedError
	}
	t, err := uc.TodoService.Update(ctx, 0, todo)
	if err != nil {
		return domain.Todo{}, err
	}
	return t, nil
}

func NewRestartTodoUseCase(s todo.Service, logger *log.Logger) *RestartTodoUseCase {
	return &RestartTodoUseCase{
		TodoService: s,
		logger:      logger,
	}
}
