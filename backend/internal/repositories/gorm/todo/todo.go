package todo

import (
	"context"
	"database/sql"
	"github.com/aghex70/daps/internal/core/domain"
	"gorm.io/gorm"
	"log"
	"time"
)

type TodoGormRepository struct {
	*gorm.DB
	SqlDb  *sql.DB
	logger *log.Logger
}

type Todo struct {
	ID           int           `gorm:"primaryKey;column:id"`
	UserId       int           `gorm:"column:user_id"`
	Active       bool          `gorm:"column:active"`
	EndDate      *time.Time    `gorm:"column:end_date"`
	CategoryId   int           `gorm:"column:category_id"`
	Completed    bool          `gorm:"column:completed"`
	CreationDate time.Time     `gorm:"column:creation_date;autoCreateTime"`
	Description  string        `gorm:"column:description"`
	Duration     time.Duration `gorm:"column:duration"`
	Link         string        `gorm:"column:link"`
	Name         string        `gorm:"column:name"`
	//Prerequisite int           `gorm:"column:prerequisite_id"`
	Priority  int        `gorm:"column:priority"`
	StartDate *time.Time `gorm:"column:start_date"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (Todo) TableName() string {
	return "daps_todos"
}

func (gr *TodoGormRepository) Create(ctx context.Context, td domain.Todo) error {
	ntd := fromDto(td)
	result := gr.DB.Create(&ntd)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *TodoGormRepository) Delete(ctx context.Context, id uint) error {
	td := Todo{ID: int(id)}
	result := gr.DB.Delete(&td)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (gr *TodoGormRepository) List(ctx context.Context, userId uint) ([]domain.Todo, error) {
	var todos []Todo
	var todes []domain.Todo
	result := gr.DB.Find(&todos)
	if result.Error != nil {
		return []domain.Todo{}, result.Error
	}

	for _, t := range todos {
		todo := t.ToDto()
		todes = append(todes, todo)
	}
	return todes, nil
}
func (gr *TodoGormRepository) GetById(ctx context.Context, id uint, userId int) (domain.Todo, error) {
	var td Todo
	result := gr.DB.Where(&Todo{ID: int(id), UserId: userId}).First(&td)
	if result.Error != nil {
		return domain.Todo{}, result.Error
	}
	return td.ToDto(), nil
}
func (gr *TodoGormRepository) GetByNameAndLink(ctx context.Context, name string, link string, userId int) (domain.Todo, error) {
	var td Todo
	result := gr.DB.Where(&Todo{Name: name, Link: link, UserId: userId}).First(&td)
	if result.Error != nil {
		return domain.Todo{}, result.Error
	}

	return td.ToDto(), nil
}

func (gr *TodoGormRepository) Update(ctx context.Context, td domain.Todo) error {
	ntd := fromDto(td)
	result := gr.DB.Model(&ntd).Update("status", true)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *TodoGormRepository) Complete(ctx context.Context, id uint, userId int) error {
	result := gr.DB.Model(&Todo{ID: int(id), UserId: userId}).Update("completed", true).Update("end_date", time.Now())
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *TodoGormRepository) Enable(ctx context.Context, id uint, userId int) error {
	result := gr.DB.Model(&Todo{ID: int(id), UserId: userId}).Update("active", true).Update("start_date", time.Now())
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func NewTodoGormRepository(db *gorm.DB) (*TodoGormRepository, error) {
	return &TodoGormRepository{
		DB: db,
	}, nil
}

func (td Todo) ToDto() domain.Todo {
	return domain.Todo{
		Active:       td.Active,
		EndDate:      td.EndDate,
		Category:     td.CategoryId,
		Completed:    td.Completed,
		CreationDate: td.CreationDate,
		Description:  td.Description,
		Duration:     td.Duration,
		Link:         td.Link,
		Name:         td.Name,
		//Prerequisite: td.Prerequisite,
		//Priority:     td.Priority,
		StartDate: td.StartDate,
		User:      td.UserId,
	}
}

func fromDto(td domain.Todo) Todo {
	return Todo{
		Active:       td.Active,
		EndDate:      td.EndDate,
		CategoryId:   td.Category,
		Completed:    td.Completed,
		CreationDate: td.CreationDate,
		Description:  td.Description,
		Duration:     td.Duration,
		Link:         td.Link,
		Name:         td.Name,
		//Prerequisite: td.Prerequisite,
		//Priority:     td.Priority,
		StartDate: td.StartDate,
		UserId:    td.User,
	}
}
