package domain

type UserCategoryRelationship struct {
	ID         int `json:"id"`
	CategoryID int `json:"category_id"`
	UserID     int `json:"user_id"`
}
