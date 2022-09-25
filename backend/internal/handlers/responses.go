package handlers

import (
	"github.com/aghex70/daps/internal/core/domain"
	"time"
)

type TodoResponse struct {
	Name         string    `json:"name"`
	CreationDate time.Time `json:"creation_date"`
	Priority     string
	Category     string
	Completed    bool          `json:"completed"`
	Active       bool          `json:"active"`
	StartDate    *time.Time    `json:"startDate"`
	EndDate      *time.Time    `json:"end_date"`
	Description  string        `json:"description"`
	Duration     time.Duration `json:"duration"`
	ID           int           `json:"id"`
	Link         string        `json:"link"`
	//User         User
}

type ListTodosResponse struct {
	Todos []domain.Todo `json:"todos"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}
