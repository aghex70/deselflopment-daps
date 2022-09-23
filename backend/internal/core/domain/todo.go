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
	Active       bool
	EndDate      *time.Time
	Category     int
	Completed    bool
	CreationDate time.Time
	Description  string
	Duration     time.Duration
	Link         string
	Name         string
	//Prerequisite *Todo
	Priority  Priority
	StartDate *time.Time
	User      int
}
