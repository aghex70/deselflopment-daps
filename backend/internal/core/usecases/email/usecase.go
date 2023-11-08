package email

import (
	"context"
	"github.com/aghex70/daps/internal/ports/repositories/email"
	"github.com/aghex70/daps/internal/ports/requests/email"
	"log"
	"net/http"
)

type Service struct {
	logger     *log.Logger
	repository email.Repository
}

func (s Service) Send(ctx context.Context, r *http.Request, req requests.SendEmailRequest) error {
	//TODO implement me
	panic("implement me")
}

func NewEmailService(r email.Repository, logger *log.Logger) Service {
	return Service{
		logger:     logger,
		repository: r,
	}
}
