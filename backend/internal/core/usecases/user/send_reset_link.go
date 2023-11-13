package user

import (
	"context"
	"github.com/aghex70/daps/internal/core/services/email"
	"github.com/aghex70/daps/internal/core/services/user"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	requests "github.com/aghex70/daps/internal/ports/requests/user"
	"log"
)

type SendResetLinkUseCase struct {
	UserService  user.Service
	EmailService email.Service
	logger       *log.Logger
}

func (uc *SendResetLinkUseCase) Execute(ctx context.Context, r requests.ResetLinkRequest) error {
	u, err := uc.UserService.GetByEmail(ctx, r.Email)
	if err != nil {
		return err
	}

	e := domain.Email{
		Subject:   "📣 DAPS - Activate your account 📣",
		Body:      "In order to reset your password, please follow this link: " + pkg.ResetPasswordLink + u.ResetPasswordCode,
		From:      pkg.FromEmail,
		Source:    pkg.ProjectName,
		To:        u.Name,
		Recipient: u.Email,
		UserID:    u.ID,
	}

	s, err := uc.EmailService.Send(ctx, e)
	if !s && err != nil {
		return err
	}

	return nil
}

func NewSendResetLinkUseCase(userService user.Service, emailService email.Service, logger *log.Logger) *SendResetLinkUseCase {
	return &SendResetLinkUseCase{
		UserService:  userService,
		EmailService: emailService,
		logger:       logger,
	}
}
