package email

import (
	"context"
	"github.com/aghex70/daps/internal/ports/domain"
)

type Servicer interface {
	Send(ctx context.Context, e domain.Email) (bool, error)
}
