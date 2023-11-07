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
	repository "github.com/aghex70/daps/internal/repositories/gorm"
	"github.com/aghex70/daps/pkg"
	"github.com/aghex70/daps/server"
	"github.com/golang-jwt/jwt/v4"
)

type Service struct {
	logger     *log.Logger
	repository *repository.GormRepository
}

type MyCustomClaims struct {
	UserId int  `json:"user_id"`
	Admin  bool `json:"admin"`
	jwt.RegisteredClaims
}

func (s Service) Register(ctx context.Context, r ports.CreateUserRequest) error {
	preexistent := s.CheckExistentUser(ctx, r.Email)
	if preexistent {
		return errors.New("user already registered")
	}
	cipheredPassword := s.EncryptPassword(ctx, r.Password)

	categories, err := s.repository.GetCategories(ctx, pkg.BaseCategoriesIds)
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

	nu, err := s.repository.CreateUser(ctx, u)
	if err != nil {
		return err
	}

	nuc := domain.UserConfig{
		UserId:      nu.Id,
		AutoSuggest: false,
		Language:    "en",
	}

	//err = s.userConfigurationRepository.Create(ctx, nuc)
	//if err != nil {
	//	return err
	//}

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

		err = s.repository.DeleteUser(ctx, 0, nu.Id)
		if err != nil {
			return err
		}

		_, errz := s.repository.CreateEmail(ctx, e)
		if errz != nil {
			fmt.Printf("Error saving email: %+v", errz)
			return errz
		}
		return err
	}

	e.Sent = true
	_, err = s.repository.CreateEmail(ctx, e)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) Login(ctx context.Context, r ports.LoginUserRequest) (string, int, error) {
	u, err := s.repository.GetUser(ctx, r.Email)
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
		Admin:  u.Admin,
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
	u, err := s.repository.GetUser(ctx, int(userId))
	if err != nil {
		return "", errors.New("invalid token")
	}

	newClaims := MyCustomClaims{
		UserId: u.Id,
		Admin:  u.Admin,
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
	u, err := s.repository.GetUser(ctx, int(userId))
	if err != nil {
		return 0, errors.New("invalid token")
	}

	if !u.Admin {
		return 0, errors.New("unauthorized")
	}

	return int(userId), nil
}

func (s Service) Delete(ctx context.Context, r *http.Request, req ports.DeleteUserRequest) error {
	adminId, err := s.CheckAdmin(ctx, r)
	if err != nil {
		return err
	}

	err = s.repository.DeleteUser(ctx, adminId, int(req.UserId))
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

	u, err := s.repository.GetUser(ctx, int(req.UserId))
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

	nu, err := s.repository.CreateUser(ctx, u)
	if err != nil {
		return err
	}

	nuc := domain.UserConfig{
		UserId:      nu.Id,
		AutoSuggest: false,
		Language:    "en",
	}

	//err = s.userConfigurationRepository.Create(ctx, nuc)
	//err = s.userConfigurationRepository.Create(ctx, nuc)
	//if err != nil {
	//	return err
	//}

	demoCategory := domain.Category{
		OwnerId:     nu.Id,
		Description: "Home tasks",
		Custom:      true,
		Name:        "Home",
		Users:       []domain.User{u},
	}

	c, err := s.repository.CreateCategory(ctx, demoCategory, nu.Id)
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

	ac, err := s.repository.CreateCategory(ctx, anotherDemoCategory, nu.Id)
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

	yac, err := s.repository.CreateCategory(ctx, yetAnotherDemoCategory, nu.Id)
	if err != nil {
		return err
	}

	todos := pkg.GenerateDemoTodos(c.Id, ac.Id, yac.Id, req.Language)

	for _, t := range todos {
		err, _ = s.repository.CreateTodo(ctx, t)
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

	users, err := s.repository.GetUsers(ctx)
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

		err, _ = s.repository.CreateTodo(ctx, domain.Todo{
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
	err := s.repository.ActivateUser(ctx, r.ActivationCode)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) SendResetLink(ctx context.Context, r ports.ResetLinkRequest) error {
	u, err := s.repository.CreateResetLink(ctx, r.Email)
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
		_, errz := s.repository.CreateEmail(ctx, e)
		if errz != nil {
			fmt.Printf("Error saving email: %+v", errz)
			return errz
		}
		return err
	}

	e.Sent = true
	_, err = s.repository.CreateEmail(ctx, e)
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
	err := s.repository.ResetPassword(ctx, encryptedPassword, r.ResetPasswordCode)
	if err != nil {
		return err
	}

	return nil
}

func NewUserService(gr *repository.GormRepository, logger *log.Logger) Service {
	return Service{
		logger:     logger,
		repository: gr,
	}
}
