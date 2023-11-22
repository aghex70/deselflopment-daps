package user

import (
	"context"
	"github.com/aghex70/daps/internal/core/services/user"
	requests "github.com/aghex70/daps/internal/ports/requests/user"
	utils "github.com/aghex70/daps/utils/user"
	"log"
)

type ResetPasswordUseCase struct {
	UserService user.Service
	logger      *log.Logger
}

func (uc *ResetPasswordUseCase) Execute(ctx context.Context, r requests.ResetPasswordRequest, userID uint) error {
	encryptedPassword := utils.EncryptPassword(ctx, r.Password)
	if err := uc.UserService.ResetPassword(ctx, userID, encryptedPassword, r.ResetPasswordCode); err != nil {
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
