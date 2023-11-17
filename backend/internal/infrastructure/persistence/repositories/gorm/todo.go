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
		Priority:    domain.Priority(t.Priority),
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

func TodoFromDto(t domain.Todo) Todo {
	return Todo{
		Name:        t.Name,
		Description: t.Description,
		Completed:   t.Completed,
		CompletedAt: t.CompletedAt,
		Active:      t.Active,
		Priority:    int(t.Priority),
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

type TodoRepository struct {
	*gorm.DB
}

func NewGormTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{DB: db}
}

func (gr *TodoRepository) Create(ctx context.Context, t domain.Todo) (domain.Todo, error) {
	td := TodoFromDto(t)
	result := gr.DB.Create(&td)
	if result.Error != nil {
		return domain.Todo{}, result.Error
	}
	return td.ToDto(), nil
}

func (gr *TodoRepository) Get(ctx context.Context, id uint) (domain.Todo, error) {
	var t Todo
	result := gr.DB.First(&t, id)
	if result.Error != nil {
		return domain.Todo{}, result.Error
	}
	return t.ToDto(), nil
}

func (gr *TodoRepository) Delete(ctx context.Context, id uint) error {
	result := gr.DB.Delete(&Todo{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *TodoRepository) List(ctx context.Context, ids *[]uint, filters *map[string]interface{}) ([]domain.Todo, error) {
	var ts []Todo
	var todos []domain.Todo

	if filters != nil {
		result := gr.DB.Where(filters).Find(&ts)
		if result.Error != nil {
			return []domain.Todo{}, result.Error
		}
	} else {
		result := gr.DB.Find(&ts)
		if result.Error != nil {
			return []domain.Todo{}, result.Error
		}
	}

	for _, c := range ts {
		ts := c.ToDto()
		todos = append(todos, ts)
	}
	return todos, nil
}

func (gr *TodoRepository) Update(ctx context.Context, id uint, t domain.Todo) (domain.Todo, error) {
	result := gr.DB.Model(&Todo{}).Where("id = ?", id).Updates(Todo{
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
	})
	if result.Error != nil {
		return domain.Todo{}, result.Error
	}
	return t, nil
}
