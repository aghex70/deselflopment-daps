package email

import (
	"github.com/aghex70/daps/internal/ports/handlers"
	"github.com/aghex70/daps/internal/ports/requests/email"
	"net/http"
)

type Handler struct {
	//emailService email.Servicer
}

func (h Handler) CreateEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)

		payload := requests.SendEmailRequest{}
		err := handlers.ValidateRequest(r, &payload)
		if err != nil {
			handlers.ThrowError(err, http.StatusBadRequest, w)
			return
		}

		//err = h.emailService.Send(context.TODO(), r, payload)
		//if err != nil {
		//	handlers.ThrowError(err, http.StatusBadRequest, w)
		//	return
		//}
		w.WriteHeader(http.StatusCreated)
	}
}

//func NewEmailHandler(es email.Servicer) Handler {
//	return Handler{
//emailService: es,
//}
//}
