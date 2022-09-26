package domain

type Category struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	User              *int   `json:"user_id"`
	Custom            bool   `json:"custom"`
	Description       string `json:"description"`
	InternationalName string `json:"international_name"`
}
