package user

import (
	"context"
	requests "github.com/aghex70/daps/internal/ports/requests/user"
	"github.com/aghex70/daps/internal/ports/services/user"
	utils "github.com/aghex70/daps/utils/user"
	"log"
)

type ResetPasswordUseCase struct {
	UserService user.Servicer
	logger      *log.Logger
}

func (uc *ResetPasswordUseCase) Execute(ctx context.Context, r requests.ResetPasswordRequest, userID uint) error {
	encryptedPassword := utils.EncryptPassword(ctx, r.Password)
	if err := uc.UserService.ResetPassword(ctx, userID, encryptedPassword, r.ResetPasswordCode); err != nil {
		return err
	}
	return nil
}

func NewResetPasswordUseCase(userService user.Servicer, logger *log.Logger) *ResetPasswordUseCase {
	return &ResetPasswordUseCase{
		UserService: userService,
		logger:      logger,
	}
}
