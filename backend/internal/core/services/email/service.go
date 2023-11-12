package email

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
	"github.com/aghex70/daps/internal/ports/repositories/email"
	"log"
)

type Service struct {
	logger     *log.Logger
	repository email.Repository
}

func (s Service) Send(ctx context.Context, e domain.Email) (bool, error) {
	// Create email in database
	// Send email
	// Update email in database
	return true, nil
}

func NewEmailService(r email.Repository, logger *log.Logger) Service {
	return Service{
		logger:     logger,
		repository: r,
	}
}
