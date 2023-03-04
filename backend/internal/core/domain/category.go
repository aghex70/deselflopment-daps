package domain

type Category struct {
	Id          int    `json:"id"`
	OwnerId     int    `json:"owner_id"`
	Name        string `json:"name"`
	Shared      *bool  `json:"shared"`
	Notifiable  bool   `json:"notifiable"`
	Custom      bool   `json:"custom"`
	Description string `json:"description"`
	Users       []User `json:"users"`
}
