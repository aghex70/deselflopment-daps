package domain

import (
	repository "github.com/aghex70/daps/internal/repositories/gorm"
	"time"
)

type Email struct {
	ID        uint
	CreatedAt time.Time
	From      string
	To        string
	Subject   string
	Body      string
	UserID    uint
	Sent      bool
	Source    string
	Error     *string
}

func (e Email) FromDto() repository.Email {
	return repository.Email{
		From:    e.From,
		To:      e.To,
		Subject: e.Subject,
		Body:    e.Body,
		UserID:  e.UserID,
		Sent:    e.Sent,
		Source:  e.Source,
		Error:   e.Error,
	}
}
