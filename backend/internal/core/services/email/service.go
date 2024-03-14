package email

import (
	"context"
	adapter "github.com/aghex70/daps/internal/adapters/email"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	"github.com/aghex70/daps/internal/ports/repositories/email"
	"log"
)

type Service struct {
	logger          *log.Logger
	emailRepository email.Repository
}

func (s Service) Send(ctx context.Context, e domain.Email) (bool, error) {
	err := adapter.SendMail(e)
	if err != nil {
		e.Sent = false
		_, _err := s.emailRepository.Create(ctx, e)
		if _err != nil {
			log.Println(pkg.SaveEmailError.Error())
			return false, err
		}

	}
	e.Sent = true
	_, _err := s.emailRepository.Create(ctx, e)
	if _err != nil {
		log.Println(pkg.SaveEmailError.Error())
		return false, err
	}
	return true, err
}

func NewEmailService(r email.Repository, logger *log.Logger) Service {
	return Service{
		logger:          logger,
		emailRepository: r,
	}
}
