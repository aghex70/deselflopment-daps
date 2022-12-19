package domain

type Category struct {
	Id                int    `json:"id"`
	OwnerId           int    `json:"owner_id"`
	Name              string `json:"name"`
	Shared            *bool  `json:"shared"`
	Custom            bool   `json:"custom"`
	Description       string `json:"description"`
	InternationalName string `json:"international_name"`
	Users             []User `json:"users"`
}
