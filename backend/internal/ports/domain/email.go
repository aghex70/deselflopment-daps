package domain

import (
	"time"
)

type Email struct {
	ID           uint
	CreatedAt    time.Time
	Subject      string
	Body         string
	From         string
	Source       string
	To           string
	Recipient    string
	Sent         bool
	ErrorMessage *string
	UserID       uint
}

//func (e Email) FromDto() repository.Email {
//	return repository.Email{
//		From:    e.From,
//		To:      e.To,
//		Subject: e.Subject,
//		Body:    e.Body,
//		UserID:  e.UserID,
//		Sent:    e.Sent,
//		Source:  e.Source,
//		Error:   e.Error,
//	}
//}
