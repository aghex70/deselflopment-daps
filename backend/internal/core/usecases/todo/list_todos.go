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

func (uc *ListTodosUseCase) Execute(ctx context.Context, filters *map[string]interface{}, userID uint) ([]domain.Todo, error) {
	u, err := uc.UserService.Get(ctx, userID)
	if err != nil {
		return []domain.Todo{}, err
	}

	if !u.Active {
		return []domain.Todo{}, pkg.InactiveUserError
	}

	// Set the user ID into the fields map (retrieve only own todos)
	(*filters)["owner_id"] = userID

	// Temporary filter
	if _, ok := (*filters)["recurring"]; !ok {
		(*filters)["recurring"] = false
	}

	todos, err := uc.TodoService.List(ctx, filters)
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
