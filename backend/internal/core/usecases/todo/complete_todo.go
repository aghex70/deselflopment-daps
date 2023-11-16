package todo

import (
	"context"
	"github.com/aghex70/daps/internal/core/services/todo"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	//"github.com/aghex70/daps/server"
	utils "github.com/aghex70/daps/utils/todo"
	"log"
)

type CompleteTodoUseCase struct {
	TodoService todo.Service
	logger      *log.Logger
}

func (uc *CompleteTodoUseCase) Execute(ctx context.Context, todo domain.Todo) (domain.Todo, error) {
	////userId, _ := server.RetrieveJWTClaims(r, req)
	authorized := utils.HasWritePermissions(todo, todo.CategoryID)
	if !authorized {
		return domain.Todo{}, pkg.UnauthorizedError
	}
	t, err := uc.TodoService.Update(ctx, 0, todo)
	if err != nil {
		return domain.Todo{}, err
	}
	return todo, nil
}

func NewCompleteTodoUseCase(s todo.Service, logger *log.Logger) *CompleteTodoUseCase {
	return &CompleteTodoUseCase{
		TodoService: s,
		logger:      logger,
	}
}
