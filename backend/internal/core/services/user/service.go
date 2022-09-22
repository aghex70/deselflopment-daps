package user

import (
	"context"
	"errors"
	"fmt"
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
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
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

func (s UserService) Logout(ctx context.Context, r ports.LogoutUserRequest) error {
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

func (s UserService) RefreshToken(ctx context.Context, r ports.RefreshTokenRequest) (string, error) {
	token, err := jwt.Parse(r.AccessToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return hmacSampleSecret, nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", err
	}
	//} else if errors.Is(err, jwt.ErrTokenMalformed) {
	//	fmt.Println("That's not even a token")
	//	return "", err
	//} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
	//	// Token is either expired or not active yet
	//	fmt.Println("Timing is everything")
	//	return "", err
	//} else {
	//	fmt.Println("Couldn't handle this token:", err)
	//	return "", err
	//}
	userId := claims["user_id"].(float64)
	user, err := s.userRepository.Get(ctx, uint(userId))
	if err != nil {
		return "", errors.New("invalid token")
	}

	newClaims := MyCustomClaims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.Email,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
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
