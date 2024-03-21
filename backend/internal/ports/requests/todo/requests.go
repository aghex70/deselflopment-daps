package requests

type CreateTodoRequest struct {
	CategoryID  uint    `json:"category_id" validate:"required"`
	Description string  `json:"description"`
	Link        string  `json:"link"`
	Name        string  `json:"name" validate:"required"`
	Recurring   bool    `json:"recurring"`
	Recurrency  *int    `json:"recurrency"`
	Priority    uint    `json:"priority" validate:"required,gte=1,lte=5"`
	TargetDate  *string `json:"target_date"`
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
	Recurring   bool    `json:"recurring"`
	Recurrency  *int    `json:"recurrency"`
	Priority    *uint   `json:"priority"`
	Suggestable *bool   `json:"suggestable"`
	TodoID      uint    `json:"todo_id"`
	TargetDate  *string `json:"target_date"`
}

type ListTodosRequest struct {
	CategoryID uint `json:"category_id" validate:"required"`
}
