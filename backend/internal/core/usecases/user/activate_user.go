package user

import (
	"context"
	"github.com/aghex70/daps/internal/core/services/user"
	requests "github.com/aghex70/daps/internal/ports/requests/user"
	"log"
)

type ActivateUserUseCase struct {
	UserService user.Service
	logger      *log.Logger
}

func (uc *ActivateUserUseCase) Execute(ctx context.Context, r requests.ActivateUserRequest, userID uint) error {
	if err := uc.UserService.Activate(ctx, userID, r.ActivationCode); err != nil {
		return err
	}
	return nil
}

func NewActivateUserUseCase(userService user.Service, logger *log.Logger) *ActivateUserUseCase {
	return &ActivateUserUseCase{
		UserService: userService,
		logger:      logger,
	}
}
