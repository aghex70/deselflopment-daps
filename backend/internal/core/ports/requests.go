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
	Category     string        `json:"category" validate:"required"`
	Description  string        `json:"description"`
	Duration     time.Duration `json:"duration" validate:"required"`
	Link         string        `json:"link"`
	Name         string        `json:"name" validate:"required"`
	Prerequisite string        `json:"prerequisite"`
	Priority     uint32        `json:"priority" validate:"required"`
}

type CompleteTodoRequest struct {
	TodoId uint64 `json:"to_do_id"`
}

type DeleteTodoRequest struct {
	TodoId uint64 `json:"to_do_id"`
}

type GetTodoRequest struct {
	TodoId uint64 `json:"to_do_id"`
}

type ListTodosRequest struct {
	TodoId uint `json:"inventado"`
}

type CreateUserRequest struct {
	Email          string `json:"email" validate:"required,email"`
	Password       string `json:"password" validate:"required,min=14"`
	RepeatPassword string `json:"repeat_password" validate:"required,min=14"`
}

type DeleteUserRequest struct {
}

type LoginUserRequest struct {
}

type LogoutUserRequest struct {
}
