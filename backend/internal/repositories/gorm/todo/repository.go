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
	ID           int        `gorm:"primaryKey;column:id"`
	Active       bool       `gorm:"column:active"`
	EndDate      *time.Time `gorm:"column:end_date"`
	CategoryId   int        `gorm:"column:category_id"`
	CategoryName string     `json:"category_name"`
	Completed    bool       `gorm:"column:completed"`
	CreationDate time.Time  `gorm:"column:creation_date;autoCreateTime"`
	Description  string     `gorm:"column:description"`
	Link         string     `gorm:"column:link"`
	Name         string     `gorm:"column:name"`
	Priority     int        `gorm:"column:priority"`
	Recurring    bool       `gorm:"column:recurring"`
	StartDate    *time.Time `gorm:"column:start_date"`
}

type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (Todo) TableName() string {
	return "daps_todos"
}

func (gr *TodoGormRepository) Create(ctx context.Context, td domain.Todo) error {
	fmt.Printf("\n\ntd: %+v", td)
	ntd := fromDto(td)
	fmt.Printf("\n\nntd: %+v", ntd)
	result := gr.DB.Create(&ntd)
	fmt.Printf("\n\nnresult: %+v", result)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *TodoGormRepository) Update(ctx context.Context, td domain.Todo) error {
	ntd := fromDto(td)
	fmt.Printf("\n\nntd: %+v", ntd)
	result := gr.DB.Model(&ntd).Where(Todo{ID: ntd.ID}).Updates(map[string]interface{}{
		"category_id": ntd.CategoryId,
		"description": ntd.Description,
		"link":        ntd.Link,
		"name":        ntd.Name,
		"priority":    ntd.Priority,
		"recurring":   ntd.Recurring,
	})

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *TodoGormRepository) Complete(ctx context.Context, id int) error {
	result := gr.DB.Model(&Todo{ID: id}).Update("completed", true).Update("end_date", time.Now())

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *TodoGormRepository) Activate(ctx context.Context, id int) error {
	result := gr.DB.Model(&Todo{ID: id}).Update("completed", false).Update("end_date", nil)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *TodoGormRepository) Start(ctx context.Context, id int) error {
	result := gr.DB.Model(&Todo{ID: id}).Update("active", true).Update("start_date", time.Now())

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *TodoGormRepository) GetById(ctx context.Context, id int, userId int) (domain.Todo, error) {
	var td Todo

	query := fmt.Sprintf("SELECT daps_todos.id, daps_todos.category_id, daps_todos.end_date, daps_todos.creation_date, daps_todos.completed, daps_todos.description, daps_todos.link, daps_todos.name, daps_todos.priority, daps_todos.recurring, daps_todos.start_date, daps_categories.name as category_name FROM daps_todos JOIN daps_categories ON daps_todos.category_id = daps_categories.id WHERE daps_todos.id = %d", id)
	fmt.Println(query)
	result := gr.DB.Raw(query).Scan(&td)

	//result := gr.DB.Where(&Todo{ID: int(id)}).Select()First(&td)
	//result := gr.DB.Where(&Todo{ID: int(id)}).First(&td)
	if result.Error != nil {
		return domain.Todo{}, result.Error
	}
	return td.ToDto(), nil
}

func (gr *TodoGormRepository) List(ctx context.Context, categoryId int) ([]domain.Todo, error) {
	var todos []Todo
	var todes []domain.Todo
	fmt.Printf("asdasdasdsdadasd")
	//result := gr.DB.Where(&Todo{CategoryId: categoryId, Completed: false}).Find(&todos)
	result := gr.DB.Where("category_id = ? AND completed = ?", categoryId, false).Find(&todos)
	if result.Error != nil {
		return []domain.Todo{}, result.Error
	}

	for _, t := range todos {
		todo := t.ToDto()
		todes = append(todes, todo)
	}
	return todes, nil
}

func (gr *TodoGormRepository) ListRecurring(ctx context.Context, categoryIds []int) ([]domain.Todo, error) {
	var todos []Todo
	var todes []domain.Todo
	result := gr.DB.Where("recurring = ? AND category_id IN ?", true, categoryIds).Find(&todos)
	if result.Error != nil {
		return []domain.Todo{}, result.Error
	}

	for _, t := range todos {
		todo := t.ToDto()
		todes = append(todes, todo)
	}
	return todes, nil
}

func (gr *TodoGormRepository) ListCompleted(ctx context.Context, categoryIds []int) ([]domain.Todo, error) {
	var todos []Todo
	var todes []domain.Todo
	result := gr.DB.Where("completed = ? AND recurring = ? AND category_id IN ?", true, false, categoryIds).Find(&todos)
	if result.Error != nil {
		return []domain.Todo{}, result.Error
	}

	for _, t := range todos {
		todo := t.ToDto()
		todes = append(todes, todo)
	}
	return todes, nil
}

func (gr *TodoGormRepository) Delete(ctx context.Context, id int) error {
	td := Todo{ID: id}
	result := gr.DB.Delete(&td)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (gr *TodoGormRepository) GetByNameAndCategory(ctx context.Context, name string, categoryId int) (domain.Todo, error) {
	var td Todo
	result := gr.DB.Where(&Todo{Name: name, CategoryId: categoryId}).First(&td)
	if result.Error != nil {
		return domain.Todo{}, result.Error
	}

	return td.ToDto(), nil
}

func (gr *TodoGormRepository) GetSummary(ctx context.Context, userId int) ([]domain.CategorySummary, error) {
	var cs []domain.CategorySummary
	query := fmt.Sprintf("SELECT daps_category_users.category_id, daps_categories.id, daps_categories.name, daps_categories.owner_id, SUM(CASE WHEN daps_todos.completed = FALSE AND daps_todos.priority = 4 then 1 else 0 END) as highest_priority_tasks, SUM(CASE WHEN daps_todos.completed = FALSE then 1 else 0 END) as tasks FROM daps_category_users INNER JOIN daps_categories ON daps_categories.id = daps_category_users.category_id LEFT JOIN daps_todos ON daps_todos.category_id = daps_category_users.category_id WHERE user_id = %d GROUP BY category_id", userId)
	fmt.Println(query)
	result := gr.DB.Raw(query).Scan(&cs)
	fmt.Printf("\n\ncs: %+v", cs)
	if result.Error != nil {
		return cs, result.Error
	}
	return cs, nil
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
		CategoryName: td.CategoryName,
		Completed:    td.Completed,
		CreationDate: td.CreationDate,
		Description:  td.Description,
		ID:           td.ID,
		Link:         td.Link,
		Name:         td.Name,
		Priority:     domain.Priority(td.Priority),
		Recurring:    td.Recurring,
		StartDate:    td.StartDate,
	}
}

func fromDto(td domain.Todo) Todo {
	return Todo{
		Active:       td.Active,
		EndDate:      td.EndDate,
		CategoryId:   td.Category,
		CategoryName: td.CategoryName,
		Completed:    td.Completed,
		CreationDate: td.CreationDate,
		Description:  td.Description,
		ID:           td.ID,
		Link:         td.Link,
		Name:         td.Name,
		Priority:     int(td.Priority),
		Recurring:    td.Recurring,
		StartDate:    td.StartDate,
	}
}
