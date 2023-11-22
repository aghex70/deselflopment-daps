package user

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	requests "github.com/aghex70/daps/internal/ports/requests/user"
	"github.com/aghex70/daps/internal/ports/services/user"
	utils "github.com/aghex70/daps/utils/user"
	"log"
)

type LoginUserUseCase struct {
	UserService user.Servicer
	logger      *log.Logger
}

func (uc *LoginUserUseCase) Execute(ctx context.Context, r requests.LoginUserRequest) (string, uint, error) {
	u, err := uc.UserService.GetByEmail(ctx, r.Email)
	if err != nil {
		return "", 0, err
	}

	decryptedPassword, err := utils.DecryptPassword(ctx, u.Password)
	if err != nil {
		return "", 0, err
	}
	match := utils.PasswordsMatch(ctx, decryptedPassword, r.Password)
	if !match {
		return "", 0, pkg.InvalidCredentialsError
	}

	token, userID, err := utils.GenerateJWT(ctx, u)
	if err != nil {
		return "", 0, err
	}

	return token, userID, nil
}

func NewLoginUserUseCase(userService user.Servicer, logger *log.Logger) *LoginUserUseCase {
	return &LoginUserUseCase{
		UserService: userService,
		logger:      logger,
	}
}
