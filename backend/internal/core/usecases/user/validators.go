package user

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

func (s Service) UserExists(ctx context.Context, email string) bool {
	_, err := s.repository.GetByEmail(ctx, email)
	return !errors.Is(err, gorm.ErrRecordNotFound)
}
