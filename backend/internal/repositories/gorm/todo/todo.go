package todo

import (
	"context"
	"database/sql"
	"fmt"
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
	ID int `gorm:"primaryKey;column:id"`
	//UserId       int           `gorm:"column:user_id"`
	Active       bool          `gorm:"column:active"`
	EndDate      *time.Time    `gorm:"column:end_date"`
	Category     string        `gorm:"column:category"`
	Completed    bool          `gorm:"column:completed"`
	CreationDate time.Time     `gorm:"column:creation_date;autoCreateTime"`
	Description  string        `gorm:"column:description"`
	Duration     time.Duration `gorm:"column:duration"`
	Link         string        `gorm:"column:link"`
	Name         string        `gorm:"column:name"`
	//Prerequisite int           `gorm:"column:prerequisite_id"`
	Priority  int        `gorm:"column:priority"`
	StartDate *time.Time `gorm:"column:start_date"`
	//User         int           `gorm:"column:user_id"`
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
func (gr *TodoGormRepository) Get(ctx context.Context, userId uint) ([]domain.Todo, error) {
	var todos []Todo
	var todes []domain.Todo
	result := gr.DB.Find(&todos)
	fmt.Println("(gr *TodoGormRepository) Get\n todos")
	fmt.Printf("%v", todos)
	//result := gr.DB.Where(&Todo{UserId: int(userId)}).Find(&td)
	if result.Error != nil {
		return []domain.Todo{}, result.Error
	}

	for _, t := range todos {
		todo := t.ToDto()
		fmt.Printf("todo \n %+v", todo)
		todes = append(todes, todo)
	}
	fmt.Println("todes")
	fmt.Printf("+%v", todes)
	return todes, nil
}
func (gr *TodoGormRepository) GetById(ctx context.Context, id uint) (Todo, error) {
	var td Todo
	result := gr.DB.Where(&Todo{ID: int(id)}).First(&td)
	if result.Error != nil {
		return Todo{}, result.Error
	}
	return td, nil
}
func (gr *TodoGormRepository) GetByNameAndLink(ctx context.Context, name string, link string) (domain.Todo, error) {
	var td Todo
	result := gr.DB.Where(&Todo{Name: name, Link: link}).First(&td)
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

func (gr *TodoGormRepository) Complete(ctx context.Context, id uint) error {
	result := gr.DB.Model(&Todo{ID: int(id)}).Update("status", true)
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
		Active:  td.Active,
		EndDate: td.EndDate,
		//Category:     td.Category,
		Completed:    td.Completed,
		CreationDate: td.CreationDate,
		Description:  td.Description,
		Duration:     td.Duration,
		Link:         td.Link,
		Name:         td.Name,
		//Prerequisite: td.Prerequisite,
		//Priority:     td.Priority,
		StartDate: td.StartDate,
		//User:      td.User,
	}
}

func fromDto(td domain.Todo) Todo {
	return Todo{
		Active:  td.Active,
		EndDate: td.EndDate,
		//Category:     td.Category,
		Completed:    td.Completed,
		CreationDate: td.CreationDate,
		Description:  td.Description,
		Duration:     td.Duration,
		Link:         td.Link,
		Name:         td.Name,
		//Prerequisite: td.Prerequisite,
		//Priority:     td.Priority,
		StartDate: td.StartDate,
		//User:         td.User,
	}
}
