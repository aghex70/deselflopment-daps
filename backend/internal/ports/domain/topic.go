package domain

import (
	"time"
)

type Topic struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	OwnerID   uint      `json:"owner_id"`
}
