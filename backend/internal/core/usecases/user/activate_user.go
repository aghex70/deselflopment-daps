package user

import (
	"context"
	requests "github.com/aghex70/daps/internal/ports/requests/user"
	"github.com/aghex70/daps/internal/ports/services/user"
	"log"
)

type ActivateUserUseCase struct {
	UserService user.Servicer
	logger      *log.Logger
}

func (uc *ActivateUserUseCase) Execute(ctx context.Context, r requests.ActivateUserRequest, userID uint) error {
	if err := uc.UserService.Activate(ctx, userID, r.ActivationCode); err != nil {
		return err
	}
	return nil
}

func NewActivateUserUseCase(userService user.Servicer, logger *log.Logger) *ActivateUserUseCase {
	return &ActivateUserUseCase{
		UserService: userService,
		logger:      logger,
	}
}
