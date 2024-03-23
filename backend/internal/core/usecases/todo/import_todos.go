package todo

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/services/todo"
	"github.com/aghex70/daps/internal/ports/services/user"
	"log"
	"mime/multipart"
)

type ImportTodosUseCase struct {
	TodoService todo.Servicer
	UserService user.Servicer
	logger      *log.Logger
}

func (uc *ImportTodosUseCase) Execute(ctx context.Context, f multipart.File, userID uint) error {
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

func NewImportTodosUseCase(s todo.Servicer, u user.Servicer, logger *log.Logger) *ImportTodosUseCase {
	return &ImportTodosUseCase{
		TodoService: s,
		UserService: u,
		logger:      logger,
	}
}
