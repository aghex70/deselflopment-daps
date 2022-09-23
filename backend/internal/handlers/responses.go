package handlers

import (
	"github.com/aghex70/daps/internal/core/domain"
	"time"
)

type TodoResponse struct {
	Active  bool       `json:"active"`
	EndDate *time.Time `json:"end_date"`
	//Category     Category
	Completed    bool          `json:"completed"`
	CreationDate time.Time     `json:"creation_date"`
	Description  string        `json:"description"`
	Duration     time.Duration `json:"duration"`
	ID           int           `json:"id"`
	Link         string        `json:"link"`
	Name         string        `json:"name"`
	//Prerequisite *Todo
	//Priority  Priority
	StartDate *time.Time `json:"startDate"`
	//User         User
}

type ListTodosResponse struct {
	Todos []domain.Todo `json:"todos"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}
