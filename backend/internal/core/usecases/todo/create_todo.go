package todo

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	requests "github.com/aghex70/daps/internal/ports/requests/todo"
	"github.com/aghex70/daps/internal/ports/services/todo"
	"github.com/aghex70/daps/internal/ports/services/user"
	"log"
)

type CreateTodoUseCase struct {
	TodoService todo.Servicer
	UserService user.Servicer
	logger      *log.Logger
}

func (uc *CreateTodoUseCase) Execute(ctx context.Context, userID uint, r requests.CreateTodoRequest) (domain.Todo, error) {
	u, err := uc.UserService.Get(ctx, userID)
	if err != nil {
		return domain.Todo{}, err
	}

	if !u.Active {
		return domain.Todo{}, pkg.InactiveUserError
	}

	t := domain.Todo{
		Name:        r.Name,
		Description: nil,
		Link:        nil,
		Recurring:   r.Recurring,
		Recurrency:  nil,
		Priority:    domain.Priority(r.Priority),
		CategoryID:  r.CategoryID,
		OwnerID:     u.ID,
	}
	t, err = uc.TodoService.Create(ctx, t)
	if err != nil {
		return domain.Todo{}, err
	}

	return t, nil
}

func NewCreateTodoUseCase(s todo.Servicer, u user.Servicer, logger *log.Logger) *CreateTodoUseCase {
	return &CreateTodoUseCase{
		TodoService: s,
		UserService: u,
		logger:      logger,
	}
}
