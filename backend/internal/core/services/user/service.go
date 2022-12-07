package user

import (
	"context"
	"errors"
	"github.com/aghex70/daps/internal/core/domain"
	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/repositories/gorm/category"
	"github.com/aghex70/daps/internal/repositories/gorm/user"
	"github.com/aghex70/daps/server"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"time"
)

type UserService struct {
	logger             *log.Logger
	userRepository     *user.UserGormRepository
	categoryRepository *category.CategoryGormRepository
}

type MyCustomClaims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

var baseCategoriesIds = []int{1, 2, 3, 4, 5}
var hmacSampleSecret = []byte("random")

func (s UserService) Register(ctx context.Context, r ports.CreateUserRequest) error {
	preexistent := s.CheckExistentUser(ctx, r.Email)
	if preexistent {
		return errors.New("user already registered")
	}
	cipheredPassword := s.EncryptPassword(ctx, r.Password)

	categories, err := s.categoryRepository.GetByIds(ctx, baseCategoriesIds)
	u := domain.User{
		Name:       r.Name,
		Email:      r.Email,
		Password:   cipheredPassword,
		Categories: categories,
	}

	_, err = s.userRepository.Create(ctx, u)
	if err != nil {
		return err
	}

	return nil
}

func (s UserService) Login(ctx context.Context, r ports.LoginUserRequest) (string, int, error) {
	u, err := s.userRepository.GetByEmail(ctx, r.Email)
	if err != nil {
		return "", 0, err
	}

	decryptedPassword, err := s.DecryptPassword(ctx, u.Password)
	if err != nil {
		return "", 0, err
	}

	match := s.PasswordsMatch(ctx, decryptedPassword, r.Password)
	if !match {
		return "", 0, errors.New("invalid credentials")
	}

	claims := MyCustomClaims{
		UserID: u.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   r.Email,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(96 * time.Hour)),
		},
	}

	mySigningKey := hmacSampleSecret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", 0, err
	}

	return ss, u.ID, nil
}

func (s UserService) RefreshToken(ctx context.Context, r *http.Request) (string, error) {
	userId, err := server.RetrieveJWTClaims(r, nil)
	u, err := s.userRepository.Get(ctx, int(userId))
	if err != nil {
		return "", errors.New("invalid token")
	}

	newClaims := MyCustomClaims{
		UserID: u.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   u.Email,
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
	if err != nil {
		return err
	}

	err = s.userRepository.Delete(ctx, int(userId))
	if err != nil {
		return err
	}

	return nil
}

func NewUserService(ur *user.UserGormRepository, cr *category.CategoryGormRepository, logger *log.Logger) UserService {
	return UserService{
		logger:             logger,
		userRepository:     ur,
		categoryRepository: cr,
	}
}
