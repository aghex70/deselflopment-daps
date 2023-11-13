package user

import (
	"context"
	"github.com/aghex70/daps/internal/core/services/user"
	"log"
)

type ListUsersUseCase struct {
	UserService user.Service
	logger      *log.Logger
}

func (uc *ListUsersUseCase) Execute(ctx context.Context) error {
	_, err := uc.UserService.List(ctx, nil)
	if err != nil {
		return err
	}
	return nil
}

func NewListUsersUseCase(userService user.Service, logger *log.Logger) *ListUsersUseCase {
	return &ListUsersUseCase{
		UserService: userService,
		logger:      logger,
	}
}
