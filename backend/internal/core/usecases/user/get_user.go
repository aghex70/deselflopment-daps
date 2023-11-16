package user

import (
	"context"
	"github.com/aghex70/daps/internal/core/services/user"
	"github.com/aghex70/daps/internal/ports/domain"
	"log"
)

type GetUserUseCase struct {
	UserService user.Service
	logger      *log.Logger
}

func (uc *GetUserUseCase) Execute(ctx context.Context, id uint) (domain.User, error) {
	//_, err := s.CheckAdmin(ctx, r)
	//if err != nil {
	//	return domain2.User{}, err
	//}
	//
	u, err := uc.UserService.Get(ctx, id)
	if err != nil {
		return domain.User{}, err
	}
	return u, nil
}

func NewGetUserUseCase(userService user.Service, logger *log.Logger) *GetUserUseCase {
	return &GetUserUseCase{
		UserService: userService,
		logger:      logger,
	}
}
