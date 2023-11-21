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

type CompleteTodoRequest struct {
	TodoID uint `json:"todo_id"`
}

type ActivateTodoRequest struct {
	TodoID uint `json:"todo_id"`
}

type StartTodoRequest struct {
	TodoID uint `json:"todo_id"`
}

type DeleteTodoRequest struct {
	TodoID uint `json:"todo_id"`
}

type GetTodoRequest struct {
	//CategoryID uint `json:"category_id" validate:"required"`
	TodoID uint `json:"todo_id"`
}

type UpdateTodoRequest struct {
	Description string `json:"description"`
	Link        string `json:"link"`
	Name        string `json:"name"`
	Recurring   bool   `json:"recurring"`
	Recurrency  string `json:"recurrency"`
	Priority    uint   `json:"priority" validate:"required,gte=1,lte=5"`
	Suggestable bool   `json:"suggestable"`
	TodoID      uint   `json:"todo_id"`
}

type ListTodosRequest struct {
	CategoryID uint `json:"category_id" validate:"required"`
}
