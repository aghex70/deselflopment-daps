package domain

import "time"

type Email struct {
	Id           int       `json:"id"`
	From         string    `json:"from"`
	To           string    `json:"to"`
	Recipient    string    `json:"recipient"`
	Subject      string    `json:"subject"`
	Body         string    `json:"body"`
	User         int       `json:"user"`
	CreationDate time.Time `json:"creation_date"`
	Sent         bool      `json:"sent"`
	Error        string    `json:"error"`
}
