package user

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
	"github.com/aghex70/daps/internal/ports/repositories/user"
	"log"
	"net/http"
)

type Service struct {
	logger         *log.Logger
	userRepository user.Repository
}

func (s Service) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	return s.userRepository.GetByEmail(ctx, email)
}

func (s Service) Create(ctx context.Context, u domain.User) (domain.User, error) {
	nu, err := s.userRepository.Create(ctx, u)
	if err != nil {
		return domain.User{}, err
	}
	return nu, nil
}

func (s Service) Delete(ctx context.Context, id uint) error {
	//_, err := uc.UserService.CheckAdmin(ctx, r)
	//if err != nil {
	//	return err
	//}
	//
	err := s.userRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) Activate(ctx context.Context, activationCode string) error {
	return s.userRepository.Activate(ctx, activationCode)
}

func (s Service) List(ctx context.Context, r *http.Request) ([]domain.User, error) {
	//_, err := s.CheckAdmin(ctx, r)
	//if err != nil {
	//	return []domain.User{}, err
	//}

	users, err := s.userRepository.List(ctx, nil)
	if err != nil {
		return []domain.User{}, err
	}

	return users, nil
}

func (s Service) Get(ctx context.Context, id uint) (domain.User, error) {
	//_, err := s.CheckAdmin(ctx, r)
	//if err != nil {
	//	return domain.User{}, err
	//}
	//
	u, err := s.userRepository.Get(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}

func (s Service) ResetPassword(ctx context.Context, password, resetPasswordCode string) error {
	err := s.userRepository.ResetPassword(ctx, password, resetPasswordCode)
	if err != nil {
		return err
	}
}

func (s Service) Update(ctx context.Context, id uint, fields map[string]interface{}) error {
	//err := s.userRepository.Update(ctx, activationCode)
	//if err != nil {
	//	return err
	//}

	return nil
}

func NewUserService(ur user.Repository, logger *log.Logger) Service {
	return Service{
		logger:         logger,
		userRepository: ur,
	}
}
