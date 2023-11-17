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

	user, err := uc.UserService.Get(ctx, r.UserID)
	if err != nil {
		return domain.User{}, err
	}

	if !u.Admin || user.ID != r.UserID {
		return domain.User{}, pkg.UnauthorizedError
	}
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func NewGetUserUseCase(userService user.Service, logger *log.Logger) *GetUserUseCase {
	return &GetUserUseCase{
		UserService: userService,
		logger:      logger,
	}
}
