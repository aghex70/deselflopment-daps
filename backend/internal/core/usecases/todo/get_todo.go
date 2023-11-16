package todo

import (
	"context"
	"github.com/aghex70/daps/internal/core/services/todo"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	"log"
)

type GetTodoUseCase struct {
	TodoService todo.Service
	logger      *log.Logger
}

func (uc *GetTodoUseCase) Execute(ctx context.Context, id uint) (domain.Todo, error) {
	//userId, _ := server.RetrieveJWTClaims(r, req)
	authorized := utils.HasWritePermissions(todo, todo.CategoryID)
	if !authorized {
		return domain.Todo{}, pkg.UnauthorizedError
	}
	t, err := uc.TodoService.Get(ctx, id)
	if err != nil {
		return domain.Todo{}, err
	}
	return t, nil
}
