package user

import (
	"context"
	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/repositories/gorm/user"
	"log"
)

type UserService struct {
	logger         *log.Logger
	userRepository *user.UserGormRepository
}

func (s UserService) Login(ctx context.Context, request ports.LoginUserRequest) error {
	//TODO implement me
	panic("implement me")
}

func (s UserService) Logout(ctx context.Context, request ports.LogoutUserRequest) error {
	//TODO implement me
	panic("implement me")
}

func (s UserService) Register(context.Context, ports.CreateUserRequest) error {
	panic("foo")
}

func (s UserService) Remove(context.Context, ports.DeleteUserRequest) error {
	panic("foo")
}

func NewUserService(ur *user.UserGormRepository, logger *log.Logger) UserService {
	return UserService{
		logger:         logger,
		userRepository: ur,
	}
}
