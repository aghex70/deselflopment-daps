package domain

import (
	"time"
)

type Note struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Content   string    `json:"content"`
	Users     []User    `json:"users"`
	OwnerID   uint      `json:"owner_id"`
	Topics    []Topic   `json:"topics"`
	Subtopic  Topic     `json:"subtopic"`
	Shared    bool      `json:"shared"`
}
