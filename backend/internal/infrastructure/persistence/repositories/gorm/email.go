package gorm

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
	"gorm.io/gorm"
)

type Email struct {
	gorm.Model
	Subject      string
	Body         string
	From         string `gorm:"column:from_email"`
	Source       string
	To           string `gorm:"column:to_email"`
	Recipient    string
	Sent         bool
	ErrorMessage *string
	UserID       uint
}

func (e Email) ToDto() domain.Email {
	return domain.Email{
		ID:           e.ID,
		CreatedAt:    e.CreatedAt,
		From:         e.From,
		To:           e.To,
		Subject:      e.Subject,
		Body:         e.Body,
		UserID:       e.UserID,
		Sent:         e.Sent,
		Source:       e.Source,
		ErrorMessage: e.ErrorMessage,
		Recipient:    e.Recipient,
	}
}

func EmailFromDto(e domain.Email) Email {
	return Email{
		From:         e.From,
		To:           e.To,
		Subject:      e.Subject,
		Body:         e.Body,
		UserID:       e.UserID,
		Sent:         e.Sent,
		Source:       e.Source,
		ErrorMessage: e.ErrorMessage,
	}
}

func EmailsFromDto(es []domain.Email) []Email {
	var emails []Email
	for _, e := range es {
		emails = append(emails, EmailFromDto(e))
	}
	return emails
}

func (Email) TableName() string {
	return "deselflopment_emails"
}

type EmailRepository struct {
	*gorm.DB
}

func NewGormEmailRepository(db *gorm.DB) *EmailRepository {
	return &EmailRepository{db}
}

func (gr *EmailRepository) Create(ctx context.Context, e domain.Email) (domain.Email, error) {
	email := EmailFromDto(e)
	result := gr.DB.Create(&email)
	if result.Error != nil {
		return domain.Email{}, result.Error
	}
	return e, nil
}

func (gr *EmailRepository) Get(ctx context.Context, id uint) (domain.Email, error) {
	var e Email
	result := gr.DB.First(&e, id)
	if result.Error != nil {
		return domain.Email{}, result.Error
	}
	return e.ToDto(), nil
}

func (gr *EmailRepository) Delete(ctx context.Context, id uint) error {
	result := gr.DB.Delete(&Email{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (gr *EmailRepository) List(ctx context.Context, filters *map[string]interface{}) ([]domain.Email, error) {
	var es []Email
	result := gr.DB.Find(&es, filters)
	if result.Error != nil {
		return []domain.Email{}, result.Error
	}
	var emails []domain.Email
	for _, e := range es {
		emails = append(emails, e.ToDto())
	}
	return emails, nil
}

func (gr *EmailRepository) Update(ctx context.Context, e domain.Email) error {
	result := gr.DB.Save(&e)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
