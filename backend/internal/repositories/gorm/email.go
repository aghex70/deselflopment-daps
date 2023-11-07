package category

import (
	"context"
	"github.com/aghex70/daps/internal/core/domain"
	"gorm.io/gorm"
)

type Email struct {
	gorm.Model
	From    string
	To      string
	Subject string
	Body    string
	UserID  uint
	Sent    bool
	Source  string
	Error   *string
}

func (e Email) ToDto() domain.Email {
	return domain.Email{
		ID:        e.ID,
		CreatedAt: e.CreatedAt,
		From:      e.From,
		To:        e.To,
		Subject:   e.Subject,
		Body:      e.Body,
		UserID:    e.UserID,
		Sent:      e.Sent,
		Source:    e.Source,
		Error:     e.Error,
	}
}

func (Email) TableName() string {
	return "daps_emails"
}

func (gr *GormRepository) GetEmail(ctx context.Context, id int) (domain.Email, error) {
	return domain.Email{}, nil
}

func (gr *GormRepository) GetEmails(ctx context.Context) ([]domain.Email, error) {
	return []domain.Email{}, nil
}

func (gr *GormRepository) CreateEmail(ctx context.Context, e domain.Email) (domain.Email, error) {
	return domain.Email{}, nil
}

func (gr *GormRepository) UpdateEmail(ctx context.Context, e domain.Email) error {
	return nil
}

func (gr *GormRepository) DeleteEmail(ctx context.Context, id int) error {
	return nil
}
