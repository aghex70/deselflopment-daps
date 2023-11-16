package user

import (
	"context"
	"github.com/aghex70/daps/internal/core/services/user"
	requests "github.com/aghex70/daps/internal/ports/requests/user"
	"log"
)

type DeleteUserUseCase struct {
	UserService user.Service
	logger      *log.Logger
}

func (uc *DeleteUserUseCase) Execute(ctx context.Context, r requests.DeleteUserRequest) error {
	err := uc.UserService.Delete(ctx, r.UserID)
	if err != nil {
		return err
	}
	return nil
}

func NewDeleteUserUseCase(userService user.Service, logger *log.Logger) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		UserService: userService,
		logger:      logger,
	}
}