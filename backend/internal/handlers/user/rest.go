package user

import (
	"encoding/json"
	"errors"
	"github.com/aghex70/daps/internal/core/ports"
	"github.com/aghex70/daps/internal/handlers"
	"log"
	"net/http"
)

type UserHandler struct {
	userService ports.UserServicer
	logger      *log.Logger
}

func (h UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	err := CheckHttpMethod(http.MethodPost, w, r)
	if err != nil {
		return
	}

	payload := ports.CreateUserRequest{}
	err = handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	err = h.userService.Register(nil, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	err := CheckHttpMethod(http.MethodPost, w, r)
	if err != nil {
		return
	}

	payload := ports.LoginUserRequest{}
	err = handlers.ValidateRequest(r, &payload)
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

func (h UserHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	err := CheckHttpMethod(http.MethodPost, w, r)
	if err != nil {
		return
	}

	payload := ports.RefreshTokenRequest{}
	err = handlers.ValidateRequest(r, &payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	token, err := h.userService.RefreshToken(nil, r, payload)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}

	b, err := json.Marshal(handlers.TokenResponse{AccessToken: token})
	w.Write(b)
}

func (h UserHandler) Logout(w http.ResponseWriter, r *http.Request) {
	err := CheckHttpMethod(http.MethodPost, w, r)
	if err != nil {
		return
	}
}

func (h UserHandler) RemoveUser(w http.ResponseWriter, r *http.Request) {
	err := CheckHttpMethod(http.MethodDelete, w, r)
	if err != nil {
		return
	}

	err = h.userService.Remove(nil, r)
	if err != nil {
		handlers.ThrowError(err, http.StatusBadRequest, w)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func CheckHttpMethod(status string, w http.ResponseWriter, r *http.Request) error {
	if r.Method != status {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return errors.New("method not allowed")
	}
	return nil
}

func NewUserHandler(us ports.UserServicer, logger *log.Logger) UserHandler {
	return UserHandler{
		userService: us,
		logger:      logger,
	}
}
