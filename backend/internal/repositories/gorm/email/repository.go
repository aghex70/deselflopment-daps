package email

import (
	"context"
	"database/sql"
	"time"

	"github.com/aghex70/daps/internal/core/domain"
	"gorm.io/gorm"
)

type EmailGormRepository struct {
	*gorm.DB
	SqlDb *sql.DB
}

type Email struct {
	Id           int       `gorm:"primaryKey;column:id"`
	Subject      string    `gorm:"column:subject"`
	Body         string    `gorm:"column:body"`
	Sent         bool      `gorm:"column:sent"`
	CreationDate time.Time `gorm:"column:creation_date"`
	Error        string    `gorm:"column:error"`
	UserId       int       `gorm:"column:user_id"`
}

type Tabler interface {
	TableName() string
}

func (Email) TableName() string {
	return "daps_emails"
}

func (gr *EmailGormRepository) Create(ctx context.Context, e domain.Email) (domain.Email, error) {
	ne := EmailFromDto(e)
	result := gr.DB.Create(&ne)
	if result.Error != nil {
		return domain.Email{}, result.Error
	}
	return ne.ToDto(), nil
}

func NewEmailGormRepository(db *gorm.DB) (*EmailGormRepository, error) {
	return &EmailGormRepository{
		DB: db,
	}, nil
}

func (e Email) ToDto() domain.Email {
	return domain.Email{
		Id:           e.Id,
		Subject:      e.Subject,
		Body:         e.Body,
		Sent:         e.Sent,
		CreationDate: e.CreationDate,
		Error:        e.Error,
		User:         e.UserId,
	}
}

func EmailFromDto(u domain.Email) Email {
	return Email{
		Id:           u.Id,
		Subject:      u.Subject,
		Body:         u.Body,
		Sent:         u.Sent,
		CreationDate: time.Now(),
		Error:        u.Error,
		UserId:       u.User,
	}
}
