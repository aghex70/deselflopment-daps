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
	Priority       int        `gorm:"column:priority"`
	Recurring      bool       `gorm:"column:recurring"`
	StartDate      *time.Time `gorm:"column:start_date"`
	Suggested      bool       `gorm:"column:completed"`
	SuggestionDate *time.Time `gorm:"column:start_date"`
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

func (gr *TodoGormRepository) Delete(ctx context.Context, id int, userId int) error {
	td := Todo{ID: id, UserId: userId}
	result := gr.DB.Delete(&td)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (gr *TodoGormRepository) List(ctx context.Context, userId int, sorting string, filters string) ([]domain.Todo, error) {
	var todos []Todo
	var todes []domain.Todo
	//fields := "id = " + strconv.Itoa(userId)
	result := gr.DB.Where(&Todo{UserId: userId}).Order(sorting).Find(&todos)
	if result.Error != nil {
		return []domain.Todo{}, result.Error
	}

	for _, t := range todos {
		todo := t.ToDto()
		todes = append(todes, todo)
	}
	return todes, nil
}
func (gr *TodoGormRepository) GetById(ctx context.Context, id int, userId int) (domain.Todo, error) {
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
	result := gr.DB.Model(&ntd).Where(Todo{ID: ntd.ID, Active: false}).Updates(map[string]interface{}{
		"end_date":      nil,
		"category_id":   ntd.CategoryId,
		"completed":     false,
		"creation_date": time.Now(),
		"description":   ntd.Description,
		"duration":      ntd.Duration,
		"link":          ntd.Link,
		"name":          ntd.Name,
		"priority":      ntd.Priority,
	})

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *TodoGormRepository) Complete(ctx context.Context, id int, userId int) error {
	result := gr.DB.Model(&Todo{ID: int(id), UserId: userId}).Update("completed", true).Update("end_date", time.Now())
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *TodoGormRepository) Start(ctx context.Context, id int, userId int) error {
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
		ID:           td.ID,
		Link:         td.Link,
		Name:         td.Name,
		//Prerequisite: td.Prerequisite,
		//Priority:     td.Priority,
		Recurring: td.Recurring,
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
		ID:           td.ID,
		Link:         td.Link,
		Name:         td.Name,
		//Prerequisite: td.Prerequisite,
		//Priority:     td.Priority,
		Recurring: td.Recurring,
		StartDate: td.StartDate,
		UserId:    td.User,
	}
}
