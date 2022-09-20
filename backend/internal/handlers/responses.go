package handlers

import (
	"github.com/aghex70/daps/internal/core/domain"
	"time"
)

type TodoResponse struct {
	Active  bool
	EndDate *time.Time
	//Category     Category
	Completed    bool
	CreationDate time.Time
	Description  string
	Duration     time.Duration
	ID           int
	Link         string
	Name         string
	//Prerequisite *Todo
	//Priority  Priority
	StartDate *time.Time
	//User         User
}

type ListTodosResponse struct {
	Todos []domain.Todo
}
