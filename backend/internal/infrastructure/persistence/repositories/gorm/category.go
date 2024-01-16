package gorm

import (
	"context"
	"fmt"
	"github.com/aghex70/daps/internal/ports/domain"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Category struct {
	gorm.Model
	Name        string
	Description *string
	OwnerID     uint
	Users       []User `gorm:"many2many:daps_category_users;save_association:true"`
	Todos       []Todo
	Shared      bool
	Notifiable  bool
	Custom      bool
}

func (c Category) ToDto() domain.Category {
	var createdAt time.Time
	if !c.CreatedAt.IsZero() {
		createdAt = c.CreatedAt
	}

	var todos []domain.Todo
	if c.Todos != nil {
		for _, todo := range c.Todos {
			todos = append(todos, todo.ToDto())
		}
	}

	var users []domain.User
	if c.Users != nil {
		for _, user := range c.Users {
			users = append(users, user.ToDto())
		}
	}

	return domain.Category{
		ID:          c.ID,
		CreatedAt:   createdAt,
		Name:        c.Name,
		Description: c.Description,
		OwnerID:     c.OwnerID,
		Users:       users,
		Todos:       todos,
		Shared:      c.Shared,
		Notifiable:  c.Notifiable,
		Custom:      c.Custom,
	}
}

func CategoryFromDto(c domain.Category) Category {
	var users []User
	if c.Users != nil {
		for _, userDTO := range c.Users {
			user := UserFromDto(userDTO)
			users = append(users, user)
		}
	}

	var todos []Todo
	if c.Todos != nil {
		for _, todoDTO := range c.Todos {
			todo := TodoFromDto(todoDTO)
			todos = append(todos, todo)
		}
	}

	return Category{
		Name:        c.Name,
		Description: c.Description,
		OwnerID:     c.OwnerID,
		Users:       users,
		Todos:       todos,
		Shared:      c.Shared,
		Notifiable:  c.Notifiable,
		Custom:      c.Custom,
	}
}

func (Category) TableName() string {
	return "daps_categories"
}

func (gr *CategoryRepository) Create(ctx context.Context, c domain.Category) (domain.Category, error) {
	nc := CategoryFromDto(c)
	if result := gr.DB.Create(&nc); result.Error != nil {
		return domain.Category{}, result.Error
	}

	// Hack to get around the fact that GORM doesn't support many-to-many relationships
	if nc.Users == nil {
		if err := gr.DB.Association("Users").Append(nc.Users); err != nil {
			return domain.Category{}, err
		}
	}
	return nc.ToDto(), nil
}

func (gr *CategoryRepository) Get(ctx context.Context, id uint) (domain.Category, error) {
	var c Category
	if result := gr.DB.First(&c, id); result.Error != nil {
		return domain.Category{}, result.Error
	}

	// Retrieve users associated with the category if they exist
	if c.Users == nil {
		if err := gr.DB.Model(&c).Association("Users").Find(&c.Users); err != nil {
			return domain.Category{}, err
		}
	}
	return c.ToDto(), nil
}

func (gr *CategoryRepository) Delete(ctx context.Context, id uint) error {
	// Fetch the category along with its associations
	if err := gr.DB.Exec(
		"DELETE FROM daps_category_users WHERE category_id = ?", id).Error; err != nil {
		return err
	}

	var category Category
	if result := gr.DB.Delete(&category, id); result.Error != nil {
		return result.Error
	}

	return nil
}

func (gr *CategoryRepository) List(ctx context.Context, ids *[]uint, filters *map[string]interface{}) ([]domain.Category, error) {
	var cs []Category
	var cats []domain.Category

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

	if result := query.Find(&cs); result.Error != nil {
		return []domain.Category{}, result.Error
	}

	for _, c := range cs {
		cs := c.ToDto()
		cats = append(cats, cs)
	}
	return cats, nil
}

func (gr *CategoryRepository) Update(ctx context.Context, id uint, filters *map[string]interface{}) error {
	var c Category
	c.ID = id
	if result := gr.DB.Model(&c).Updates(*filters); result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *CategoryRepository) Share(ctx context.Context, id uint, u domain.User) error {
	var c Category
	if result := gr.DB.Preload("Users").First(&c, id); result.Error != nil {
		return result.Error
	}

	if c.Users == nil {
		c.Users = make([]User, 0)
	}

	// Check if the user is already in the category
	if !isUserInCategory(c.Users, u) {
		query := fmt.Sprintf("INSERT INTO daps_category_users (category_id, user_id) VALUES (%d, %d)", id, u.ID)
		if err := gr.DB.Exec(query).Error; err != nil {
			return err
		}
		// Update the category to be shared
		query = fmt.Sprintf("UPDATE daps_categories SET shared = TRUE WHERE id = %d", id)
		if err := gr.DB.Exec(query).Error; err != nil {
			return err
		}
	}
	return nil
}

func isUserInCategory(users []User, u domain.User) bool {
	for _, user := range users {
		if user.ID == u.ID {
			return true
		}
	}
	return false
}

func (gr *CategoryRepository) Unshare(ctx context.Context, id uint, u domain.User) error {
	if err := gr.DB.Exec(
		"DELETE FROM daps_category_users WHERE user_id = ? AND category_id = ?", u.ID, id).Error; err != nil {
		return err
	}

	var c Category
	if result := gr.DB.Preload("Users").First(&c, id); result.Error != nil {
		return result.Error
	}

	if c.Users == nil || len(c.Users) == 1 {
		query := fmt.Sprintf("UPDATE daps_categories SET shared = FALSE WHERE id = %d", id)
		if err := gr.DB.Exec(query).Error; err != nil {
			return err
		}
	}

	return nil
}

func (gr *CategoryRepository) GetSummary(ctx context.Context, id uint) ([]domain.CategorySummary, error) {
	var cs []domain.CategorySummary
	query := fmt.Sprintf("SELECT daps_category_users.category_id, daps_categories.id, daps_categories.name, daps_categories.owner_id, daps_categories.shared, SUM(CASE WHEN daps_todos.completed = FALSE AND daps_todos.priority = 5 then 1 else 0 END) as highest_priority_tasks, SUM(CASE WHEN daps_todos.completed = FALSE then 1 else 0 END) as tasks FROM daps_category_users INNER JOIN daps_categories ON daps_categories.id = daps_category_users.category_id LEFT JOIN daps_todos ON daps_todos.category_id = daps_category_users.category_id WHERE user_id = %d GROUP BY category_id", id)
	result := gr.DB.Raw(query).Scan(&cs)
	if result.Error != nil {
		return cs, result.Error
	}
	return cs, nil
}

func (gr *CategoryRepository) ListCategoryUsers(ctx context.Context, id uint) ([]domain.CategoryUser, error) {
	var cu []domain.CategoryUser
	query := fmt.Sprintf("SELECT DISTINCT daps_category_users.user_id, deselflopment_users.email, CASE WHEN daps_category_users.user_id = daps_categories.owner_id THEN 1 ELSE 0 END AS is_owner FROM daps_category_users LEFT JOIN deselflopment_users ON daps_category_users.user_id = deselflopment_users.id LEFT JOIN daps_categories ON daps_category_users.category_id = daps_categories.id WHERE daps_category_users.category_id = %d", id)
	result := gr.DB.Raw(query).Scan(&cu)
	if result.Error != nil {
		return cu, result.Error
	}
	return cu, nil
}

type CategoryRepository struct {
	*gorm.DB
}

func NewGormCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db}
}
