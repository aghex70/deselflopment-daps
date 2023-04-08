package user

import (
	"bufio"
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/aghex70/daps/internal/core/domain"
	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/repositories/gorm/category"
	"github.com/aghex70/daps/internal/repositories/gorm/email"
	"github.com/aghex70/daps/internal/repositories/gorm/todo"
	"github.com/aghex70/daps/internal/repositories/gorm/user"
	"github.com/aghex70/daps/internal/repositories/gorm/userconfig"
	"github.com/aghex70/daps/pkg"
	"github.com/aghex70/daps/server"
	"github.com/golang-jwt/jwt/v4"
)

type Service struct {
	logger                      *log.Logger
	userRepository              *user.GormRepository
	categoryRepository          *category.GormRepository
	userConfigurationRepository *userconfig.GormRepository
	todoRepository              *todo.GormRepository
	emailRepository             *email.GormRepository
}

type MyCustomClaims struct {
	UserId int `json:"user_id"`
	jwt.RegisteredClaims
}

func (s Service) Register(ctx context.Context, r ports.CreateUserRequest) error {
	preexistent := s.CheckExistentUser(ctx, r.Email)
	if preexistent {
		return errors.New("user already registered")
	}
	cipheredPassword := s.EncryptPassword(ctx, r.Password)

	categories, err := s.categoryRepository.GetByIds(ctx, pkg.BaseCategoriesIds)
	if err != nil {
		return err
	}
	u := domain.User{
		Name:              r.Name,
		Email:             r.Email,
		Password:          cipheredPassword,
		Categories:        categories,
		Active:            false,
		ActivationCode:    pkg.GenerateUUID(),
		ResetPasswordCode: pkg.GenerateUUID(),
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

	e := domain.Email{
		From:      pkg.FromEmail,
		To:        r.Email,
		Recipient: r.Name,
		Subject:   "📣 DAPS - Activate your account 📣",
		Body:      "In order to complete your registration, please click on the following link: " + pkg.ActivationCodeLink + nu.ActivationCode,
		User:      nu.Id,
	}

	err = pkg.SendEmail(e)
	if err != nil {
		fmt.Printf("Error sending email: %+v", err)
		e.Error = err.Error()
		e.Sent = false

		err = s.userRepository.Delete(ctx, 0, nu.Id)
		if err != nil {
			return err
		}

		_, errz := s.emailRepository.Create(ctx, e)
		if errz != nil {
			fmt.Printf("Error saving email: %+v", errz)
			return errz
		}
		return err
	}

	e.Sent = true
	_, err = s.emailRepository.Create(ctx, e)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) Login(ctx context.Context, r ports.LoginUserRequest) (string, int, error) {
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

func (s Service) RefreshToken(ctx context.Context, r *http.Request) (string, error) {
	userId, err := server.RetrieveJWTClaims(r, nil)
	if err != nil {
		return "", errors.New("invalid token")
	}
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

func (s Service) CheckAdmin(ctx context.Context, r *http.Request) (int, error) {
	userId, err := server.RetrieveJWTClaims(r, nil)
	if err != nil {
		return 0, errors.New("invalid token")
	}
	u, err := s.userRepository.Get(ctx, int(userId))
	if err != nil {
		return 0, errors.New("invalid token")
	}

	if !u.IsAdmin {
		return 0, errors.New("unauthorized")
	}

	return int(userId), nil
}

func (s Service) Delete(ctx context.Context, r *http.Request, req ports.DeleteUserRequest) error {
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

func (s Service) Get(ctx context.Context, r *http.Request, req ports.GetUserRequest) (domain.User, error) {
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

func (s Service) ProvisionDemoUser(ctx context.Context, r *http.Request, req ports.ProvisionDemoUserRequest) error {
	_, err := s.CheckAdmin(ctx, r)
	if err != nil {
		return err
	}

	cipheredPassword := s.EncryptPassword(ctx, req.Password)
	u := domain.User{
		Name:              pkg.DemoUserName,
		Email:             req.Email,
		Password:          cipheredPassword,
		Active:            true,
		ResetPasswordCode: pkg.GenerateUUID(),
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
	if err != nil {
		return err
	}

	anotherDemoCategory := domain.Category{
		OwnerId:     nu.Id,
		Description: "Work stuff",
		Custom:      true,
		Name:        "Work",
		Users:       []domain.User{u},
	}

	ac, err := s.categoryRepository.Create(ctx, anotherDemoCategory, nu.Id)
	if err != nil {
		return err
	}

	yetAnotherDemoCategory := domain.Category{
		OwnerId:     nu.Id,
		Description: "Purchase list",
		Custom:      true,
		Name:        "Purchases",
		Users:       []domain.User{u},
	}

	yac, err := s.categoryRepository.Create(ctx, yetAnotherDemoCategory, nu.Id)
	if err != nil {
		return err
	}

	todos := pkg.GenerateDemoTodos(c.Id, ac.Id, yac.Id, req.Language)

	for _, t := range todos {
		err = s.todoRepository.Create(ctx, t)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s Service) List(ctx context.Context, r *http.Request) ([]domain.User, error) {
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

func (s Service) ImportCSV(ctx context.Context, r *http.Request, f multipart.File) error {
	_, err := s.CheckAdmin(ctx, r)
	if err != nil {
		return err
	}

	// Create a buffer to read the file line by line
	buf := bufio.NewReader(f)

	// Parse the CSV file
	rr := csv.NewReader(buf)

	// Read and discard the first line
	_, err = rr.Read()
	if err != nil {
		return err
	}

	// Iterate over the lines of the CSV file
	for {
		// Read the next line
		record, err := rr.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		name := record[0]
		link := record[1]
		categoryId, _ := strconv.Atoi(record[2])

		err = s.todoRepository.Create(ctx, domain.Todo{
			Name:     name,
			Link:     link,
			Category: categoryId,
			Priority: domain.Priority(3),
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (s Service) Activate(ctx context.Context, r ports.ActivateUserRequest) error {
	err := s.userRepository.ActivateUser(ctx, r.ActivationCode)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) SendResetLink(ctx context.Context, r ports.ResetLinkRequest) error {
	u, err := s.userRepository.CreateResetLink(ctx, r.Email)
	if err != nil {
		return err
	}

	e := domain.Email{
		From:      pkg.FromEmail,
		To:        u.Email,
		Recipient: u.Name,
		Subject:   "📣 DAPS - Password reset request 📣",
		Body:      "In order to reset your password, please follow this link: " + pkg.ResetPasswordLink + u.ResetPasswordCode,
		User:      u.Id,
	}

	err = pkg.SendEmail(e)
	if err != nil {
		fmt.Printf("Error sending email: %+v", err)
		e.Error = err.Error()
		e.Sent = false
		_, errz := s.emailRepository.Create(ctx, e)
		if errz != nil {
			fmt.Printf("Error saving email: %+v", errz)
			return errz
		}
		return err
	}

	e.Sent = true
	_, err = s.emailRepository.Create(ctx, e)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) ResetPassword(ctx context.Context, r ports.ResetPasswordRequest) error {
	match := s.PasswordMatchesRepeatPassword(ctx, r.Password, r.RepeatPassword)
	if !match {
		return errors.New("passwords do not match")
	}

	encryptedPassword := s.EncryptPassword(ctx, r.Password)
	err := s.userRepository.ResetPassword(ctx, encryptedPassword, r.ResetPasswordCode)
	if err != nil {
		return err
	}

	return nil
}

func NewUserService(ur *user.GormRepository, cr *category.GormRepository, ucr *userconfig.GormRepository, tr *todo.GormRepository, er *email.GormRepository, logger *log.Logger) Service {
	return Service{
		logger:                      logger,
		userRepository:              ur,
		categoryRepository:          cr,
		userConfigurationRepository: ucr,
		todoRepository:              tr,
		emailRepository:             er,
	}
}
