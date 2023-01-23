package userconfig

import (
	"encoding/json"
	"errors"
	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/handlers"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type UserConfigHandler struct {
	userConfigService ports.UserConfigServicer
	logger            *log.Logger
}

func (h UserConfigHandler) HandleUserConfig(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.GetUserConfig(w, r)
	case http.MethodPut:
		h.UpdateUserConfig(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h UserConfigHandler) UpdateUserConfig(w http.ResponseWriter, r *http.Request) {
	payload := ports.UpdateUserConfigRequest{}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.userConfigService.Update(nil, r, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
}

func (h UserConfigHandler) GetUserConfig(w http.ResponseWriter, r *http.Request) {
	c, err := h.userConfigService.Get(nil, r)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	b, err := json.Marshal(c)
	w.Write(b)
}

func NewUserConfigHandler(ucs ports.UserConfigServicer, logger *log.Logger) UserConfigHandler {
	return UserConfigHandler{
		userConfigService: ucs,
		logger:            logger,
	}
}
