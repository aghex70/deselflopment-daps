package todo

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/aghex70/daps/internal/core/domain"
	"github.com/aghex70/daps/internal/errors"
	"github.com/aghex70/daps/pkg"
	"gorm.io/gorm"
)

type GormRepository struct {
	*gorm.DB
	SqlDb *sql.DB
}

type Todo struct {
	Id             int        `gorm:"primaryKey;column:id"`
	Active         bool       `gorm:"column:active"`
	EndDate        *time.Time `gorm:"column:end_date"`
	CategoryId     int        `gorm:"column:category_id"`
	Completed      bool       `gorm:"column:completed"`
	CreationDate   time.Time  `gorm:"column:creation_date;autoCreateTime"`
	Description    string     `gorm:"column:description"`
	Link           string     `gorm:"column:link"`
	Name           string     `gorm:"column:name"`
	Priority       int        `gorm:"column:priority"`
	Recurring      bool       `gorm:"column:recurring"`
	Recurrency     string     `gorm:"column:recurrency"`
	StartDate      *time.Time `gorm:"column:start_date"`
	Suggestable    bool       `gorm:"suggestable"`
	SuggestionDate *time.Time `gorm:"column:suggestion_date"`
}

type TodoInfo struct {
	KategoryName string `json:"category_name"`
	Todo
}

type Tabler interface {
	TableName() string
}

func (Todo) TableName() string {
	return "daps_todos"
}

func (gr *GormRepository) Create(ctx context.Context, td domain.Todo) error {
	ntd := fromDto(td)
	result := gr.DB.Create(&ntd)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *GormRepository) Update(ctx context.Context, td domain.Todo) error {
	ntd := fromDto(td)
	result := gr.DB.Model(&ntd).Where(Todo{Id: ntd.Id}).Updates(map[string]interface{}{
		"category_id": ntd.CategoryId,
		"description": ntd.Description,
		"link":        ntd.Link,
		"name":        ntd.Name,
		"priority":    ntd.Priority,
		"recurring":   ntd.Recurring,
		"recurrency":  ntd.Recurrency,
		"suggestable": ntd.Suggestable,
	})

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *GormRepository) Complete(ctx context.Context, id int) error {
	result := gr.DB.Model(&Todo{Id: id}).Update("completed", true).Update("end_date", time.Now())

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *GormRepository) Activate(ctx context.Context, id int) error {
	result := gr.DB.Model(&Todo{Id: id}).Update("completed", false).Update("end_date", nil)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *GormRepository) Start(ctx context.Context, id int) error {
	result := gr.DB.Model(&Todo{Id: id}).Update("active", true).Update("start_date", time.Now())

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *GormRepository) Restart(ctx context.Context, id int) error {
	result := gr.DB.Model(&Todo{Id: id}).Update("active", false).Update("start_date", nil)

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *GormRepository) GetById(ctx context.Context, id int) (domain.TodoInfo, error) {
	var ti TodoInfo
	query := fmt.Sprintf("SELECT daps_todos.id, daps_todos.category_id, daps_todos.end_date, daps_todos.creation_date, daps_todos.completed, daps_todos.description, daps_todos.link, daps_todos.name, daps_todos.priority, daps_todos.recurring, daps_todos.start_date, daps_todos.recurrency, daps_todos.suggestable, daps_categories.name as category_name FROM daps_todos JOIN daps_categories ON daps_todos.category_id = daps_categories.id WHERE daps_todos.id = %d", id)
	result := gr.DB.Raw(query).Scan(&ti)

	if result.Error != nil {
		return domain.TodoInfo{}, result.Error
	}
	return ti.ToDto(), nil
}

