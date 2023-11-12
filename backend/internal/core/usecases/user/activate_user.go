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

func (uc *ActivateUserUseCase) Execute(ctx context.Context, r requests.ActivateUserRequest) error {
	err := uc.UserService.Activate(ctx, r.ActivationCode)
	if err != nil {
		return err
	}
	return nil
}
