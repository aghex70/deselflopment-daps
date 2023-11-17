package user

import (
	"context"
	"errors"
	"github.com/aghex70/daps/internal/core/services/user"
	"github.com/aghex70/daps/internal/pkg"
	requests "github.com/aghex70/daps/internal/ports/requests/user"
	"gorm.io/gorm"
	"log"
)

type UpdateUserUseCase struct {
	UserService user.Service
	logger      *log.Logger
}

func (uc *UpdateUserUseCase) Execute(ctx context.Context, r requests.UpdateUserRequest, userID uint) error {
	u, err := uc.UserService.Get(ctx, r.UserID)
	if err != nil {
		return err
	}
	if u.ID != userID {
		return pkg.UnauthorizedError
	}

	fields := map[string]interface{}{"auto_suggest": r.AutoSuggest, "language": r.Language}
	err = uc.UserService.Update(ctx, u.ID, &fields)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return nil
}

func NewUpdateUserUseCase(us user.Service, logger *log.Logger) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		UserService: us,
		logger:      logger,
	}
}
