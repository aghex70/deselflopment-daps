package user

import (
	"context"
	"github.com/aghex70/daps/internal/core/services/user"
	"github.com/aghex70/daps/internal/pkg"
	requests "github.com/aghex70/daps/internal/ports/requests/user"
	utils "github.com/aghex70/daps/utils/user"
	"log"
)

type ResetPasswordUseCase struct {
	UserService user.Service
	logger      *log.Logger
}

func (uc *ResetPasswordUseCase) Execute(ctx context.Context, r requests.ResetPasswordRequest) error {
	match := utils.PasswordMatchesRepeatPassword(ctx, r.Password, r.RepeatPassword)
	if !match {
		return pkg.PasswordsDoNotMatchError
	}

	encryptedPassword := utils.EncryptPassword(ctx, r.Password)
	err := uc.UserService.ResetPassword(ctx, encryptedPassword, r.ResetPasswordCode)
	if err != nil {
		return err
	}
	return nil
}

func NewResetPasswordUseCase(userService user.Service, logger *log.Logger) *ResetPasswordUseCase {
	return &ResetPasswordUseCase{
		UserService: userService,
		logger:      logger,
	}
}