func (gr *GormRepository) List(ctx context.Context, categoryId int) ([]domain.Todo, error) {
	var todos []Todo
	var todes []domain.Todo
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

func (gr *GormRepository) ListRecurring(ctx context.Context, categoryIds []int) ([]domain.Todo, error) {
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

func (gr *GormRepository) ListCompleted(ctx context.Context, categoryIds []int) ([]domain.Todo, error) {
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

func (gr *GormRepository) ListSuggested(ctx context.Context, userId int) ([]domain.TodoInfo, error) {
	var tis []TodoInfo
	var todosInfo []domain.TodoInfo
	query := fmt.Sprintf("SELECT daps_todos.name, daps_todos.id, daps_todos.category_id, daps_todos.active, daps_todos.suggestion_date, daps_categories.name as category_name FROM daps_todos JOIN daps_categories ON daps_todos.category_id = daps_categories.id WHERE daps_todos.suggested = true AND daps_todos.completed = false AND daps_categories.owner_id = %d AND daps_todos.suggestable = true ORDER BY RAND() LIMIT 8", userId)
	result := gr.DB.Raw(query).Scan(&tis)

	if result.Error != nil {
		return []domain.TodoInfo{}, result.Error
	}

	for _, ti := range tis {
		todoInfo := ti.ToDto()
		todosInfo = append(todosInfo, todoInfo)
	}

	return todosInfo, nil
}

func (gr *GormRepository) Suggest(ctx context.Context, userId int) error {
	var suggestedTodosNumber int

	// Retrieve current number of suggested todos
	query := fmt.Sprintf("SELECT COUNT(*) FROM daps_todos JOIN daps_categories ON daps_todos.category_id = daps_categories.id WHERE daps_todos.suggested = true AND daps_todos.completed = false AND daps_categories.owner_id = %d", userId)
	_ = gr.DB.Raw(query).Scan(&suggestedTodosNumber)

	if suggestedTodosNumber >= pkg.MaximumConcurrentSuggestions {
		return nil
	}

	newSuggestedTodosNumber := pkg.MaximumConcurrentSuggestions - suggestedTodosNumber

	// Retrieve todos that are going to be suggested
	var ids []int
	query = fmt.Sprintf("SELECT daps_todos.id FROM daps_todos JOIN daps_categories ON daps_todos.category_id = daps_categories.id WHERE daps_todos.recurring = false AND daps_todos.suggested = false AND daps_todos.completed = false AND daps_todos.active = false AND daps_categories.owner_id = %d AND daps_todos.suggestable = true ORDER BY RAND() LIMIT %d", userId, newSuggestedTodosNumber)
	_ = gr.DB.Raw(query).Scan(&ids)

	var idList string
	for i, id := range ids {
		idList += strconv.Itoa(id)
		if i < len(ids)-1 {
			idList += ","
		}
	}

	// Set them to suggested
	query2 := fmt.Sprintf("UPDATE daps_todos SET suggested = true, suggestion_date = NOW() WHERE id IN (%s)", idList)
	result := gr.DB.Exec(query2)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *GormRepository) Delete(ctx context.Context, id int) error {
	td := Todo{Id: id}
	result := gr.DB.Delete(&td)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func (gr *GormRepository) GetByNameAndCategory(ctx context.Context, name string, categoryId int) (domain.Todo, error) {
	var td Todo
	result := gr.DB.Where(&Todo{Name: name, CategoryId: categoryId}).First(&td)
	if result.Error != nil {
		return domain.Todo{}, result.Error
	}

	return td.ToDto(), nil
}

func (gr *GormRepository) GetSummary(ctx context.Context, userId int) ([]domain.CategorySummary, error) {
	var cs []domain.CategorySummary
	query := fmt.Sprintf("SELECT daps_category_users.category_id, daps_categories.id, daps_categories.name, daps_categories.owner_id, daps_categories.shared, SUM(CASE WHEN daps_todos.completed = FALSE AND daps_todos.priority = 5 then 1 else 0 END) as highest_priority_tasks, SUM(CASE WHEN daps_todos.completed = FALSE then 1 else 0 END) as tasks FROM daps_category_users INNER JOIN daps_categories ON daps_categories.id = daps_category_users.category_id LEFT JOIN daps_todos ON daps_todos.category_id = daps_category_users.category_id WHERE user_id = %d GROUP BY category_id", userId)
	result := gr.DB.Raw(query).Scan(&cs)
	if result.Error != nil {
		return cs, result.Error
	}
	return cs, nil
}

func (gr *GormRepository) GetRemindSummary(ctx context.Context, userId int) ([]domain.RemindSummary, error) {
	var rs []domain.RemindSummary

	// Check if reminder has already been sent
	var e domain.Email
	subject := fmt.Sprintf("📣 DAPS - Tareas pendientes (%s) 📣", time.Now().Format("02/01/2006"))
	query := fmt.Sprintf("SELECT * FROM daps_emails WHERE subject IN ('%s') AND sent = true AND user_id = %d", subject, userId)
	result := gr.DB.Raw(query).Scan(&e)

	if result.Error != nil {
		return rs, result.Error
	}

	if result.RowsAffected != 0 {
		return rs, errors.ReminderAlreadySent
	}

	// Retrieve todos to be reminded
	// - Todos that are not completed
	// - Belong to a category that is notifiable
	// - Belong to a category that belongs to the user
	query = fmt.Sprintf("SELECT t.name as todo_name, c.name as category_name, t.priority as todo_priority, t.description as todo_description, t.link as todo_link FROM daps_users u JOIN daps_user_configs uc ON uc.user_id = u.id AND uc.auto_remind = true JOIN daps_categories c ON c.owner_id = u.id JOIN (SELECT id, name, priority, description, completed, link, category_id, ROW_NUMBER() OVER (PARTITION BY category_id ORDER BY priority DESC, RAND()) as rn FROM daps_todos) t ON t.category_id = c.id AND t.rn <= 3 WHERE c.owner_id = %d AND t.completed = false and c.notifiable = true ORDER BY u.id, c.id, t.id", userId)
	result = gr.DB.Raw(query).Scan(&rs)
	if result.Error != nil {
		return rs, result.Error
	}

	if result.RowsAffected == 0 {
		return rs, gorm.ErrRecordNotFound
	}
	return rs, nil
}

func NewTodoGormRepository(db *gorm.DB) (*GormRepository, error) {
	return &GormRepository{
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
		Id:           td.Id,
		Link:         td.Link,
		Name:         td.Name,
		Priority:     domain.Priority(td.Priority),
		Recurring:    td.Recurring,
		Recurrency:   td.Recurrency,
		Suggestable:  td.Suggestable,
		StartDate:    td.StartDate,
	}
}

func fromDto(td domain.Todo) Todo {
	return Todo{
		Active:         td.Active,
		EndDate:        td.EndDate,
		CategoryId:     td.Category,
		Completed:      td.Completed,
		CreationDate:   td.CreationDate,
		Description:    td.Description,
		Id:             td.Id,
		Link:           td.Link,
		Name:           td.Name,
		Priority:       int(td.Priority),
		Recurring:      td.Recurring,
		Recurrency:     td.Recurrency,
		StartDate:      td.StartDate,
		Suggestable:    td.Suggestable,
		SuggestionDate: td.SuggestionDate,
	}
}

func (ti TodoInfo) ToDto() domain.TodoInfo {
	return domain.TodoInfo{
		Todo: domain.Todo{
			Active:         ti.Active,
			EndDate:        ti.EndDate,
			Category:       ti.CategoryId,
			KategoryName:   ti.KategoryName,
			Completed:      ti.Completed,
			CreationDate:   ti.CreationDate,
			Description:    ti.Description,
			Id:             ti.Id,
			Link:           ti.Link,
			Name:           ti.Name,
			Priority:       domain.Priority(ti.Priority),
			Recurring:      ti.Recurring,
			Recurrency:     ti.Recurrency,
			StartDate:      ti.StartDate,
			Suggestable:    ti.Suggestable,
			SuggestionDate: ti.SuggestionDate,
		},
		CategoryInfo: domain.CategoryInfo{KategoryName: ti.KategoryName},
	}
}
