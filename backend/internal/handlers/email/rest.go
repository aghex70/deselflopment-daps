package email

import (
	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/handlers"
	"log"
	"net/http"
)

type EmailHandler struct {
	emailService ports.EmailServicer
	logger       *log.Logger
}

func (h EmailHandler) CreateEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)

		payload := ports.SendEmailRequest{}
		err := handlers.ValidateRequest(r, &payload)
		if err != nil {
			handlers.ThrowError(err, http.StatusBadRequest, w)
			return
		}

		err = h.emailService.Send(nil, r, payload)
		if err != nil {
			handlers.ThrowError(err, http.StatusBadRequest, w)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}

func NewEmailHandler(es ports.EmailServicer, logger *log.Logger) EmailHandler {
	return EmailHandler{
		emailService: es,
		logger:       logger,
	}
}
