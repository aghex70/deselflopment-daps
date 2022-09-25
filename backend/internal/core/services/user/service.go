package user

import (
	"context"
	"errors"
	"github.com/aghex70/daps/internal/core/domain"
	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/repositories/gorm/user"
	"github.com/aghex70/daps/server"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

type UserService struct {
	logger         *log.Logger
	userRepository *user.UserGormRepository
}

type MyCustomClaims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

var hmacSampleSecret = []byte("random")

func (s UserService) Login(ctx context.Context, r ports.LoginUserRequest) (string, error) {
	user, err := s.userRepository.GetByEmail(ctx, r.Email)
	if err != nil {
		return "", err
	}

	claims := MyCustomClaims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   r.Email,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(96 * time.Hour)),
		},
	}

	mySigningKey := hmacSampleSecret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return ss, nil
}

func (s UserService) Logout(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (s UserService) Register(ctx context.Context, r ports.CreateUserRequest) error {
	preexistent := s.CheckExistentUser(ctx, r.Email)
	if preexistent {
		return errors.New("user already registered")
	}

	nu := domain.User{
		Email:    r.Email,
		Password: r.Password,
	}

	err := s.userRepository.Create(ctx, nu)
	if err != nil {
		return err
	}
	return nil
}

func (s UserService) RefreshToken(ctx context.Context, r *http.Request) (string, error) {
	userId, err := server.RetrieveJWTClaims(r, nil)
	user, err := s.userRepository.Get(ctx, uint(userId))
	if err != nil {
		return "", errors.New("invalid token")
	}

	newClaims := MyCustomClaims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.Email,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(96 * time.Hour)),
		},
	}

	mySigningKey := hmacSampleSecret
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	ss, err := newToken.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return ss, nil
}

func (s UserService) Remove(ctx context.Context, r *http.Request) error {
	userId, err := server.RetrieveJWTClaims(r, nil)
	err = s.userRepository.Delete(ctx, uint(userId))
	if err != nil {
		return err
	}
	return nil
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
