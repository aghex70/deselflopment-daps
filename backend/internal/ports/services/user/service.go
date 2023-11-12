package user

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
	requests "github.com/aghex70/daps/internal/ports/requests/user"
	"net/http"
)

type Servicer interface {
	GetByEmail(ctx context.Context, email string) (domain.User, error)
	Create(ctx context.Context, u domain.User) (domain.User, error)
	Delete(ctx context.Context, id uint) error
	Activate(ctx context.Context, activationCode string) error
	Update(ctx context.Context, id uint, fields map[string]interface{}) error
	Get(ctx context.Context, r *http.Request, req requests.GetUserRequest) (domain.User, error)
	List(ctx context.Context, r *http.Request) ([]domain.User, error)
}
