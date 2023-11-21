package user

import (
	"context"
	"github.com/aghex70/daps/internal/core/services/user"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	"log"
)

type ListUsersUseCase struct {
	UserService user.Service
	logger      *log.Logger
}

func (uc *ListUsersUseCase) Execute(ctx context.Context, fields *map[string]interface{}, userID uint) ([]domain.User, error) {
	u, err := uc.UserService.Get(ctx, userID)
	if err != nil {
		return []domain.User{}, err
	}
	if !u.Active {
		return []domain.User{}, pkg.InactiveUserError
	}

	if !u.Admin {
		return []domain.User{}, pkg.UnauthorizedError
	}

	us, err := uc.UserService.List(ctx, fields)
	if err != nil {
		return []domain.User{}, err
	}
	return us, nil
}

func NewListUsersUseCase(userService user.Service, logger *log.Logger) *ListUsersUseCase {
	return &ListUsersUseCase{
		UserService: userService,
		logger:      logger,
	}
}
