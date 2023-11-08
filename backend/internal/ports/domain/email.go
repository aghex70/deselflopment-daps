package domain

import (
	"time"
)

type Email struct {
	ID        uint
	CreatedAt time.Time
	Source    string
	From      string
	Recipient string
	To        string
	Subject   string
	Body      string
	UserID    uint
	Sent      bool
	Error     *string
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
