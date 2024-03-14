package domain

import (
	"time"
)

type Priority int

const (
	Lowest Priority = iota + 1
	Low
	Medium
	High
	Highest
)

type Todo struct {
	ID          uint       `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	Name        string     `json:"name"`
	Description *string    `json:"description"`
	Completed   bool       `json:"completed"`
	CompletedAt *time.Time `json:"completed_at"`
	Active      bool       `json:"active"`
	Priority    Priority   `json:"priority"`
	CategoryID  uint       `json:"category_id"`
	Link        *string    `json:"link"`
	Recurring   bool       `json:"recurring"`
	Recurrency  *string    `json:"recurrency"`
	StartedAt   *time.Time `json:"started_at"`
	Suggestable bool       `json:"suggestable"`
	Suggested   bool       `json:"suggested"`
	SuggestedAt *time.Time `json:"suggested_at"`
	OwnerID     uint       `json:"owner_id"`
}
