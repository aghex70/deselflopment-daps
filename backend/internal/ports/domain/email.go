package domain

import (
	"time"
)

type Email struct {
	ID           uint
	CreatedAt    time.Time
	Subject      string
	Body         string
	From         string
	Source       string
	To           string
	Recipient    string
	Sent         bool
	ErrorMessage *string
	UserID       uint
}
