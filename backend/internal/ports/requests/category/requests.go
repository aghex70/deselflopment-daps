package requests

type CreateCategoryRequest struct {
	Description string `json:"description"`
	Name        string `json:"name" validate:"required"`
	Notifiable  bool   `json:"notifiable"`
}

type DeleteCategoryRequest struct {
	CategoryID uint `json:"category_id"`
}

type GetCategoryRequest struct {
	CategoryID uint `json:"category_id"`
}

type UpdateCategoryRequest struct {
	CategoryID  uint   `json:"category_id"`
	Description *string `json:"description"`
	Name        string `json:"name"`
	Notifiable  bool   `json:"notifiable"`
}
