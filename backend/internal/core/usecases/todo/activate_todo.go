package todo

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	requests "github.com/aghex70/daps/internal/ports/requests/todo"
	"github.com/aghex70/daps/internal/ports/services/todo"
	"github.com/aghex70/daps/internal/ports/services/user"
	utils "github.com/aghex70/daps/utils/todo"

	"log"
)

type ActivateTodoUseCase struct {
	TodoService todo.Servicer
	UserService user.Servicer
	logger      *log.Logger
}

func (uc *ActivateTodoUseCase) Execute(ctx context.Context, r requests.GetTodoRequest, userID uint) error {
	u, err := uc.UserService.Get(ctx, userID)
	if err != nil {
		return err
	}

	if !u.Active {
		return pkg.InactiveUserError
	}

	t, err := uc.TodoService.Get(ctx, r.TodoID)
	if err != nil {
		return err
	}
	owner := utils.IsTodoOwner(t.OwnerID, userID)
	if !owner {
		return pkg.UnauthorizedError
	}

	if err = uc.TodoService.Activate(ctx, t.ID); err != nil {
		return err
	}
	return nil
}

func NewActivateTodoUseCase(s todo.Servicer, u user.Servicer, logger *log.Logger) *ActivateTodoUseCase {
	return &ActivateTodoUseCase{
		TodoService: s,
		UserService: u,
		logger:      logger,
	}
}
