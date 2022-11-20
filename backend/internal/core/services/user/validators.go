package user

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s UserService) CheckExistentUser(ctx context.Context, email, password string) bool {
	_, err := s.userRepository.GetByEmail(ctx, email, password)
	return !errors.Is(err, gorm.ErrRecordNotFound)
}

func (s UserService) PasswordsMatch(ctx context.Context, password, repeatPassword string) bool {
	return password == repeatPassword
}
