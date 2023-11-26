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
	var category Category
	if result := gr.DB.Preload("Users").First(&category, id); result.Error != nil {
		return result.Error
	}

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
	//var c Category
	//if result := gr.DB.Preload("Users").First(&c, id); result.Error != nil {
	//	return result.Error
	//}
	//
	//for _, user := range c.Users {
	//	fmt.Printf("User: %+v\n", user)
	//}
	//// User is not in the category, so add them
	//if !isUserInCategory(c.Users, u) {
	//	c.Shared = true
	//	c.Users = append(c.Users, UserFromDto(u))
	//	if err := gr.DB.Association("Users").Replace(c.Users); err != nil {
	//		return err
	//	}
	//}
	//
	////fmt.Printf("New user list: %+v\n", c.Users)
	////
	////	Save the updated category, including the many-to-many association
	////if result := gr.DB.Save(&c); result.Error != nil {
	////	return result.Error
	////}
	////}
	//
	//return nil
	//var c Category
	//if result := gr.DB.Preload("Users").First(&c, id); result.Error != nil {
	//	return result.Error
	//}
	//
	//fmt.Printf("Category: %+v\n", c)
	//fmt.Printf("Category users: %+v\n", c.Users)
	//
	//if c.Users == nil {
	//	c.Users = make([]User, 0)
	//}
	//// Check if the user is already in the category
	//if !isUserInCategory(c.Users, u) {
	//	// Initialize the Users slice if it's nil
	//
	//	// User is not in the category, so add them
	//	//c.Users = append(c.Users, UserFromDto(u))
	//
	//	// Set the shared flag to true
	//	c.Shared = true
	//	c.Users = append(c.Users, UserFromDto(u))
	//	fmt.Println("Category users length:", len(c.Users))
	//	fmt.Printf("New Category users: %+v\n", c.Users)
	//
	//	// Replace the users association with the updated slice
	//	gr.DB.Session(&gorm.Session{FullSaveAssociations: false})
	//	if err := gr.DB.Model(&c).Association("Users").Replace(c.Users); err != nil {
	//		return err
	//	}
	//
	//	fmt.Printf("New user list: %+v\n", c.Users)
	//	////Save the updated category
	//	//if result := gr.DB.Save(&c); result.Error != nil {
	//	//	return result.Error
	//	//}
	//}
	////}
	//
	//return nil
	var c Category
	if result := gr.DB.Preload("Users").First(&c, id); result.Error != nil {
		return result.Error
	}

	fmt.Printf("c.Users: %+v\n", c.Users)
	users := append(c.Users, UserFromDto(u))
	fmt.Printf("Users: %+v\n", users)
	// Initialize the Users slice if it's nil
	if c.Users == nil {
		c.Users = make([]User, 0)
	}

	// Remove existing associations
	if err := gr.DB.Model(&c).Association("Users").Replace(nil); err != nil {
		return err
	}

	// Add new users
	fmt.Println("Length users:", len(users))
	for _, u := range users {
		//user := UserFromDto(u)
		c.Users = append(c.Users, u)
	}

	// Set the shared flag to true
	c.Shared = true

	// Save the updated category
	if result := gr.DB.Save(&c); result.Error != nil {
		return result.Error
	}

	// Replace the users association with the updated slice
	if err := gr.DB.Model(&c).Association("Users").Replace(c.Users); err != nil {
		return err
	}

	//// Save the updated category
	//if result := gr.DB.Save(&c); result.Error != nil {
	//	return result.Error
	//}

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
	var c Category
	if result := gr.DB.Preload("Users").First(&c, id); result.Error != nil {
		return result.Error
	}

	// Check if the user is associated with the category
	index := findUserIndex(c.Users, u)
	if index != -1 {
		// User is in the category, so remove them
		c.Users = append(c.Users[:index], c.Users[index+1:]...)

		// If the Users slice is empty, set the Shared flag to false
		if len(c.Users) == 0 {
			c.Shared = false
		}

		if result := gr.DB.Save(&c); result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func findUserIndex(users []User, u domain.User) int {
	for i, user := range users {
		if user.ID == u.ID {
			return i
		}
	}
	return -1
}

type CategoryRepository struct {
	*gorm.DB
}

func NewGormCategoryRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{db}
}
