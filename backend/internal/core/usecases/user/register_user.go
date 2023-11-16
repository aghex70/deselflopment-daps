package user

import (
	"context"
	"errors"
	"github.com/aghex70/daps/internal/core/services/category"
	"github.com/aghex70/daps/internal/core/services/email"
	"github.com/aghex70/daps/internal/core/services/user"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	requests "github.com/aghex70/daps/internal/ports/requests/user"
	common "github.com/aghex70/daps/utils"
	utils "github.com/aghex70/daps/utils/user"
	"gorm.io/gorm"
	"log"
)

type RegisterUserUseCase struct {
	UserService     user.Service
	CategoryService category.Service
	EmailService    email.Service
	logger          *log.Logger
}

func (uc *RegisterUserUseCase) Execute(ctx context.Context, r requests.CreateUserRequest) error {
	_, err := uc.UserService.GetByEmail(ctx, r.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return pkg.UserAlreadyRegisteredError
		}
		return err
	}

	categories, err := uc.CategoryService.List(ctx, &pkg.BaseCategoriesIds, nil)
	if err != nil {
		return err
	}
	log.Printf("Categories: %+v", categories)

	encryptedPassword := utils.EncryptPassword(ctx, r.Password)

	u := domain.User{
		Name:              r.Name,
		Email:             r.Email,
		Password:          encryptedPassword,
		ActivationCode:    common.GenerateUUID(),
		ResetPasswordCode: common.GenerateUUID(),
		Categories:        &categories,
	}

	log.Printf("User: %+v", u)
	nu, err := uc.UserService.Create(ctx, u)
	if err != nil {
		return err
	}

	e := domain.Email{
		Subject:   "📣 DAPS - Activate your account 📣",
		Body:      "In order to complete your registration, please click on the following link: " + pkg.ActivationCodeLink + nu.ActivationCode,
		From:      pkg.FromEmail,
		Source:    pkg.ProjectName,
		To:        r.Name,
		Recipient: r.Email,
		UserID:    nu.ID,
	}

	s, err := uc.EmailService.Send(ctx, e)
	if !s && err != nil {
		uerr := uc.UserService.Delete(ctx, nu.ID)
		if uerr != nil {
			return uerr
		}
		return err
	}
	log.Println("Registering!!!!!")
	log.Println("Registering!!!!!")
	log.Println("Registering!!!!!")
	log.Println("Registering!!!!!")
	log.Println("Registering!!!!!")
	log.Println("Registering!!!!!")
	return nil
}

func NewRegisterUserUseCase(us user.Service, cs category.Service, es email.Service, logger *log.Logger) RegisterUserUseCase {
	return RegisterUserUseCase{
		UserService:     us,
		CategoryService: cs,
		EmailService:    es,
		logger:          logger,
	}
}
