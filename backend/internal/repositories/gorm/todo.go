package category

import (
	"context"
	"time"

	"github.com/aghex70/daps/internal/core/domain"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Name        string
	Description *string
	Completed   bool
	CompletedAt *time.Time
	Active      bool
	Priority    int
	CategoryID  uint
	Link        string
	Recurring   bool
	Recurrency  string
	StartedAt   *time.Time
	Suggestable bool
	SuggestedAt *time.Time
	UserID      uint
}

func (td Todo) ToDto() domain.Todo {
	return domain.Todo{
		ID:          td.ID,
		CreatedAt:   td.CreatedAt,
		Name:        td.Name,
		Description: td.Description,
		Completed:   td.Completed,
		CompletedAt: td.CompletedAt,
		Active:      td.Active,
		Priority:    domain.Priority(td.Priority),
		CategoryID:  td.CategoryID,
		Link:        td.Link,
		Recurring:   td.Recurring,
		Recurrency:  td.Recurrency,
		StartedAt:   td.StartedAt,
		Suggestable: td.Suggestable,
		SuggestedAt: td.SuggestedAt,
		UserID:      td.UserID,
	}
}

func (gr *GormRepository) GetTodo(ctx context.Context, id int) (domain.Todo, error) {
	return domain.Todo{}, nil
}

func (gr *GormRepository) GetTodos(ctx context.Context) ([]domain.Todo, error) {
	return []domain.Todo{}, nil
}

func (gr *GormRepository) CreateTodo(ctx context.Context, t domain.Todo) (domain.Todo, error) {
	return domain.Todo{}, nil
}

func (gr *GormRepository) UpdateTodo(ctx context.Context, t domain.Todo) error {
	return nil
}

func (gr *GormRepository) DeleteTodo(ctx context.Context, id int) error {
	return nil
}
