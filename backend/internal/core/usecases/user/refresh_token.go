package user

import (
	"context"
	"github.com/aghex70/daps/internal/core/services/user"
	requests "github.com/aghex70/daps/internal/ports/requests/user"
	utils "github.com/aghex70/daps/utils/user"
	"log"
)

type RefreshTokenUseCase struct {
	UserService user.Service
	logger      *log.Logger
}

func (uc *RefreshTokenUseCase) Execute(ctx context.Context, r requests.RefreshTokenRequest) (string, uint, error) {
	u, err := uc.UserService.Get(ctx, r.UserID)
	if err != nil {
		return "", 0, err
	}

	token, userID, err := utils.GenerateJWT(ctx, u)
	if err != nil {
		return "", 0, err
	}

	return token, userID, nil
}

func NewRefreshTokenUseCase(userService user.Service, logger *log.Logger) *RefreshTokenUseCase {
	return &RefreshTokenUseCase{
		UserService: userService,
		logger:      logger,
	}
}
