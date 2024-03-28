package todo

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	requests "github.com/aghex70/daps/internal/ports/requests/todo"
	"github.com/aghex70/daps/internal/ports/services/todo"
	"github.com/aghex70/daps/internal/ports/services/user"
	"log"
	"time"
)

type CreateTodoUseCase struct {
	TodoService todo.Servicer
	UserService user.Servicer
	logger      *log.Logger
}

func (uc *CreateTodoUseCase) Execute(ctx context.Context, r requests.CreateTodoRequest, userID uint) (domain.Todo, error) {
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
		Recurrency:  r.Recurrency,
		Priority:    domain.Priority(r.Priority),
		CategoryID:  r.CategoryID,
		OwnerID:     u.ID,
	}
	if *r.TargetDate != "" {
		var targetDate = setTargetDate(r.Recurrency, r.TargetDate)
		t.TargetDate = &targetDate
	} else {
		t.TargetDate = nil
	}

	t, err = uc.TodoService.Create(ctx, t)
	if err != nil {
		return domain.Todo{}, err
	}

	return t, nil
}

func setTargetDate(recurrency *int, targetDate *string) time.Time {
	// Define the layout of the date string
	layout := "2006-01-02"

	// Check if targetDate is nil
	if targetDate == nil {
		// Default to today's date
		currentDate := time.Now()

		// Add days to the current date
		newTargetDate := currentDate.AddDate(0, 0, *recurrency)

		// Set the time to 23:59:59
		newTargetDate = time.Date(newTargetDate.Year(), newTargetDate.Month(), newTargetDate.Day(), 23, 59, 59, 0, newTargetDate.Location())

		return newTargetDate
	}

	// If targetDate is already set, parse it into time.Time and set time to 23:59:59
	parsedTargetDate, _ := time.Parse(layout, *targetDate)
	parsedTargetDate = time.Date(parsedTargetDate.Year(), parsedTargetDate.Month(), parsedTargetDate.Day(), 23, 59, 59, 0, parsedTargetDate.Location())
	return parsedTargetDate
}

func NewCreateTodoUseCase(s todo.Servicer, u user.Servicer, logger *log.Logger) *CreateTodoUseCase {
	return &CreateTodoUseCase{
		TodoService: s,
		UserService: u,
		logger:      logger,
	}
}
