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

func (t Todo) ToDto() domain.Todo {
	return domain.Todo{
		ID:          t.ID,
		CreatedAt:   t.CreatedAt,
		Name:        t.Name,
		Description: t.Description,
		Completed:   t.Completed,
		CompletedAt: t.CompletedAt,
		Active:      t.Active,
		Priority:    domain.Priority(t.Priority),
		CategoryID:  t.CategoryID,
		Link:        t.Link,
		Recurring:   t.Recurring,
		Recurrency:  t.Recurrency,
		StartedAt:   t.StartedAt,
		Suggestable: t.Suggestable,
		SuggestedAt: t.SuggestedAt,
		UserID:      t.UserID,
	}
}

func (Todo) TableName() string {
	return "daps_todos"
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
