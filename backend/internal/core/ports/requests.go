package ports

import (
	"time"
)

type CreateCategoryRequest struct {
}

type DeleteCategoryRequest struct {
}

type GetCategoryRequest struct {
}

type ListCategoriesRequest struct {
}

type CreateTodoRequest struct {
	Category     string        `json:"category_id"`
	Description  string        `json:"description"`
	Duration     time.Duration `json:"duration" validate:"required"`
	Link         string        `json:"link"`
	Name         string        `json:"name" validate:"required"`
	Prerequisite string        `json:"prerequisite"`
	Priority     uint32        `json:"priority" validate:"required"`
}

type CompleteTodoRequest struct {
	TodoId int64 `json:"todo_id"`
}

type StartTodoRequest struct {
	TodoId int64 `json:"todo_id"`
}

type DeleteTodoRequest struct {
	TodoId int64 `json:"todo_id"`
}

type GetTodoRequest struct {
	TodoId int64 `json:"todo_id"`
}

type UpdateTodoRequest struct {
	Category     string        `json:"category_id"`
	Description  string        `json:"description"`
	Duration     time.Duration `json:"duration" validate:"required"`
	Link         string        `json:"link"`
	Name         string        `json:"name" validate:"required"`
	Prerequisite string        `json:"prerequisite"`
	Priority     uint32        `json:"priority" validate:"required"`
	TodoId       int64         `json:"todo_id"`
}

// CreateUserRequest Users
type CreateUserRequest struct {
	Email          string `json:"email" validate:"required,email"`
	Password       string `json:"password" validate:"required,min=14"`
	RepeatPassword string `json:"repeat_password" validate:"required,min=14"`
}

// LoginUserRequest
type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=14"`
}

// RefreshTokenRequest
type RefreshTokenRequest struct {
}
