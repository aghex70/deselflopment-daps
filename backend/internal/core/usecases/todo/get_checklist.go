package todo

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	"github.com/aghex70/daps/internal/ports/services/todo"
	"github.com/aghex70/daps/internal/ports/services/user"
	"time"

	//"github.com/aghex70/daps/server"
	"log"
)

type GetChecklistUseCase struct {
	TodoService todo.Servicer
	UserService user.Servicer
	logger      *log.Logger
}

func (uc *GetChecklistUseCase) Execute(ctx context.Context, userID uint) ([]domain.Todo, error) {
	u, err := uc.UserService.Get(ctx, userID)
	if err != nil {
		return []domain.Todo{}, err
	}

	if !u.Active {
		return []domain.Todo{}, pkg.InactiveUserError
	}

	// Set the user ID into the fields map (retrieve only own todos)
	filters := &map[string]interface{}{}
	(*filters)["owner_id"] = userID

	// Set the target date to the end of the day
	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	(*filters)["target_date"] = today

	todos, err := uc.TodoService.List(ctx, filters)
	if err != nil {
		return []domain.Todo{}, err
	}
	return todos, nil
}

func NewGetChecklistUseCase(s todo.Servicer, u user.Servicer, logger *log.Logger) *GetChecklistUseCase {
	return &GetChecklistUseCase{
		TodoService: s,
		UserService: u,
		logger:      logger,
	}
}
