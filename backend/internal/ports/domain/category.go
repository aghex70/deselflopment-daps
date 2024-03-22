package domain

import (
	"time"
)

type Category struct {
	ID          uint      `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	OwnerID     uint      `json:"owner_id"`
	Users       []User    `json:"users"`
	Todos       []Todo    `json:"todos"`
	Shared      bool      `json:"shared"`
	Notifiable  bool      `json:"notifiable"`
	Custom      bool      `json:"custom"`
}

type FilteredCategory struct {
	ID          uint    `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Shared      bool    `json:"shared"`
	Notifiable  bool    `json:"notifiable"`
	Custom      bool    `json:"custom"`
}
