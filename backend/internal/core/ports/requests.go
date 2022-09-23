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
	TodoId uint64 `json:"todo_id"`
}

type DeleteTodoRequest struct {
	TodoId uint64 `json:"todo_id"`
}

type GetTodoRequest struct {
	TodoId uint64 `json:"todo_id"`
	UserId uint64 `json:"user_id"`
}

type ListTodosRequest struct {
	TodoId uint64 `json:"todo_id"`
	UserId uint64 `json:"user_id"`
}

type CreateUserRequest struct {
	Email          string `json:"email" validate:"required,email"`
	Password       string `json:"password" validate:"required,min=14"`
	RepeatPassword string `json:"repeat_password" validate:"required,min=14"`
}

type DeleteUserRequest struct {
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=14"`
}

type RefreshTokenRequest struct {
	AccessToken string `json:"access_token" validate:"required"`
}

type LogoutUserRequest struct {
}
