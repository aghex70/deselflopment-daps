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
	ID          uint
	CreatedAt   time.Time
	Name        string
	Description *string
	Completed   bool
	CompletedAt *time.Time
	Active      bool
	Priority    Priority
	CategoryID  uint
	Link        *string
	Recurring   bool
	Recurrency  *string
	StartedAt   *time.Time
	Suggestable bool
	Suggested   bool
	SuggestedAt *time.Time
	UserID      uint
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
