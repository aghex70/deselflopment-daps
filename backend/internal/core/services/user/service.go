package user

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
	"github.com/aghex70/daps/internal/ports/repositories/user"
	"log"
)

type Service struct {
	logger         *log.Logger
	userRepository user.Repository
}

func (s Service) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := s.userRepository.GetByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return u, nil
}

func (s Service) Create(ctx context.Context, u domain.User) (domain.User, error) {
	nu, err := s.userRepository.Create(ctx, u)
	if err != nil {
		return domain.User{}, err
	}
	return nu, nil
}

func (s Service) Delete(ctx context.Context, id uint) error {
	if err := s.userRepository.Delete(ctx, id); err != nil {
		return err
	}
	return nil
}

func (s Service) Activate(ctx context.Context, id uint, activationCode string) error {
	return s.userRepository.Activate(ctx, id, activationCode)
}

func (s Service) List(ctx context.Context, fields *map[string]interface{}) ([]domain.User, error) {
	users, err := s.userRepository.List(ctx, fields)
	if err != nil {
		return []domain.User{}, err
	}

	return users, nil
}

func (s Service) Get(ctx context.Context, id uint) (domain.User, error) {
	u, err := s.userRepository.Get(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	return u, nil
}

func (s Service) ResetPassword(ctx context.Context, userID uint, password, resetPasswordCode string) error {
	if err := s.userRepository.ResetPassword(ctx, userID, password, resetPasswordCode); err != nil {
		return err
	}
	return nil
}

func (s Service) Update(ctx context.Context, id uint, fields *map[string]interface{}) error {
	if err := s.userRepository.Update(ctx, id, fields); err != nil {
		return err
	}
	return nil
}

func NewUserService(ur user.Repository, logger *log.Logger) Service {
	return Service{
		logger:         logger,
		userRepository: ur,
	}
}
