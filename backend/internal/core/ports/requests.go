package ports

type CreateCategoryRequest struct {
	Description       string `json:"description"`
	Name              string `json:"name" validate:"required"`
	InternationalName string `json:"international_name"`
}

type DeleteCategoryRequest struct {
	CategoryId int64 `json:"category_id"`
}

type GetCategoryRequest struct {
	CategoryId int64 `json:"category_id"`
}

type UpdateCategoryRequest struct {
	CategoryId        int64  `json:"category_id"`
	Description       string `json:"description"`
	Name              string `json:"name"`
	InternationalName string `json:"international_name"`
	Shared            *bool  `json:"shared"`
	Email             string `json:"email"`
}

type CreateTodoRequest struct {
	Category    int    `json:"category_id" validate:"required"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Name        string `json:"name" validate:"required"`
	Recurring   bool   `json:"recurring"`
	Priority    int32  `json:"priority" validate:"required,gte=1,lte=5"`
}

type CompleteTodoRequest struct {
	Category int   `json:"category_id" validate:"required"`
	TodoId   int64 `json:"todo_id"`
}

type ActivateTodoRequest struct {
	Category int   `json:"category_id" validate:"required"`
	TodoId   int64 `json:"todo_id"`
}

type StartTodoRequest struct {
	Category int   `json:"category_id" validate:"required"`
	TodoId   int64 `json:"todo_id"`
}

type DeleteTodoRequest struct {
	Category int   `json:"category_id" validate:"required"`
	TodoId   int64 `json:"todo_id"`
}

type GetTodoRequest struct {
	Category int   `json:"category_id" validate:"required"`
	TodoId   int64 `json:"todo_id"`
}

type UpdateTodoRequest struct {
	Category    int    `json:"category_id" validate:"required"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Name        string `json:"name"`
	Recurring   bool   `json:"recurring"`
	Priority    int32  `json:"priority" validate:"required,gte=1,lte=5"`
	TodoId      int64  `json:"todo_id"`
}

type ListTodosRequest struct {
	Category int `json:"category_id" validate:"required"`
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=4"`
}

type RefreshTokenRequest struct {
}
