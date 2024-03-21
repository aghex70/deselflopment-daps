package todo

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	requests "github.com/aghex70/daps/internal/ports/requests/todo"
	"github.com/aghex70/daps/internal/ports/services/todo"
	"github.com/aghex70/daps/internal/ports/services/user"
	common "github.com/aghex70/daps/utils"
	utils "github.com/aghex70/daps/utils/todo"

	"log"
)

type UpdateTodoUseCase struct {
	TodoService todo.Servicer
	UserService user.Servicer
	logger      *log.Logger
}

func (uc *UpdateTodoUseCase) Execute(ctx context.Context, r requests.UpdateTodoRequest, userID uint) error {
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

	fields := common.StructToMap(r, "todo_id")
	var targetDate = setTargetDate(r.Recurrency, r.TargetDate)
	fields["target_date"] = &targetDate
	if err = uc.TodoService.Update(ctx, t.ID, &fields); err != nil {
		return err
	}
	return nil
}

func NewUpdateTodoUseCase(s todo.Servicer, u user.Servicer, logger *log.Logger) *UpdateTodoUseCase {
	return &UpdateTodoUseCase{
		TodoService: s,
		UserService: u,
		logger:      logger,
	}
}
