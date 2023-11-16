package user

import (
	"context"
	"github.com/aghex70/daps/internal/core/services/user"
	"github.com/aghex70/daps/internal/pkg"
	"log"
)

type CheckAdminUseCase struct {
	UserService user.Service
	logger      *log.Logger
}

func (uc *CheckAdminUseCase) Execute(ctx context.Context, id uint) (bool, error) {
	u, err := uc.UserService.Get(ctx, id)
	if err != nil {
		return false, err
	}

	if !u.Admin {
		return false, pkg.UnauthorizedError
	}

	return true, nil
}

func NewCheckAdminUseCase(userService user.Service, logger *log.Logger) *CheckAdminUseCase {
	return &CheckAdminUseCase{
		UserService: userService,
		logger:      logger,
	}
}
