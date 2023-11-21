package gorm

import (
	"context"
	"fmt"
	"github.com/aghex70/daps/internal/ports/domain"
	"strings"
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
	OwnerID     uint
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
		OwnerID:     t.OwnerID,
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
		OwnerID:     t.OwnerID,
	}
}

func TodosFromDto(ts []domain.Todo) []Todo {
	var todos []Todo
	for _, t := range ts {
		todos = append(todos, TodoFromDto(t))
	}
	return todos
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

	query := gr.DB
	if filters != nil {
		// Convert map[string]interface{} to a slice of arguments
		var args []interface{}
		var conditions []string
		for key, value := range *filters {
			conditions = append(conditions, fmt.Sprintf("%s = ?", key))
			args = append(args, value)
		}
		query = query.Where(strings.Join(conditions, " AND "), args...)
	}

	result := query.Find(&ts)
	if result.Error != nil {
		return []domain.Todo{}, result.Error
	}

	for _, c := range ts {
		ts := c.ToDto()
		todos = append(todos, ts)
	}
	return todos, nil
}

func (gr *TodoRepository) Update(ctx context.Context, id uint, filters *map[string]interface{}) error {
	var t Todo
	t.ID = id
	result := gr.DB.Model(&t).Updates(*filters)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
