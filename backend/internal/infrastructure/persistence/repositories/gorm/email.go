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

func (Email) TableName() string {
	return "daps_emails"
}

type GormEmailRepository struct {
	*gorm.DB
}

func NewGormEmailRepository(db *gorm.DB) *GormEmailRepository {
	return &GormEmailRepository{DB: db}
}

func (gr *GormEmailRepository) Get(ctx context.Context, id uint) (domain.Email, error) {
	return domain.Email{}, nil
}

func (gr *GormEmailRepository) List(ctx context.Context, filters *map[string]interface{}) ([]domain.Email, error) {
	return []domain.Email{}, nil
}

func (gr *GormEmailRepository) Create(ctx context.Context, e domain.Email) (domain.Email, error) {
	return domain.Email{}, nil
}

func (gr *GormEmailRepository) Update(ctx context.Context, e domain.Email) error {
	return nil
}

func (gr *GormEmailRepository) Delete(ctx context.Context, id uint) error {
	return nil
}
