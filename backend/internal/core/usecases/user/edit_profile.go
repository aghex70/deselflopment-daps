package user

import (
	"context"
	"errors"
	"github.com/aghex70/daps/internal/pkg"
	requests "github.com/aghex70/daps/internal/ports/requests/user"
	"github.com/aghex70/daps/internal/ports/services/user"
	common "github.com/aghex70/daps/utils"
	"gorm.io/gorm"
	"log"
)

type EditProfileUseCase struct {
	UserService user.Servicer
	logger      *log.Logger
}

func (uc *EditProfileUseCase) Execute(ctx context.Context, r requests.EditProfileRequest, userID uint) error {
	u, err := uc.UserService.Get(ctx, userID)
	if err != nil {
		return err
	}

	if !u.Active {
		return pkg.InactiveUserError
	}

	if !u.Admin {
		return pkg.UnauthorizedError
	}

	fields := common.StructToMap(r, "user_id")
	err = uc.UserService.Update(ctx, u.ID, &fields)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return nil
}

func NewEditProfileUseCase(userService user.Servicer, logger *log.Logger) *EditProfileUseCase {
	return &EditProfileUseCase{
		UserService: userService,
		logger:      logger,
	}
}
