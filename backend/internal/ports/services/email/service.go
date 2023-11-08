package email

import (
	"context"
	requests "github.com/aghex70/daps/internal/ports/requests/email"
	"net/http"
)

type Servicer interface {
	Send(ctx context.Context, r *http.Request, req requests.SendEmailRequest) error
}
