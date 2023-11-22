package todo

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	"github.com/aghex70/daps/internal/ports/services/todo"
	"github.com/aghex70/daps/internal/ports/services/user"

	//"github.com/aghex70/daps/server"
	"log"
)

type ListTodosUseCase struct {
	TodoService todo.Servicer
	UserService user.Servicer
	logger      *log.Logger
}

func (uc *ListTodosUseCase) Execute(ctx context.Context, fields *map[string]interface{}, userID uint) ([]domain.Todo, error) {
	u, err := uc.UserService.Get(ctx, userID)
	if err != nil {
		return []domain.Todo{}, err
	}

	if !u.Active {
		return []domain.Todo{}, pkg.InactiveUserError
	}

	// Set the user ID into the fields map
	if fields == nil {
		fields = &map[string]interface{}{}
		(*fields)["owner_id"] = userID
	} else {
		(*fields)["owner_id"] = userID
	}
	todos, err := uc.TodoService.List(ctx, nil, fields)
	if err != nil {
		return []domain.Todo{}, err
	}
	return todos, nil
}

func NewListTodosUseCase(s todo.Servicer, u user.Servicer, logger *log.Logger) *ListTodosUseCase {
	return &ListTodosUseCase{
		TodoService: s,
		UserService: u,
		logger:      logger,
	}
}
