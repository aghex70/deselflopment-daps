package user

import (
	"context"
	"github.com/aghex70/daps/internal/pkg"
	"github.com/aghex70/daps/internal/ports/domain"
	requests "github.com/aghex70/daps/internal/ports/requests/user"
	common "github.com/aghex70/daps/utils"
	utils "github.com/aghex70/daps/utils/user"
	"log"
)

func (s Service) Register(ctx context.Context, r requests.CreateUserRequest) error {
	preexistent := s.UserExists(ctx, r.Email)
	if preexistent {
		return pkg.UserAlreadyRegisteredError
	}

	cipheredPassword := utils.EncryptPassword(ctx, r.Password)

	categories, err := s.categoryRepository.List(ctx, &pkg.BaseCategoryIdFilter)
	if err != nil {
		return err
	}

	log.Printf("Categories: %+v", categories)

	u := domain.User{
		Name:              r.Name,
		Email:             r.Email,
		Password:          cipheredPassword,
		ActivationCode:    common.GenerateUUID(),
		ResetPasswordCode: common.GenerateUUID(),
		Categories:        &categories,
	}

	log.Printf("User: %+v", u)
	nu, err := s.userRepository.Create(ctx, u)
	if err != nil {
		return err
	}

	//e := domain.Email{
	//	Subject:   "📣 DAPS - Activate your account 📣",
	//	Body:      "In order to complete your registration, please click on the following link: " + pkg.ActivationCodeLink + nu.ActivationCode,
	//	From:      pkg.FromEmail,
	//	Source:    pkg.ProjectName,
	//	To:        r.Email,
	//	Recipient: r.Name,
	//	UserID:    nu.ID,
	//}

	//err = pkg2.SendEmail(e)
	//if err != nil {
	//	fmt.Printf("Error sending email: %+v", err)
	//	//e.Error = err.Error()
	//	e.Sent = false
	//
	//	err = s.repository.DeleteUser(ctx, nu.ID)
	//	if err != nil {
	//		return err
	//	}
	//
	//	_, errz := s.repository.CreateEmail(ctx, e)
	//	if errz != nil {
	//		fmt.Printf("Error saving email: %+v", errz)
	//		return errz
	//	}
	//	return err
	//}
	//
	//e.Sent = true
	//_, err = s.repository.CreateEmail(ctx, e)
	//if err != nil {
	//	return err
	//}
	log.Println("Registering!!!!!")
	log.Println("Registering!!!!!")
	log.Println("Registering!!!!!")
	log.Println("Registering!!!!!")
	log.Println("Registering!!!!!")
	log.Println("Registering!!!!!")
	return nil
}
