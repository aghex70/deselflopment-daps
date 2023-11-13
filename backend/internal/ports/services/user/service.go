package user

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
	"net/http"
)

type Servicer interface {
	GetByEmail(ctx context.Context, email string) (domain.User, error)
	Create(ctx context.Context, u domain.User) (domain.User, error)
	Delete(ctx context.Context, id uint) error
	Activate(ctx context.Context, activationCode string) error
	List(ctx context.Context, r *http.Request) ([]domain.User, error)
	Get(ctx context.Context, id uint) (domain.User, error)
	ResetPassword(ctx context.Context, password, resetPasswordCode string) error
	Update(ctx context.Context, id uint, fields map[string]interface{}) error
}
