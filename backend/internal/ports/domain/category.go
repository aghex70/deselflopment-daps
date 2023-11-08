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

//func (c Category) FromDto() repository.Category {
//	return repository.Category{
//		Name:        c.Name,
//		Description: c.Description,
//		OwnerID:     c.OwnerID,
//		//Users:       c.Users,
//		//Todos:       c.Todos,
//		Shared:     c.Shared,
//		Notifiable: c.Notifiable,
//		Custom:     c.Custom,
//	}
//}
