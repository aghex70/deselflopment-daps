package user

import (
	"context"
	"errors"
	"github.com/aghex70/daps/internal/core/domain"
	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/repositories/gorm/user"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"log"
	"time"
)

type UserService struct {
	logger         *log.Logger
	userRepository *user.UserGormRepository
}

func (s UserService) Login(ctx context.Context, r ports.LoginUserRequest) error {
	//TODO implement me
	panic("implement me")
}

func (s UserService) Logout(ctx context.Context, r ports.LogoutUserRequest) error {
	//TODO implement me
	panic("implement me")
}

func (s UserService) Register(ctx context.Context, r ports.CreateUserRequest) error {
	preexistent := s.CheckExistentUser(ctx, r.Email)
	if preexistent {
		return errors.New("user already registered")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	nu := domain.User{
		//Category:     r.Category,
		Email:        r.Email,
		Password:     r.Password,
		AccessToken:  token,
		RefreshToken: token,
	}

	err := s.userRepository.Create(ctx, nu)
	if err != nil {
		return err
	}
	return nil
}

func (s UserService) Remove(ctx context.Context, r ports.DeleteUserRequest) error {
	panic("foo")
}

func (s UserService) CheckExistentUser(ctx context.Context, email string) bool {
	_, err := s.userRepository.GetByEmail(ctx, email)
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

func NewUserService(ur *user.UserGormRepository, logger *log.Logger) UserService {
	return UserService{
		logger:         logger,
		userRepository: ur,
	}
}
