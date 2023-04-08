package userconfig

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/handlers"
	"gorm.io/gorm"
)

type Handler struct {
	userConfigService ports.UserConfigServicer
}

func (h Handler) HandleUserConfig(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetUserConfig(w, r)
	case http.MethodPut:
		h.UpdateUserConfig(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h Handler) UpdateUserConfig(w http.ResponseWriter, r *http.Request) {
	payload := ports.UpdateUserConfigRequest{}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.userConfigService.Update(context.TODO(), r, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
}

func (h Handler) GetUserConfig(w http.ResponseWriter, r *http.Request) {
	c, err := h.userConfigService.Get(context.TODO(), r)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	b, err := json.Marshal(c)
	if err != nil {
		return
	}
	_, err = w.Write(b)
	if err != nil {
		return
	}
}

func NewUserConfigHandler(ucs ports.UserConfigServicer) Handler {
	return Handler{
		userConfigService: ucs,
	}
}
