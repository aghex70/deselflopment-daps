package gorm

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
	"time"

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
	Link        *string
	Recurring   bool
	Recurrency  *string
	StartedAt   *time.Time
	Suggestable bool
	Suggested   bool
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
		//Priority:    Priority(t.Priority),
		CategoryID:  t.CategoryID,
		Link:        t.Link,
		Recurring:   t.Recurring,
		Recurrency:  t.Recurrency,
		StartedAt:   t.StartedAt,
		Suggestable: t.Suggestable,
		Suggested:   t.Suggested,
		SuggestedAt: t.SuggestedAt,
		UserID:      t.UserID,
	}
}

func (Todo) TableName() string {
	return "daps_todos"
}

type GormTodoRepository struct {
	*gorm.DB
}

func NewGormTodoRepository(db *gorm.DB) *GormTodoRepository {
	return &GormTodoRepository{DB: db}
}

func (gr *GormTodoRepository) Get(ctx context.Context, id uint) (domain.Todo, error) {
	return domain.Todo{}, nil
}

func (gr *GormTodoRepository) List(ctx context.Context, filters *map[string]interface{}) ([]domain.Todo, error) {
	return []domain.Todo{}, nil
}

func (gr *GormTodoRepository) Create(ctx context.Context, t domain.Todo) (domain.Todo, error) {
	return domain.Todo{}, nil
}

func (gr *GormTodoRepository) Update(ctx context.Context, t domain.Todo) error {
	return nil
}

func (gr *GormTodoRepository) Delete(ctx context.Context, id uint) error {
	return nil
}
