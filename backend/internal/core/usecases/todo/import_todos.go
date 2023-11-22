package todo

import (
	"context"
	"github.com/aghex70/daps/internal/core/services/todo"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/services/user"
	"log"
	"mime/multipart"
)

type ImportTodosUseCase struct {
	TodoService todo.Service
	UserService user.Servicer
	logger      *log.Logger
}

func (uc *ImportTodosUseCase) Execute(ctx context.Context, userID uint, f multipart.File) error {
	u, err := uc.UserService.Get(ctx, userID)
	if err != nil {
		return err
	}

	if !u.Active {
		return pkg.InactiveUserError
	}

	if !u.Admin {
		return pkg.UnauthorizedError
	}

	if err = uc.TodoService.Import(ctx, f); err != nil {
		return err
	}
	return nil
}

func NewImportTodosUseCase(s todo.Service, u user.Servicer, logger *log.Logger) *ImportTodosUseCase {
	return &ImportTodosUseCase{
		TodoService: s,
		UserService: u,
		logger:      logger,
	}
}