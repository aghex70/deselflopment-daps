package domain

import (
	"time"
)

type Category struct {
	ID          uint
	CreatedAt   time.Time
	Name        string
	Description *string
	OwnerID     uint
	Users       *[]User
	Todos       *[]Todo
	Shared      bool
	Notifiable  bool
	Custom      bool
}
