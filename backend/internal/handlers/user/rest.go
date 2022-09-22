package user

import (
	"encoding/json"
	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/handlers"
	"log"
	"net/http"
)

type UserHandler struct {
	userService ports.UserServicer
	logger      *log.Logger
}

func (h UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	payload := ports.LoginUserRequest{}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	token, err := h.userService.Login(nil, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	b, err := json.Marshal(handlers.TokenResponse{AccessToken: token})
	w.Write(b)
}

func (h UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	panic("foo")
}

func (h UserHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	payload := ports.RefreshTokenRequest{}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	token, err := h.userService.RefreshToken(nil, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	b, err := json.Marshal(handlers.TokenResponse{AccessToken: token})
	w.Write(b)
}

func (h UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	payload := ports.CreateUserRequest{}
	err := handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.userService.Register(nil, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
}

func (h UserHandler) RemoveUser(w http.ResponseWriter, r *http.Request) {
	panic("foo")
}

func NewUserHandler(us ports.UserServicer, logger *log.Logger) UserHandler {
	return UserHandler{
		userService: us,
		logger:      logger,
	}
}
