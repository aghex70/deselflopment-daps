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

//type Summmary struct {
//	Summary []CategorySummary `json:"summary"`
//}
//
//type CategorySummary struct {
//	ID                   int    `json:"id"`
//	Name                 string `json:"name"`
//	Tasks                int    `json:"tasks"`
//	HighestPriorityTasks int    `json:"highest_priority_tasks"`
//	OwnerID              int    `json:"owner_id"`
//	Shared               int    `json:"shared"`
//}
//
//type RemindSummary struct {
//	TodoName        string `json:"todo_name"`
//	TodoPriority    int    `json:"todo_priority"`
//	TodoDescription string `json:"todo_description"`
//	TodoLink        string `json:"todo_link"`
//	CategoryName    string `json:"category_name"`
//}
//
//type CategoryInfo struct {
//	KategoryName string `json:"kategory_name"`
//}
//
//type TodoInfo struct {
//	CategoryInfo
//	Todo
//}
