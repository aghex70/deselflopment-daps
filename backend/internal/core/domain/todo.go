package domain

import "time"

type Priority uint

const (
	Lowest Priority = iota
	Low
	Medium
	High
	Highest
)

type Todo struct {
	Active       bool          `json:"active"`
	Category     int           `json:"category_id"`
	Completed    bool          `json:"completed"`
	CreationDate time.Time     `json:"creation_date"`
	Description  string        `json:"description"`
	Duration     time.Duration `json:"duration"`
	EndDate      *time.Time    `json:"end_date"`
	ID           int           `json:"id"`
	Link         string        `json:"link"`
	Name         string        `json:"name"`
	Prerequisite int           `json:"prerequisite_id"`
	Priority     Priority      `json:"priority"`
	Recurring    bool          `json:"recurring"`
	StartDate    *time.Time    `json:"start_date"`
	User         int           `json:"user_id"`
}
