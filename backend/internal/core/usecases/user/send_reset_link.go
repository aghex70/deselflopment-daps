package user

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	requests "github.com/aghex70/daps/internal/ports/requests/user"
	"github.com/aghex70/daps/internal/ports/services/email"
	"github.com/aghex70/daps/internal/ports/services/user"
	utils "github.com/aghex70/daps/utils/email"
	"log"
)

type SendResetLinkUseCase struct {
	UserService  user.Servicer
	EmailService email.Servicer
	logger       *log.Logger
}

func (uc *SendResetLinkUseCase) Execute(ctx context.Context, r requests.ResetLinkRequest, userID uint) error {
	u, err := uc.UserService.GetByEmail(ctx, r.Email)
	if err != nil {
		return err
	}

	// Check if the user is the same as the one requesting the reset link
	if u.ID != userID {
		return pkg.UnauthorizedError
	}

	e := utils.GenerateResetPasswordEmail(u)
	s, err := uc.EmailService.Send(ctx, e)
	if !s && err != nil {
		return err
	}

	return nil
}

func NewSendResetLinkUseCase(userService user.Servicer, emailService email.Servicer, logger *log.Logger) *SendResetLinkUseCase {
	return &SendResetLinkUseCase{
		UserService:  userService,
		EmailService: emailService,
		logger:       logger,
	}
}
