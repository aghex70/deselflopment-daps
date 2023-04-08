package email

import (
	"context"
	"net/http"

	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/handlers"
)

type Handler struct {
	emailService ports.EmailServicer
}

func (h Handler) CreateEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)

		payload := ports.SendEmailRequest{}
		err := handlers.ValidateRequest(r, &payload)
		if err != nil {
			handlers.ThrowError(err, http.StatusBadRequest, w)
			return
		}

		err = h.emailService.Send(context.TODO(), r, payload)
		if err != nil {
			handlers.ThrowError(err, http.StatusBadRequest, w)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func NewEmailHandler(es ports.EmailServicer) Handler {
	return Handler{
		emailService: es,
	}
}
