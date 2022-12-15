package user

import (
	"context"
	"errors"
	"github.com/aghex70/daps/internal/core/domain"
	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/repositories/gorm/category"
	"github.com/aghex70/daps/internal/repositories/gorm/todo"
	"github.com/aghex70/daps/internal/repositories/gorm/user"
	"github.com/aghex70/daps/internal/repositories/gorm/userconfig"
	"github.com/aghex70/daps/pkg"
	"github.com/aghex70/daps/server"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"time"
)

type UserService struct {
	logger                      *log.Logger
	userRepository              *user.UserGormRepository
	categoryRepository          *category.CategoryGormRepository
	userConfigurationRepository *userconfig.UserConfigGormRepository
	todoRepository              *todo.TodoGormRepository
}

type MyCustomClaims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

func (s UserService) Register(ctx context.Context, r ports.CreateUserRequest) error {
	preexistent := s.CheckExistentUser(ctx, r.Email)
	if preexistent {
		return errors.New("user already registered")
	}
	cipheredPassword := s.EncryptPassword(ctx, r.Password)

	categories, err := s.categoryRepository.GetByIds(ctx, pkg.BaseCategoriesIds)
	u := domain.User{
		Name:       r.Name,
		Email:      r.Email,
		Password:   cipheredPassword,
		Categories: categories,
	}

	nu, err := s.userRepository.Create(ctx, u)
	if err != nil {
		return err
	}

	nuc := domain.UserConfig{
		UserId:      nu.ID,
		AutoSuggest: false,
		Language:    "en",
	}

	err = s.userConfigurationRepository.Create(ctx, nuc)
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

	mySigningKey := pkg.HmacSampleSecret
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

	mySigningKey := pkg.HmacSampleSecret
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, newClaims)
	ss, err := newToken.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return ss, nil
}

func (s UserService) CheckAdmin(ctx context.Context, r *http.Request) error {
	userId, err := server.RetrieveJWTClaims(r, nil)
	u, err := s.userRepository.Get(ctx, int(userId))
	if err != nil {
		return errors.New("invalid token")
	}

	if !u.IsAdmin {
		return errors.New("unauthorized")
	}

	return nil
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

func (s UserService) ProvisionDemoUser(ctx context.Context, r *http.Request, req ports.ProvisionDemoUserRequest) error {
	err := s.CheckAdmin(ctx, r)
	if err != nil {
		return err
	}

	cipheredPassword := s.EncryptPassword(ctx, pkg.DemoUserPassword)
	u := domain.User{
		Name:     pkg.DemoUserName,
		Email:    req.Email,
		Password: cipheredPassword,
	}

	nu, err := s.userRepository.Create(ctx, u)
	if err != nil {
		return err
	}

	nuc := domain.UserConfig{
		UserId:      nu.ID,
		AutoSuggest: false,
		Language:    "en",
	}

	err = s.userConfigurationRepository.Create(ctx, nuc)
	if err != nil {
		return err
	}

	demoCategory := domain.Category{
		OwnerID:     nu.ID,
		Description: "Home tasks",
		Custom:      true,
		Name:        "Home",
		Users:       []domain.User{u},
	}

	c, err := s.categoryRepository.Create(ctx, demoCategory, nu.ID)

	anotherDemoCategory := domain.Category{
		OwnerID:     nu.ID,
		Description: "Work stuff",
		Custom:      true,
		Name:        "Work",
		Users:       []domain.User{u},
	}

	ac, err := s.categoryRepository.Create(ctx, anotherDemoCategory, nu.ID)

	todos := pkg.GenerateDemoTodos(c.ID, ac.ID, req.Language)

	for _, t := range todos {
		err = s.todoRepository.Create(ctx, t)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s UserService) List(ctx context.Context, r *http.Request) ([]domain.User, error) {
	err := s.CheckAdmin(ctx, r)
	if err != nil {
		return []domain.User{}, err
	}

	users, err := s.userRepository.List(ctx)
	if err != nil {
		return []domain.User{}, err
	}

	return users, nil
}

func NewUserService(ur *user.UserGormRepository, cr *category.CategoryGormRepository, ucr *userconfig.UserConfigGormRepository, tr *todo.TodoGormRepository, logger *log.Logger) UserService {
	return UserService{
		logger:                      logger,
		userRepository:              ur,
		categoryRepository:          cr,
		userConfigurationRepository: ucr,
		todoRepository:              tr,
	}
}
