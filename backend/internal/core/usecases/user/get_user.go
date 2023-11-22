package user

import (
	"context"
	"github.com/aghex70/daps/internal/core/services/user"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	requests "github.com/aghex70/daps/internal/ports/requests/user"
	"log"
)

type GetUserUseCase struct {
	UserService user.Service
	logger      *log.Logger
}

func (uc *GetUserUseCase) Execute(ctx context.Context, r requests.GetUserRequest, userID uint) (domain.User, error) {
	u, err := uc.UserService.Get(ctx, userID)
	if err != nil {
		return domain.User{}, err
	}

	if !u.Active {
		return domain.User{}, pkg.InactiveUserError
	}

	if !u.Admin {
		return domain.User{}, pkg.UnauthorizedError
	}

	ur, err := uc.UserService.Get(ctx, r.UserID)
	if err != nil {
		return domain.User{}, err
	}

	if ur.ID != r.UserID {
		return domain.User{}, pkg.UnauthorizedError
	}
	return ur, nil
}

func NewGetUserUseCase(userService user.Service, logger *log.Logger) *GetUserUseCase {
	return &GetUserUseCase{
		UserService: userService,
		logger:      logger,
	}
}
