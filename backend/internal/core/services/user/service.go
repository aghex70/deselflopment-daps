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
	UserId int `json:"user_id"`
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
		UserId:      nu.Id,
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
		UserId: u.Id,
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

	return ss, u.Id, nil
}

func (s UserService) RefreshToken(ctx context.Context, r *http.Request) (string, error) {
	userId, err := server.RetrieveJWTClaims(r, nil)
	u, err := s.userRepository.Get(ctx, int(userId))
	if err != nil {
		return "", errors.New("invalid token")
	}

	newClaims := MyCustomClaims{
		UserId: u.Id,
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

func (s UserService) CheckAdmin(ctx context.Context, r *http.Request) (int, error) {
	userId, err := server.RetrieveJWTClaims(r, nil)
	u, err := s.userRepository.Get(ctx, int(userId))
	if err != nil {
		return 0, errors.New("invalid token")
	}

	if !u.IsAdmin {
		return 0, errors.New("unauthorized")
	}

	return int(userId), nil
}

func (s UserService) Delete(ctx context.Context, r *http.Request, req ports.DeleteUserRequest) error {
	adminId, err := s.CheckAdmin(ctx, r)
	if err != nil {
		return err
	}

	err = s.userRepository.Delete(ctx, adminId, int(req.UserId))
	if err != nil {
		return err
	}

	return nil
}

func (s UserService) Get(ctx context.Context, r *http.Request, req ports.GetUserRequest) (domain.User, error) {
	_, err := s.CheckAdmin(ctx, r)
	if err != nil {
		return domain.User{}, err
	}

	u, err := s.userRepository.Get(ctx, int(req.UserId))
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}

func (s UserService) ProvisionDemoUser(ctx context.Context, r *http.Request, req ports.ProvisionDemoUserRequest) error {
	_, err := s.CheckAdmin(ctx, r)
	if err != nil {
		return err
	}

	cipheredPassword := s.EncryptPassword(ctx, req.Password)
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
		UserId:      nu.Id,
		AutoSuggest: false,
		Language:    "en",
	}

	err = s.userConfigurationRepository.Create(ctx, nuc)
	if err != nil {
		return err
	}

	demoCategory := domain.Category{
		OwnerId:     nu.Id,
		Description: "Home tasks",
		Custom:      true,
		Name:        "Home",
		Users:       []domain.User{u},
	}

	c, err := s.categoryRepository.Create(ctx, demoCategory, nu.Id)

	anotherDemoCategory := domain.Category{
		OwnerId:     nu.Id,
		Description: "Work stuff",
		Custom:      true,
		Name:        "Work",
		Users:       []domain.User{u},
	}

	ac, err := s.categoryRepository.Create(ctx, anotherDemoCategory, nu.Id)

	yetAnotherDemoCategory := domain.Category{
		OwnerId:     nu.Id,
		Description: "Purchase list",
		Custom:      true,
		Name:        "Purchases",
		Users:       []domain.User{u},
	}

	yac, err := s.categoryRepository.Create(ctx, yetAnotherDemoCategory, nu.Id)

	todos := pkg.GenerateDemoTodos(c.Id, ac.Id, yac.Id, req.Language)

	for _, t := range todos {
		err = s.todoRepository.Create(ctx, t)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s UserService) List(ctx context.Context, r *http.Request) ([]domain.User, error) {
	_, err := s.CheckAdmin(ctx, r)
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
