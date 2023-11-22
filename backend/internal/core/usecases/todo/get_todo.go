package todo

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	requests "github.com/aghex70/daps/internal/ports/requests/todo"
	"github.com/aghex70/daps/internal/ports/services/todo"
	"github.com/aghex70/daps/internal/ports/services/user"
	utils "github.com/aghex70/daps/utils/todo"

	"log"
)

type GetTodoUseCase struct {
	TodoService todo.Servicer
	UserService user.Servicer
	logger      *log.Logger
}

func (uc *GetTodoUseCase) Execute(ctx context.Context, r requests.GetTodoRequest, userID uint) (domain.Todo, error) {
	u, err := uc.UserService.Get(ctx, userID)
	if err != nil {
		return domain.Todo{}, err
	}

	if !u.Active {
		return domain.Todo{}, pkg.InactiveUserError
	}

	t, err := uc.TodoService.Get(ctx, r.TodoID)
	if err != nil {
		return domain.Todo{}, err
	}
	if owner := utils.IsTodoOwner(t.OwnerID, u.ID); !owner {
		return domain.Todo{}, pkg.UnauthorizedError
	}

	return t, nil
}

func NewGetTodoUseCase(s todo.Servicer, u user.Servicer, logger *log.Logger) *GetTodoUseCase {
	return &GetTodoUseCase{
		TodoService: s,
		UserService: u,
		logger:      logger,
	}
}
