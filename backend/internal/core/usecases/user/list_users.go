package user

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	"github.com/aghex70/daps/internal/ports/services/user"
	"log"
)

type ListUsersUseCase struct {
	UserService user.Servicer
	logger      *log.Logger
}

func (uc *ListUsersUseCase) Execute(ctx context.Context, filters *map[string]interface{}, userID uint) ([]domain.User, error) {
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

	us, err := uc.UserService.List(ctx, filters)
	if err != nil {
		return []domain.User{}, err
	}
	return us, nil
}

func NewListUsersUseCase(userService user.Servicer, logger *log.Logger) *ListUsersUseCase {
	return &ListUsersUseCase{
		UserService: userService,
		logger:      logger,
	}
}
