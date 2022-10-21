package domain

import "time"

type Priority int

const (
	Lowest Priority = iota
	Low
	Medium
	High
	Highest
)

type Todo struct {
	Active       bool       `json:"active"`
	Category     int        `json:"category_id"`
	Completed    bool       `json:"completed"`
	CreationDate time.Time  `json:"creation_date"`
	Description  string     `json:"description"`
	EndDate      *time.Time `json:"end_date"`
	ID           int        `json:"id"`
	Link         string     `json:"link"`
	Name         string     `json:"name"`
	Priority     Priority   `json:"priority"`
	Recurring    bool       `json:"recurring"`
	StartDate    *time.Time `json:"start_date"`
}

type Summmary struct {
	Summary []CategorySummary `json:"summary"`
}

type CategorySummary struct {
	ID                   int    `json:"id"`
	Name                 string `json:"name"`
	Tasks                int    `json:"tasks"`
	HighestPriorityTasks int    `json:"highest_priority_tasks"`
}
