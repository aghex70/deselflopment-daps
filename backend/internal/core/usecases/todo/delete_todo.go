package todo

import (
	"context"
	"github.com/aghex70/daps/internal/core/services/todo"
	"github.com/aghex70/daps/internal/pkg"
	//"github.com/aghex70/daps/server"
	utils "github.com/aghex70/daps/utils/todo"
	"log"
)

type DeleteTodoUseCase struct {
	TodoService todo.Service
	logger      *log.Logger
}

func (uc *DeleteTodoUseCase) Execute(ctx context.Context, id uint) error {
	//userId, _ := server.RetrieveJWTClaims(r, req)
	authorized := utils.HasWritePermissions(todo, todo.CategoryID)
	if !authorized {
		return pkg.UnauthorizedError
	}
	err := uc.TodoService.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func NewDeleteTodoUseCase(s todo.Service, logger *log.Logger) *DeleteTodoUseCase {
	return &DeleteTodoUseCase{
		TodoService: s,
		logger:      logger,
	}
}
