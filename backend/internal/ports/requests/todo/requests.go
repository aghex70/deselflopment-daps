package requests

type CreateTodoRequest struct {
	CategoryID  uint   `json:"category_id" validate:"required"`
	Description string `json:"description"`
	Link        string `json:"link"`
	Name        string `json:"name" validate:"required"`
	Recurring   bool   `json:"recurring"`
	Recurrency  string `json:"recurrency"`
	Priority    uint   `json:"priority" validate:"required,gte=1,lte=5"`
}

type DeleteTodoRequest struct {
	TodoID uint `json:"todo_id"`
}

type GetTodoRequest struct {
	TodoID uint `json:"todo_id"`
}

type UpdateTodoRequest struct {
	Description *string `json:"description"`
	Link        *string `json:"link"`
	Name        *string `json:"name"`
	Recurring   *bool   `json:"recurring"`
	Recurrency  *string `json:"recurrency"`
	Priority    *uint   `json:"priority"`
	Suggestable *bool   `json:"suggestable"`
	TodoID      uint    `json:"todo_id"`
}

type ListTodosRequest struct {
	CategoryID uint `json:"category_id" validate:"required"`
}
